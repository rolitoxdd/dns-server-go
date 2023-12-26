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

func BufToQuestion(buf []byte, start int) (Question, error) {

	question := Question{}

	var i int
	for i = start; i < start+len(buf[start:]); i++ {
		// search for the null byte (that's the end of the Name)
		if buf[i] == 0 { // 0x00 = 0000 0000
			question.Name = buf[start : i+1]
			break
		}
		if buf[i]&0xC0 == 0xC0 { // 0xC0 = 1100 0000
			// if the first two bits are 11, then the next 14 bits are an offset to the real location of the Name
			// the real location of the Name is the offset concatenated with the rest of the buffer
			// the offset is the last 14 bits of the first byte concatenated with the second byte
			offset := binary.BigEndian.Uint16([]byte{buf[i], buf[i+1]}) & 0x3FFF // 0x3FFF = 0011 1111 1111 1111
			r_question, err := BufToQuestion(buf, int(offset))
			if err != nil {
				panic(err)
			}
			question.Name = append(buf[start:i], r_question.Name...)
			break
		}
	}
	question.Type = binary.BigEndian.Uint16(buf[i+1 : i+3])
	question.Class = binary.BigEndian.Uint16(buf[i+3 : i+5])

	return question, nil
}

func BufToQuestions(header Header, buf []byte) ([]Question, error) {

	questions := make([]Question, 0)
	start := 12
	for i := 0; i < int(header.QDCOUNT); i++ {
		question, err := BufToQuestion(buf, start)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
		if start <= len(buf) {
			start += len(question.Name) + 4
		}
	}

	return questions, nil
}

func BufToAnswer(buf []byte) (Answer, error) {

	answer := Answer{}

	var i int
	for i = 0; i < len(buf); i++ {
		// search for the null byte (that's the end of the Name)
		if buf[i] == 0 { // 0x00 = 0000 0000
			answer.Name = buf[:i+1]
			break
		}
	}
	answer.Type = binary.BigEndian.Uint16(buf[i+1 : i+3])
	answer.Class = binary.BigEndian.Uint16(buf[i+3 : i+5])
	answer.TTL = binary.BigEndian.Uint32(buf[i+5 : i+9])
	answer.RDLENGTH = binary.BigEndian.Uint16(buf[i+9 : i+11])
	answer.RDATA = buf[i+11 : i+11+int(answer.RDLENGTH)]

	return answer, nil
}

// function that converts a byte slice to a Message
func BufToMessage(buf []byte) (Message, error) {

	header, err := BufToHeader(buf)
	if err != nil {
		return Message{}, err
	}

	questions, err := BufToQuestions(header, buf)
	if err != nil {
		return Message{}, err
	}

	return Message{
		Header:    header,
		Questions: questions,
	}, nil
}
