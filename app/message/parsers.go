package message

import (
	"encoding/binary"
	"errors"
)

func BufToHeader(buf []byte) (Header, error) {

	if len(buf) < 12 {
		return Header{}, errors.New("Header too short")
	}

	header := Header{
		ID:      binary.BigEndian.Uint16(buf[0:2]), // bits 0 - 16
		QR:      buf[2]&0x80 != 0,                  // bit 17: 0x80 = 1000 0000 , will be 0 if 0, otherwise
		OPCODE:  buf[2] & 0x78 >> 3,                // bits 18 - 21: 0x78 = 0111 1000, shift right 3 bits
		AA:      buf[2]&0x04 != 0,                  // bit 22: 0x04 = 0000 0100
		TC:      buf[2]&0x02 != 0,                  // bit 23: 0x02 = 0000 0010
		RD:      buf[2]&0x01 != 0,                  // bit 24: 0x01 = 0000 0001
		RA:      buf[3]&0x80 != 0,                  // bit 25: 0x80 = 1000 0000
		Z:       buf[3] & 0x70 >> 4,                // bits 26 - 28: 0x70 = 0111 0000, shift right 4 bits
		RCODE:   buf[3] & 0x0F,                     // bits 29 - 32: 0x0F = 0000 1111
		QDCOUNT: binary.BigEndian.Uint16(buf[4:6]),
		ANCOUNT: binary.BigEndian.Uint16(buf[6:8]),
		NSCOUNT: binary.BigEndian.Uint16(buf[8:10]),
		ARCOUNT: binary.BigEndian.Uint16(buf[10:12]),
	}

	return header, nil
}

func BufToBody(buf []byte) (Question, error) {

	body := Question{}

	var i int
	for i = 0; i < len(buf); i++ {
		// search for the null byte (that's the end of the Name)
		if buf[i] == 0 { // 0x00 = 0000 0000
			body.Name = buf[:i+1]
			break
		}
	}
	body.Type = binary.BigEndian.Uint16(buf[i+1 : i+3])
	body.Class = binary.BigEndian.Uint16(buf[i+3 : i+5])

	return body, nil
}

func BufToMessage(buf []byte) (Message, error) {

	header, err := BufToHeader(buf)
	if err != nil {
		return Message{}, err
	}

	body, err := BufToBody(buf[12:])
	if err != nil {
		return Message{}, err
	}

	return Message{
		Header: header,
		Body:   body,
	}, nil
}