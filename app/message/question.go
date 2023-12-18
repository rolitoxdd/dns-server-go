package message

type Question struct {
	// The question section contains a list of questions (usually just 1) that the sender wants to ask the receiver. This section is present in both query and reply packets.
	// Each question has the following structure:

	Name []byte // Domain Name -  A domain name, represented as a sequence of "labels" (variable length)

	Type uint16 // Type - the type of record (1 for an A record, 5 for a CNAME record etc., full list: https://www.rfc-editor.org/rfc/rfc1035#section-3.2.2) (16 bits)

	Class uint16 // Class - the class of record (1 for Internet addresses, full list: https://www.rfc-editor.org/rfc/rfc1035#section-3.2.4) (16 bits)
}

func (x Question) ToBuf() []byte {
	// function that converts the body to a byte slice.
	buf := make([]byte, 0)

	buf = append(buf, x.Name...)
	buf = append(buf, byte(x.Type>>8))  // bits 0 - 8
	buf = append(buf, byte(x.Type))     // bits 9 - 16
	buf = append(buf, byte(x.Class>>8)) // bits 0 - 8
	buf = append(buf, byte(x.Class))    // bits 9 - 16

	return buf
}
