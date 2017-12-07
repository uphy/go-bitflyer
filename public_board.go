package bitflyer

import "net/url"

type Board struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []BoardPrice
	Asks     []BoardPrice
}

type BoardPrice struct {
	Price float64
	Size  float64
}

// Board gets the board information.
func (b *BitFlyer) Board(productCode string) (*Board, error) {
	var board Board
	params := url.Values{}
	params.Add("product_code", productCode)
	if err := b.get("board", params, &board); err != nil {
		return nil, err
	}
	return &board, nil
}
