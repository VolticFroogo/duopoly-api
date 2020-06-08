package message

const (
	RequestGreeting = iota
	RequestChat
	RequestStart
	RequestRoll
	RequestEndTurn
	RequestBuy
	RequestAuction
	RequestBid
	RequestPay
	RequestBuild
	RequestDemolish
	RequestMortgage
	RequestUnmortgage
	RequestBankrupt
)

const (
	ResponseError = iota
	ResponsePlaying
	ResponseGreeting
	ResponseJoined
	ResponseLeft
	ResponseChat
	ResponseStart
	ResponseRoll
	ResponseSetTurn
	ResponseBuy
	ResponseAuction
	ResponseBid
	ResponseAuctionWon
	ResponsePay
	ResponseBuild
	ResponseDemolish
	ResponseMortgage
	ResponseUnmortgage
	ResponseBankrupt
)

type Message struct {
	Type int         `json:"type"`
	Data interface{} `json:"data,omitempty"`
}
