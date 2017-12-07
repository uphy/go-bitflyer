package bitflyer

import (
	"net/url"
)

type Execution struct {
	ID                         int64
	Side                       string
	Price                      float64
	Size                       float64
	ExecDate                   string `json:"exec_date"`
	BuyChildOrdeAcceptanceID   string `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceID string `json:"sell_child_order_acceptance_id"`
}

// Executions gets the executions.
func (b *BitFlyer) Executions(productCode string, pageParams *PageParams) ([]Execution, error) {
	var executions []Execution
	params := url.Values{}
	params.Add("product_code", productCode)
	pageParams.QueryParams(params)
	if err := b.get("executions", params, &executions); err != nil {
		return nil, err
	}
	return executions, nil
}
