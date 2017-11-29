package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/uphy/go-bitflyer"
)

func main() {
	client := bitflyer.New()
	root := &cobra.Command{
		Use:   "bitflyer",
		Short: "bitflyer parent command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	ticker := &cobra.Command{
		Use:   "ticker",
		Short: "Get ticker.",
		RunE: func(cmd *cobra.Command, args []string) error {
			productCode, _ := cmd.Flags().GetString("productcode")
			ticker, err := client.Ticker(productCode)
			if err != nil {
				return err
			}
			fmt.Printf("Product Code: %s\n", ticker.ProductCode)
			fmt.Printf("LTP         : %-8.3f\n", ticker.LTP)
			return nil
		},
	}
	ticker.Flags().String("productcode", bitflyer.ProductCodeBTCJPY, "the product code")
	markets := &cobra.Command{
		Use:   "markets",
		Short: "Get markets.",
		RunE: func(cmd *cobra.Command, args []string) error {
			markets, err := client.Markets()
			if err != nil {
				return err
			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Product Code", "Alias"})
			for _, market := range markets {
				table.Append([]string{market.ProductCode, market.Alias})
			}
			table.Render()
			return nil
		},
	}
	root.AddCommand(ticker)
	root.AddCommand(markets)
	if err := root.Execute(); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Command failed. (err=%s)", err))
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
