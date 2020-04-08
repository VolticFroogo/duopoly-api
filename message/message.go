package message

const (
	RequestChat = iota
)

const (
	ResponseError = iota
	ResponsePlaying
	ResponseGreeting
	ResponseJoined
	ResponseLeft
	ResponseChat
)

type Message struct {
	Type int         `json:"type"`
	Data interface{} `json:"data,omitempty"`
}
