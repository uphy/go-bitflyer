package bitflyer

import (
	"encoding/json"

	"github.com/pubnub/go/messaging"
)

func (b *BitFlyer) RealtimeTicker(productCode string, r chan<- Ticker) {
	pubnub := messaging.NewPubnub("", "sub-c-52a9ab50-291b-11e5-baaa-0619f8945a4f", "", "", false, "", nil)
	successChannel, errorChannel := messaging.CreateSubscriptionChannels()
	pubnub.Subscribe("lightning_ticker_"+productCode, "", successChannel, false, errorChannel)
	go func() {
		for {
			value := <-successChannel
			resp := []interface{}{}
			if err := json.Unmarshal(value, &resp); err != nil {
				break
			}
			j, _ := json.Marshal(resp[0])
			if string(j) == "1" {
				continue
			}
			var ticker []Ticker
			if err := json.Unmarshal(j, &ticker); err != nil {
				break
			}
			r <- ticker[0]
		}
	}()
}
