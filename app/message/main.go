package message

// The message is the packet that is sent between the client and the server. It contains a header and a body.
type Message struct {

	Header Header

	// Body
	Questions []Question
	Answers   []Answer
}

// function that converts the message to a byte slice.
func (x Message) ToBuf() []byte {
	buf := make([]byte, 0)

	buf = append(buf, x.Header.ToBuf()...)

	for _, question := range x.Questions {
		buf = append(buf, question.ToBuf()...)
	}

	for _, answer := range x.Answers {
		buf = append(buf, answer.ToBuf()...)
	}


	return buf
}
