package game

const (
	_ = iota // Null action.
	actionTurn
	actionPurchaseDecision
	actionAuction
	actionPay
)

type Action struct {
	Type int         `json:"type"`
	Data interface{} `json:"data,omitempty"`
}

func (action Action) Tick(game *Game) {
	switch action.Type {
	case actionTurn:
		tickTurn(game)
		break

	case actionPurchaseDecision:
		tickPurchaseDecision(game)
		break

	case actionAuction:
		tickAuction(game)
		break

	case actionPay:
		tickPay(game)
		break
	}
}
