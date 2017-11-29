package bitflyer

import (
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
