package message

// The message is the packet that is sent between the client and the server. It contains a header and a body.
type Message struct {

	Header Header

	// Body
	Question Question
	Answer   Answer
}

// function that converts the message to a byte slice.
func (x Message) ToBuf() []byte {
	buf := make([]byte, 0)

	buf = append(buf, x.Header.ToBuf()...)
	buf = append(buf, x.Question.ToBuf()...)
	buf = append(buf, x.Answer.ToBuf()...)

	return buf
}
