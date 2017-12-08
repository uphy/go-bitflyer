package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	bitflyer "github.com/uphy/go-bitflyer"
)

func init() {
	executions := &cobra.Command{
		Use:   "executions",
		Short: "Get executions.",
		RunE: func(cmd *cobra.Command, args []string) error {
			var productCode string
			if len(args) == 0 {
				productCode = bitflyer.ProductCodeBTCJPY
			} else {
				productCode = args[0]
			}
			isRealtime, _ := cmd.Flags().GetBool("realtime")
			if isRealtime {
				executions := make(chan bitflyer.Execution, 10)
				client.RealtimeExecutions(bitflyer.ProductCodeBTCJPY, executions)
				for t := range executions {
					fmt.Println(t)
				}
			} else {
				pageParams := bitflyer.NewPageParams()
				pageParams.Count, _ = cmd.Flags().GetInt("count")
				pageParams.Before, _ = cmd.Flags().GetInt64("before")
				pageParams.After, _ = cmd.Flags().GetInt64("after")
				executions, err := client.Executions(productCode, pageParams)
				if err != nil {
					return err
				}
				for _, execution := range executions {
					fmt.Println(execution)
				}
			}
			return nil
		},
	}
	executions.Flags().String("productcode", bitflyer.ProductCodeBTCJPY, "the product code")
	executions.Flags().BoolP("realtime", "f", false, "realtime")
	executions.Flags().IntP("count", "c", 200, "the count of the page")
	executions.Flags().Int64P("before", "b", -1, "get before this ID")
	executions.Flags().Int64P("after", "a", -1, "get after this ID")
	root.AddCommand(executions)
}
