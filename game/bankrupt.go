package game

func (game *Game) tryBankrupt(playerID int) {
	if game.Action.Type != actionPay {
		game.error(playerID, "cannot bankrupt: not action pay")
		return
	}

	data := game.Action.Data.(dataPay)

	if data.Payer != playerID {
		game.error(playerID, "cannot bankrupt: not payer")
		return
	}

	game.bankrupt(data.Payer, data.Payee)
}

func (game *Game) bankrupt(payer int, payee int) {
	// TODO: implement proper bankrupting.
	game.Players[payer].Bankrupt = true
	game.endTurn(payer, true)
}
