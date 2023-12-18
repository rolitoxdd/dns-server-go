package message

type Message struct {
	// The message is the packet that is sent between the client and the server. It contains a header and a body.

	Header Header

	// Body
	Question Question
	Answer   Answer
}

func (x Message) ToBuf() []byte {
	// function that converts the message to a byte slice.
	buf := make([]byte, 0)

	buf = append(buf, x.Header.ToBuf()...)
	buf = append(buf, x.Question.ToBuf()...)
	buf = append(buf, x.Answer.ToBuf()...)

	return buf
}
