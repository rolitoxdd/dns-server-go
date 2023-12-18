package message

type Header struct {
	// The header section is always 12 bytes long. Integers are encoded in big-endian format.

	ID uint16 // Packet Identifier - A random ID assigned to query packets. Response packets must reply with the same ID. (16 bits)

	QR     bool // Query/Response Indicator - 1 for a reply packet, 0 for a question packet. (1 bit)
	OPCODE byte // Operation Code - Specifies the kind of query in a message. (4 bits)
	AA     bool // Authoritative Answer - 1 if the responding server "owns" the domain queried, i.e., it's authoritative. (1 bit)
	TC     bool // Truncation - 1 if the message is larger than 512 bytes. Always 0 in UDP responses. (1 bit)
	RD     bool // Recursion Desired - Sender sets this to 1 if the server should recursively resolve this query, 0 otherwise. (1 bit)

	RA    bool // Recursion Available - Server sets this to 1 to indicate that recursion is available. (1 bit)
	Z     byte // Used by DNSSEC queries. At inception, it was reserved for future use. (3 bits)
	RCODE byte // Response Code - Response code indicating the status of the response. (4 bits)

	QDCOUNT uint16 // Question Count - Number of questions in the Question section. (16 bits)

	ANCOUNT uint16 // Answer Record Count - Number of records in the Answer section. (16 bits)

	NSCOUNT uint16 // Authority Record Count - Number of records in the Authority section. (16 bits)

	ARCOUNT uint16 // Additional Record Count - Number of records in the Additional section. (16 bits)

}

func (x Header) ToBuf() []byte {
	// function that converts the header to a byte slice.
	buf := make([]byte, 12)

	// bits 0 - 16
	buf[0] = byte(x.ID >> 8) // bits 0 - 8
	buf[1] = byte(x.ID)      // bits 9 - 16

	// bits 17 - 24
	buf[2] = 0
	if x.QR {
		// bit 17 will come from QR indicator
		buf[2] |= 0x80 // 0x80 = 1000 0000
	}
	// bits 18 - 21 will come from OPCODE
	buf[2] |= x.OPCODE << 3 // shift left 3 bits to get bits 18 - 21 0??? ?000
	if x.AA {
		// bit 22 will come from AA indicator
		buf[2] |= 0x04 // 0x04 = 0000 0100
	}
	if x.TC {
		// bit 23 will come from TC indicator
		buf[2] |= 0x02 // 0x02 = 0000 0010
	}
	if x.RD {
		// bit 24 will come from RD indicator
		buf[2] |= 0x01 // 0x01 = 0000 0001
	}

	// bits 25 - 32
	buf[3] = 0
	if x.RA {
		// bit 25 will come from RA indicator
		buf[3] |= 0x80 // 0x80 = 1000 0000
	}
	// bits 26 - 28 will come from Z
	buf[3] |= x.Z << 4 // shift left 4 bits to get bits 26 - 28 0??? ?000
	// bits 29 - 32 will come from RCODE
	buf[3] |= x.RCODE // 0x0F = 0000 ????

	// bits 33 - 48
	buf[4] = byte(x.QDCOUNT >> 8) // bits 33 - 40
	buf[5] = byte(x.QDCOUNT)      // bits 41 - 48

	// bits 49 - 64
	buf[6] = byte(x.ANCOUNT >> 8) // bits 49 - 56
	buf[7] = byte(x.ANCOUNT)      // bits 57 - 64

	// bits 65 - 80
	buf[8] = byte(x.NSCOUNT >> 8) // bits 65 - 72
	buf[9] = byte(x.NSCOUNT)      // bits 73 - 80

	// bits 81 - 96
	buf[10] = byte(x.ARCOUNT >> 8) // bits 81 - 88
	buf[11] = byte(x.ARCOUNT)      // bits 89 - 96

	return buf
}

// The header section is always 12 bytes long. Integers are encoded in big-endian format.
