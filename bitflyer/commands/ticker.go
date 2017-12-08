package commands

import (
	"fmt"
	"os"

	"github.com/gizak/termui"
	"github.com/spf13/cobra"
	bitflyer "github.com/uphy/go-bitflyer"
	"github.com/uphy/go-bitflyer/io"
)

func init() {
	ticker := &cobra.Command{
		Use:   "ticker",
		Short: "Get tickers",
		RunE: func(cmd *cobra.Command, args []string) error {
			var productCode string
			if len(args) == 0 {
				productCode = bitflyer.ProductCodeBTCJPY
			} else {
				productCode = args[0]
			}
			isRealtime, _ := cmd.Flags().GetBool("realtime")
			if isRealtime {
				ticker := make(chan bitflyer.Ticker, 10)
				client.RealtimeTicker(bitflyer.ProductCodeBTCJPY, ticker)
				if output, _ := cmd.Flags().GetString("output"); output != "" {
					f, err := os.Create(output)
					if err != nil {
						return err
					}
					defer f.Close()
					writer := io.NewTickerWriter(f)
					for t := range ticker {
						writer.Write(&t)
						fmt.Println(t)
					}
				} else {
					if err := termui.Init(); err != nil {
						return err
					}
					defer termui.Close()
					table := termui.NewTable()
					table.Width = 70
					table.Height = 9
					table.X = 0
					table.Y = 0
					go func() {
						for t := range ticker {
							table.BorderLabel = "Ticker: " + t.Timestamp
							rows := [][]string{
								[]string{"LTP", fmt.Sprintf("%.f", t.LTP)},
								[]string{"Best", fmt.Sprintf("Ask %.f (%6f) / Bid %.f (%6f)", t.BestAsk, t.BestAskSize, t.BestBid, t.BestBidSize)},
								[]string{"Depth", fmt.Sprintf("Ask %6f / Bid %6f", t.TotalAskDepth, t.TotalBidDepth)},
								[]string{"Volume", fmt.Sprintf("%6f", t.Volume)},
							}

							table.Rows = rows
							termui.Render(table)
						}
					}()
					termui.Handle("/sys/kbd/q", func(termui.Event) {
						termui.StopLoop()
					})
					termui.Loop()
				}
			} else {
				ticker, err := client.Ticker(productCode)
				if err != nil {
					return err
				}
				fmt.Println(ticker)
			}
			return nil
		},
	}
	ticker.Flags().String("productcode", bitflyer.ProductCodeBTCJPY, "the product code")
	ticker.Flags().BoolP("realtime", "f", false, "realtime")
	ticker.Flags().StringP("output", "o", "", "output filepath")
	root.AddCommand(ticker)
}
