// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	ce "github.com/cloudevents/sdk-go/v2"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	log "k8s.io/klog/v2"

	"github.com/fatih/color"
	"github.com/linkall-labs/vanus/internal/kv"
	"github.com/linkall-labs/vanus/internal/kv/etcd"
	ctrlpb "github.com/linkall-labs/vanus/proto/pkg/controller"
	"github.com/linkall-labs/vanus/proto/pkg/meta"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	VolumeKeyPrefixInKVStore         = "/vanus/internal/resource/volume/metadata"
	BlockKeyPrefixInKVStore          = "/vanus/internal/resource/volume/block"
	VolumeInstanceKeyPrefixInKVStore = "/vanus/internal/resource/volume/instance"

	EventbusKeyPrefixInKVStore = "/vanus/internal/resource/eventbus"
	EventlogKeyPrefixInKVStore = "/vanus/internal/resource/eventlog"
	SegmentKeyPrefixInKVStore  = "/vanus/internal/resource/segment"

	EventlogSegmentsKeyPrefixInKVStore = "/vanus/internal/resource/segs_of_eventlog"
)

const (
	HttpPrefix = "http://"
	EventBus   = "e2e"
)

var (
	Sink        = "http://quick-display:80"
	Source      = ""
	Filters     = ""
	Transformer = ""

	EventType   = "examples"
	EventBody   = "Hello Vanus"
	EventSource = "e2e"

	HttpClient = resty.New()
	Endpoint   = os.Getenv("VANUS_GATEWAY")
	EtcdClient kv.Client
	err        error
)

func init() {
	kvStoreEndpoints := []string{"192.168.49.2:30007"}
	kvKeyPrefix := "/vanus"
	EtcdClient, err = etcd.NewEtcdClientV3(kvStoreEndpoints, kvKeyPrefix)
	if err != nil {
		log.Fatalf("NewEtcdClientV3 failed, err: %+v\n", err)
	}
}

func mustGetControllerProxyConn(ctx context.Context) *grpc.ClientConn {
	splits := strings.Split(os.Getenv("VANUS_GATEWAY"), ":")
	port, err := strconv.Atoi(splits[1])
	if err != nil {
		log.Error("parsing gateway port failed")
		return nil
	}
	leaderConn := createGRPCConn(ctx, fmt.Sprintf("%s:%d", splits[0], port+2))
	if leaderConn == nil {
		log.Error("failed to connect to gateway")
		return nil
	}
	return leaderConn
}

func createGRPCConn(ctx context.Context, addr string) *grpc.ClientConn {
	if addr == "" {
		return nil
	}
	addr = strings.TrimPrefix(addr, "http://")
	var err error
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())
	ctx, cancel := context.WithCancel(ctx)
	timeout := false
	go func() {
		ticker := time.Tick(time.Second)
		select {
		case <-ctx.Done():
		case <-ticker:
			cancel()
			timeout = true
		}
	}()
	conn, err := grpc.DialContext(ctx, addr, opts...)
	cancel()
	if timeout {
		color.Yellow("dial to controller: %s timeout, try to another controller", addr)
		return nil
	} else if err != nil {
		color.Red("dial to controller: %s failed", addr)
		return nil
	}
	return conn
}

func createEventbus(eb string) error {
	ctx := context.Background()
	grpcConn := mustGetControllerProxyConn(ctx)
	defer func() {
		_ = grpcConn.Close()
	}()

	cli := ctrlpb.NewEventBusControllerClient(grpcConn)
	res, err := cli.CreateEventBus(ctx, &ctrlpb.CreateEventBusRequest{
		Name: eb,
	})
	if err != nil {
		log.Errorf("create eventbus failed, err: %s", err)
		return err
	}
	log.Infof("create eventbus[%s] success.", res.Name)
	return nil
}

