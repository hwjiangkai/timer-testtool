package function

import (
	"github.com/spf13/cobra"
)

const (
	xceVanusEventbus                     = "xvanuseventbus"
	xceVanusDeliveryTime                 = "xvanusdeliverytime"
	timerBuiltInEventbusReceivingStation = "__Timer_RS"
)

func NewFunctionTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ft sub-command ",
		Short: "convenient operations for function testing",
	}
	cmd.AddCommand(testcase1())
	cmd.AddCommand(testcase2())
	cmd.AddCommand(testcase3())
	cmd.AddCommand(testcase4())
	return cmd
}
