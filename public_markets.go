package bitflyer

// Market represents market
type Market struct {
	ProductCode string `json:"product_code"`
	// 以下の呼出で product_code を指定するときに、代わりに使用できます。
	Alias string
}

// GetMarkets get the list of markets.
func (b *BitFlyer) Markets() ([]Market, error) {
	var markets []Market
	if err := b.get("markets", nil, &markets); err != nil {
		return nil, err
	}
	return markets, nil
}
