package message

type Message struct {
	Header Header
	Body   Question
}

func (x Message) ToBuf() []byte {
	// function that converts the message to a byte slice.
	buf := make([]byte, 0)

	buf = append(buf, x.Header.ToBuf()...)
	buf = append(buf, x.Body.ToBuf()...)

	return buf
}
