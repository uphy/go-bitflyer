package bitflyer

import "net/url"

const (
	HealthNormal    = "NORMAL"
	HealthBusy      = "BUSY"
	HealthVeryBusy  = "VERY BUSY"
	HealthSuperBusy = "SUPER BUSY"
	HealthNoOrder   = "NO ORDER"
	HealthStop      = "STOP"

	// 通常稼働中
	StateRunning = "RUNNING"
	// 取引停止中
	StateClosed = "CLOSED"
	// 再起動中
	StateStarting = "STARTING"
	// 板寄せ中
	StatePreOpen = "PREOPEN"
	// サーキットブレイク発動中
	StateCircuitBreak = "CIRCUIT BREAK"
	// Lightning Futures の取引終了後 SQ（清算値）の確定前
	StateAwaitingSQ = "AWAITING SQ"
	// Lightning Futures の満期に到達
	StateMatured = "MATURED"
)

type BoardState struct {
	Health string
	State  string
	Data   *BoardStateData
}

type BoardStateData struct {
	SpecialQuotation int64 `json:"special_quotation"`
}

// BoardState gets the state of the board.
func (b *BitFlyer) BoardState(productCode string) (*BoardState, error) {
	var boardState BoardState
	params := url.Values{}
	params.Add("product_code", productCode)
	if err := b.get("getboardstate", params, &boardState); err != nil {
		return nil, err
	}
	return &boardState, nil
}
