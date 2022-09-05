package utils

import (
	"context"
	"fmt"
	"strings"

	client "github.com/linkall-labs/vanus/client"
	eb "github.com/linkall-labs/vanus/client/pkg/eventbus"
	eventlog "github.com/linkall-labs/vanus/client/pkg/eventlog"
	"github.com/linkall-labs/vanus/observability/log"
	ctrlpb "github.com/linkall-labs/vanus/proto/pkg/controller"
	meta "github.com/linkall-labs/vanus/proto/pkg/meta"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	EtcdEndpoints = []string{"vanus-controller-0.vanus-controller:2379", "vanus-controller-1.vanus-controller:2379", "vanus-controller-2.vanus-controller:2379"}
	CtrlEndpoints = []string{"vanus-controller-0.vanus-controller.vanus.svc:2048", "vanus-controller-1.vanus-controller.vanus.svc:2048", "vanus-controller-2.vanus-controller.vanus.svc:2048"}
)

func ListEventbus(ctx context.Context) (*ctrlpb.ListEventbusResponse, error) {
	eventbusCtrlCli := NewClient(CtrlEndpoints).makeSureClient(ctx, true)
	return eventbusCtrlCli.ListEventBus(ctx, &emptypb.Empty{})
}

func CreateEventbus(ctx context.Context, eventbus string) error {
	eventbusCtrlCli := NewClient(CtrlEndpoints).makeSureClient(ctx, true)
	_, err := eventbusCtrlCli.CreateEventBus(ctx, &ctrlpb.CreateEventBusRequest{
		Name: eventbus,
	})
	if err != nil {
		log.Error(ctx, "create eventbus failed", map[string]interface{}{
			log.KeyError: err,
			"eventbus":   eventbus,
		})
		return err
	}
	log.Info(ctx, "create eventbus success.", map[string]interface{}{
		"eventbus": eventbus,
	})
	return nil
}

func DeleteEventbus(ctx context.Context, eventbus string) error {
	eventbusCtrlCli := NewClient(CtrlEndpoints).makeSureClient(ctx, true)
	_, err := eventbusCtrlCli.DeleteEventBus(ctx, &meta.EventBus{
		Name: eventbus,
	})
	if err != nil {
		log.Error(ctx, "delete eventbus failed", map[string]interface{}{
			log.KeyError: err,
			"eventbus":   eventbus,
		})
		return err
	}
	log.Info(ctx, "delete eventbus success.", map[string]interface{}{
		"eventbus": eventbus,
	})
	return nil
}

type ebClient struct {
	eventbus       string
	eventlogWriter eventlog.LogWriter
	eventlogReader eventlog.LogReader
}

func NewEventlogClient(ctx context.Context, eventbus string) *ebClient {
	vrn := fmt.Sprintf("vanus:///eventbus/%s?controllers=%s", eventbus, strings.Join(CtrlEndpoints, ","))
	ls, _ := client.LookupReadableLogs(ctx, vrn)
	eventlogReader, _ := client.OpenLogReader(ls[0].VRN)
	eventlogWriter, _ := client.OpenLogWriter(ls[0].VRN)
	return &ebClient{
		eventbus:       eventbus,
		eventlogReader: eventlogReader,
		eventlogWriter: eventlogWriter,
	}
}

func NewEventbusClient(ctx context.Context, eventbus string) eb.BusWriter {
	vrn := fmt.Sprintf("vanus:///eventbus/%s?controllers=%s", eventbus, strings.Join(CtrlEndpoints, ","))
	eventbusWriter, _ := client.OpenBusWriter(vrn)
	return eventbusWriter
}