func createSubscription(eventbus, sink, source, filters, transformer string) (uint64, error) {
	ctx := context.Background()
	grpcConn := mustGetControllerProxyConn(ctx)
	defer func() {
		_ = grpcConn.Close()
	}()
	var filter []*meta.Filter
	if filters != "" {
		err := json.Unmarshal([]byte(filters), &filter)
		if err != nil {
			log.Errorf("the filter invalid, err: %s", err)
			return 0, err
		}
	}

	var trans *meta.Transformer
	if transformer != "" {
		err := json.Unmarshal([]byte(transformer), &trans)
		if err != nil {
			log.Errorf("the transformer invalid, err: %s", err)
			return 0, err
		}
	}

	cli := ctrlpb.NewTriggerControllerClient(grpcConn)
	res, err := cli.CreateSubscription(ctx, &ctrlpb.CreateSubscriptionRequest{
		Subscription: &ctrlpb.SubscriptionRequest{
			Source:      source,
			Filters:     filter,
			Sink:        sink,
			EventBus:    eventbus,
			Transformer: trans,
		},
	})
	if err != nil {
		log.Errorf("create subscription failed, err: %s", err)
		return 0, err
	}
	log.Infof("create subscription[%d] success.", res.Id)
	return res.Id, nil
}

func deleteSubscription(id uint64) error {
	ctx := context.Background()
	grpcConn := mustGetControllerProxyConn(ctx)
	defer func() {
		_ = grpcConn.Close()
	}()

	cli := ctrlpb.NewTriggerControllerClient(grpcConn)
	_, err := cli.DeleteSubscription(ctx, &ctrlpb.DeleteSubscriptionRequest{
		Id: id,
	})
	if err != nil {
		log.Errorf("delete subscription failed, err: %s", err)
		return err
	}
	log.Infof("delete subscription[%d] success.", id)
	return nil
}

func putEvent(eventbus, eventID, eventType, eventBody, eventSource string) error {
	client, err := ce.NewClientHTTP()
	if err != nil {
		log.Errorf("new cloudevents client failed, err: %s\n", err.Error())
		return err
	}

	if eventID == "" {
		eventID = uuid.NewString()
	}

	ctx := ce.ContextWithTarget(context.Background(), fmt.Sprintf("%s%s/gateway/%s", HttpPrefix, Endpoint, eventbus))
	event := ce.NewEvent()
	event.SetID(eventID)
	event.SetSource(eventSource)
	event.SetType(eventType)
	err = event.SetData(ce.TextPlain, eventBody)
	if err != nil {
		log.Errorf("set data failed: %s\n", err)
		return err
	}
	if result := client.Send(ctx, event); ce.IsUndelivered(result) {
		log.Errorf("failed to send event, err: %s\n", result.Error())
		return err
	}
	// log.Infof("put event[%s] success.", event.ID())
	return nil
}

func putEvents(offset, eventNum, threadNum int64, eventBus, eventBody, eventSource string) error {
	var (
		i       int64
		eventid int64 = offset
		wg      sync.WaitGroup
	)
	for i = 1; i <= threadNum; i++ {
		first := eventid
		last := eventid + eventNum/threadNum
		wg.Add(1)
		go func(first, last int64) {
			for n := first; n < last; n++ {
				putEvent(eventBus, fmt.Sprintf("%d", n), EventType, eventBody, eventSource)
			}
			wg.Done()
		}(first, last)
		eventid = eventid + eventNum/threadNum
	}
	wg.Wait()
	log.Infof("put %d events success.", eventNum)
	return nil
}

func getEvent(eventbus, offset, number string) error {
	idx := strings.LastIndex(Endpoint, ":")
	port, err := strconv.Atoi(Endpoint[idx+1:])
	if err != nil {
		log.Errorf("parse gateway port failed: %s, endpoint: %s", err, Endpoint)
		return err
	}
	newEndpoint := fmt.Sprintf("%s:%d", Endpoint[:idx], port+1)
	url := fmt.Sprintf("%s%s/getEvents?eventbus=%s&offset=%s&number=%s", HttpPrefix, newEndpoint, eventbus, offset, number)
	event, err := HttpClient.NewRequest().Get(url)
	if err != nil {
		log.Errorf("get event from eventbus[%s]&offset[%s]&number[%s] failed, err: %s\n", eventbus, offset, number, err)
		return err
	}
	log.Infof("get event from eventbus[%s]&offset[%s]&number[%s] success, event: %s\n", eventbus, offset, number, event.String())
	return nil
}

