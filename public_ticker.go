package bitflyer

import (
	"fmt"
	"net/url"
)

// Ticker represents ticker
type Ticker struct {
	ProductCode string `json:"product_code"`
	// 時刻は UTC（協定世界時）で表されます。
	Timestamp     string  `json:"timestamp"`
	TickID        int64   `json:"tick_id"`
	BestBid       float64 `json:"best_bid"`
	BestAsk       float64 `json:"best_ask"`
	BestBidSize   float64 `json:"best_bid_size"`
	BestAskSize   float64 `json:"best_ask_size"`
	TotalBidDepth float64 `json:"total_bid_depth"`
	TotalAskDepth float64 `json:"total_ask_depth"`
	// 最終取引価格
	LTP float64 `json:"ltp"`
	// 24 時間の取引量
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

func (t Ticker) String() string {
	return fmt.Sprintf("Ticker(date=%s, id=%d, ltp=%8.1f, bestBid=%8.1f, bestAsk=%8.1f, bestBidSize=%5.4f, bestAskSize=%5.4f, totalBidDepth=%f, totalAskDepth=%f)", t.Timestamp, t.TickID, t.LTP, t.BestBid, t.BestAsk, t.BestBidSize, t.BestAskSize, t.TotalBidDepth, t.TotalAskDepth)
}

const (
	ProductCodeBTCJPY = "BTC_JPY"
)

// Ticker gets the ticker.
func (b *BitFlyer) Ticker(productCode string) (*Ticker, error) {
	var ticker Ticker
	params := url.Values{}
	params.Add("product_code", productCode)
	if err := b.get("ticker", params, &ticker); err != nil {
		return nil, err
	}
	return &ticker, nil
}
