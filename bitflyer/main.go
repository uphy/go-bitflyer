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
	executions := &cobra.Command{
		Use:   "executions",
		Short: "Get executions.",
		RunE: func(cmd *cobra.Command, args []string) error {
			pageParams := bitflyer.NewPageParams()
			pageParams.Count = 5
			productCode, _ := cmd.Flags().GetString("productcode")
			executions, err := client.Executions(productCode, pageParams)
			if err != nil {
				return err
			}
			for _, execution := range executions {
				fmt.Println(execution)
			}
			return nil
		},
	}
	executions.Flags().String("productcode", bitflyer.ProductCodeBTCJPY, "the product code")
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
	board := &cobra.Command{
		Use:   "board",
		Short: "Get board information.",
		RunE: func(cmd *cobra.Command, args []string) error {
			productCode, _ := cmd.Flags().GetString("productcode")
			board, err := client.Board(productCode)
			if err != nil {
				return err
			}
			fmt.Println(board)
			return nil
		},
	}
	board.Flags().String("productcode", bitflyer.ProductCodeBTCJPY, "the product code")
	boardState := &cobra.Command{
		Use:   "boardstate",
		Short: "Get board state.",
		RunE: func(cmd *cobra.Command, args []string) error {
			productCode, _ := cmd.Flags().GetString("productcode")
			boardState, err := client.BoardState(productCode)
			if err != nil {
				return err
			}
			fmt.Printf("State  : %s\n", boardState.State)
			fmt.Printf("Health : %s\n", boardState.Health)
			return nil
		},
	}
	boardState.Flags().String("productcode", bitflyer.ProductCodeBTCJPY, "the product code")
	realtimeTicker := &cobra.Command{
		Use:   "realtime-ticker",
		Short: "Get ticker.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ticker := make(chan bitflyer.Ticker, 10)
			client.RealtimeTicker(bitflyer.ProductCodeBTCJPY, ticker)
			for t := range ticker {
				fmt.Println(int(t.LTP))
			}
			return nil
		},
	}
	realtimeExecutions := &cobra.Command{
		Use:   "realtime-executions",
		Short: "Get executions.",
		RunE: func(cmd *cobra.Command, args []string) error {
			executions := make(chan bitflyer.Execution, 10)
			client.RealtimeExecutions(bitflyer.ProductCodeBTCJPY, executions)
			for t := range executions {
				fmt.Println(t)
			}
			return nil
		},
	}
	root.AddCommand(ticker)
	root.AddCommand(executions)
	root.AddCommand(markets)
	root.AddCommand(board)
	root.AddCommand(boardState)
	root.AddCommand(realtimeTicker)
	root.AddCommand(realtimeExecutions)
	if err := root.Execute(); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Command failed. (err=%s)", err))
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