func Init() error {
	err := createEventbus(EventBus)
	if err != nil {
		return err
	}

	// wait for eventbus ready
	time.Sleep(3 * time.Second)
	return nil
}

func Test_e2e_base() {
	log.Info("Start test e2e basecase")
	eventBus := "eventbus-base"
	err = createEventbus(eventBus)
	if err != nil {
		return
	}

	id, err := createSubscription(EventBus, Sink, Source, Filters, Transformer)
	if err != nil {
		log.Error("create subsription failed")
		return
	}
	log.Info("create subsription success")

	err = putEvents(0, 10000, 100, EventBus, EventBody, EventSource)
	if err != nil {
		log.Error("put events failed")
		return
	}
	log.Info("put events success")

	err = getEvent(EventBus, "9999", "1")
	if err != nil {
		log.Error("get event failed")
		return
	}
	log.Info("get event success")

	deleteSubscription(id)
	log.Info("Finish test e2e basecase")
}

func Test_e2e_filter() {
	eventBus := "eventbus-filter"
	err = createEventbus(EventBus)
	if err != nil {
		return
	}

	filters := "[{\"exact\": {\"source\":\"filter\"}}]"
	err = createSubscription(EventBus, Sink, Source, filters, Transformer)
	if err != nil {
		return
	}

	filters = "[{\"cel\": \"$key.(string) == \\\"value\\\"\"}]"
	err = createSubscription(EventBus, Sink, Source, filters, Transformer)
	if err != nil {
		return
	}

	putEvents(0, 2000, 100, EventBus, EventBody, EventSource)
	eventSource := "filter"
	putEvents(2000, 4000, 10, EventBus, EventBody, eventSource)
	eventBody := "{\"key\":\"value\"}"
	putEvents(4000, 4000, 100, EventBus, eventBody, EventSource)

	err = getEvent(EventBus, "0", "8000", false)
	if err != nil {
		log.Error("Test_e2e_filter get event failed")
		return
	}
	log.Info("Test_e2e_filter get event success")
}

func Test_e2e_transformation() {
	eventBus := "eventbus-transformation"
	err = createEventbus(eventBus)
	if err != nil {
		return
	}

	transformer := "{\"template\": \"{\\\"transKey\\\": \\\"transValue\\\"}\"}"
	err = createSubscription(eventBus, Sink, Source, Filters, transformer)
	if err != nil {
		return
	}

	putEvents(0, 10000, 100, eventBus, EventBody, EventSource)

	err = getEvent(eventBus, "0", "10000", false)
	if err != nil {
		log.Error("Test_e2e_transformation get event failed")
		return
	}
	log.Info("Test_e2e_filter get event success")
}

func Test_e2e_metadata() {
	eventBus := "eventbus-meta"
	err = createEventbus(eventBus)
	if err != nil {
		return
	}

	// Currently, only check metadata of eventbus
	var path string = fmt.Sprintf("%s/%s", EventbusKeyPrefixInKVStore, eventBus)
	ctx := context.Background()
	meta, err := EtcdClient.Get(ctx, path)
	if err != nil {
		log.Errorf("get metadata failed, path: %s, err: %s\n", path, err.Error())
		return
	}
	log.Infof("get metadata success, path: %s, mata: %s\n", path, string(meta))
}

func main() {
	log.Info("start e2e test base case...")

	Test_e2e_base()

	// Test_e2e_filter()

	// Test_e2e_transformation()

	// Test_e2e_metadata()

	log.Info("finish e2e test base case...")
}
