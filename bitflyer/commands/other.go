package commands

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	bitflyer "github.com/uphy/go-bitflyer"
)

func init() {
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

	root.AddCommand(markets)
	root.AddCommand(board)
	root.AddCommand(boardState)
}
