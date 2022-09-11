package function

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	ce "github.com/cloudevents/sdk-go/v2"
	"github.com/fatih/color"
	"github.com/linkall-labs/vanus/observability/log"
	"github.com/linkall-labs/vanus/test/timer/utils"
	"github.com/spf13/cobra"
)

func testcase2() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "testcase2",
		Short: "exec testcase2 of function testing",
		Run: func(cmd *cobra.Command, args []string) {
			// 1. environmental preparation
			var err error
			ctx := context.Background()
			eventbus := "testcase2"
			err = utils.CreateEventbus(ctx, eventbus)
			if err != nil {
				log.Error(ctx, "create eventbus failed", map[string]interface{}{
					"eventbus": eventbus,
				})
				utils.CmdFailedf(cmd, "create eventbus failed, eventbus: %s, err: %s", eventbus, err.Error())
			}
			ebClient := utils.NewEventbusClient(ctx, timerBuiltInEventbusReceivingStation)
			defer ebClient.Close(ctx)

			// 2. test data injection
			// record the time spent writing messages
			now := time.Now().UTC()
			numberOfEvents := 50
			var wg sync.WaitGroup
			for i := 0; i < numberOfEvents; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					event := ce.NewEvent()
					event.SetExtension(xceVanusEventbus, eventbus)
					event.SetExtension(xceVanusDeliveryTime, time.Now().UTC().Format(time.RFC3339))
					off, err := ebClient.Append(ctx, &event)
					if err != nil {
						utils.CmdFailedf(cmd, "ft testcase2 put event failed, off: %s, err: %s", off, err.Error())
					}
				}()
			}
			wg.Wait()
			log.Info(ctx, "time spent, unit: millisecounds", map[string]interface{}{
				"time": time.Now().UTC().Sub(now).Milliseconds(),
			})

			// 3. wait for program processing to complete
			log.Info(ctx, "waitting 4 until all event expired", nil)
			time.Sleep(4 * time.Second)

			// 4. result verification
			// 4.1 check the number of events received
			log.Info(ctx, "start check numbers of event in real eventbus", nil)
			elClient := utils.NewEventlogClient(ctx, eventbus)
			events, err := elClient.GetEvent(ctx, 0, int16(numberOfEvents+1))
			if err != nil {
				utils.CmdFailedf(cmd, "ft testcase2 get events failed, err: %s", err.Error())
			}
			if len(events) == numberOfEvents {
				log.Info(ctx, "check numbers of event in real eventbus success", map[string]interface{}{
					"expectl": numberOfEvents,
					"actual":  len(events),
				})
			} else {
				log.Error(ctx, "check numbers of event in real eventbus failed", map[string]interface{}{
					"expectl": numberOfEvents,
					"actual":  len(events),
				})
			}
			// 4.2 check timing message delivery time delay
			// todo, after long polling option merge

			// 5. environmental cleaning
			err = utils.DeleteEventbus(ctx, eventbus)
			if err != nil {
				log.Error(ctx, "delete eventbus failed", map[string]interface{}{
					"eventbus": eventbus,
				})
				utils.CmdFailedf(cmd, "delete eventbus failed, eventbus: %s, err: %s", eventbus, err.Error())
			}

			data, _ := json.Marshal(map[string]interface{}{"Result": 200})
			color.Green(string(data))
		},
	}
	return cmd
}
