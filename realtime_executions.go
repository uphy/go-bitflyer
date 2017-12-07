package bitflyer

import (
	"encoding/json"
	"fmt"

	"github.com/pubnub/go/messaging"
)

func (b *BitFlyer) RealtimeExecutions(productCode string, r chan<- Execution) {
	pubnub := messaging.NewPubnub("", "sub-c-52a9ab50-291b-11e5-baaa-0619f8945a4f", "", "", false, "", nil)
	successChannel, errorChannel := messaging.CreateSubscriptionChannels()
	pubnub.Subscribe("lightning_executions_"+productCode, "", successChannel, false, errorChannel)
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
			var executions [][]Execution
			if err := json.Unmarshal(j, &executions); err != nil {
				fmt.Println(err)
				break
			}
			e := executions[0]
			for _, ee := range e {
				r <- ee
			}
		}
	}()
}
