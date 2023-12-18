package message

type IMessage interface {
	// function that converts the message to a byte slice.
	ToBuf() []byte
}
