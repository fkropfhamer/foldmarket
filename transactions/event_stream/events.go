package event_stream

const DepositEventType string = "DepositEvent"
const WithdrawEventType string = "WithdrawEvent"

type DepositEvent struct {
	AccountId int32 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

type WithdrawEvent struct {
	AccountId int32 `json:"account_id"`
	Amount    int64 `json:"amount"`
}
