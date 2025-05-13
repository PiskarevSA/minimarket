package events

type Event string

func (e Event) String() string {
	return string(e)
}

const (
	BalanceDeposited Event = "POINTS.DEPOSITED"
	BalanceWithdrawn Event = "POINTS.WITHDRAWN"
)

type BalanceChanged struct {
	OrderId string `json:"orderId"`
	UserId  string `json:"userId"`
	Amount  string `json:"amount"`
}
