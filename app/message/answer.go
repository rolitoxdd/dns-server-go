package message

type Answer struct {
	// The answer section contains a list of RRs (Resource Records), which are answers to the questions asked in the question section.
	// Each RR has the following structure:

	Name []byte // Domain Name -  A domain name, represented as a sequence of "labels" (variable length)

	Type uint16 // Type - the type of record (1 for an A record, 5 for a CNAME record etc., full list: https://www.rfc-editor.org/rfc/rfc1035#section-3.2.2) (16 bits)

	Class uint16 // Class - the class of record (1 for Internet addresses, full list: https://www.rfc-editor.org/rfc/rfc1035#section-3.2.4) (16 bits)

	TTL uint32 // TTL - the time interval (in seconds) that the record may be cached before it should be discarded. Zero values are interpreted to mean that the RR can only be used for the transaction in progress, and should not be cached (32 bits)

	RDLENGTH uint16 // RDLENGTH - the length of the RDATA field in bytes (16 bits)

	RDATA []byte // RDATA - Data specific to the record type.(variable length)
}

func (x Answer) ToBuf() []byte {
	// function that converts the body to a byte slice.
	buf := make([]byte, 0)

	buf = append(buf, x.Name...)
	buf = append(buf, byte(x.Type>>8)) // bits 0 - 8
	buf = append(buf, byte(x.Type))    // bits 9 - 16

	buf = append(buf, byte(x.Class>>8)) // bits 0 - 8
	buf = append(buf, byte(x.Class))    // bits 9 - 16

	buf = append(buf, byte(x.TTL>>24)) // bits 0 - 8
	buf = append(buf, byte(x.TTL>>16)) // bits 9 - 16
	buf = append(buf, byte(x.TTL>>8))  // bits 17 - 24
	buf = append(buf, byte(x.TTL))     // bits 25 - 32

	buf = append(buf, byte(x.RDLENGTH>>8)) // bits 0 - 8
	buf = append(buf, byte(x.RDLENGTH))    // bits 9 - 16

	buf = append(buf, x.RDATA...)

	return buf
}
