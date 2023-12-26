package main

import m "github.com/codecrafters-io/dns-server-starter-go/app/message"

func createResponse(request m.Message) m.Message {
	RCODE := 4
	if request.Header.OPCODE == 0 {
		RCODE = 0
	}
	// function that creates a response Header from a request Header
	header := m.Header{
		ID:      request.Header.ID,
		QR:      true,
		OPCODE:  request.Header.OPCODE,
		AA:      false,
		TC:      false,
		RD:      request.Header.RD,
		RA:      false,
		Z:       0,
		RCODE:   byte(RCODE),
		QDCOUNT: request.Header.QDCOUNT,
		ANCOUNT: request.Header.QDCOUNT,
		NSCOUNT: 0,
		ARCOUNT: 0,
	}

	// function that creates a response Question from a request Question
	questions := make([]m.Question, 0)
	for _, question := range request.Questions {
		questions = append(questions, m.Question{
			Name:  question.Name,
			Type:  question.Type,
			Class: question.Class,
		})
	}

	// function that creates a response Answer from a request Answer
	answers := make([]m.Answer, 0)
	for _, question := range request.Questions {
		answers = append(answers, m.Answer{
			Name:     question.Name,
			Type:     question.Type,
			Class:    question.Class,
			TTL:      60,
			RDLENGTH: 4,
			RDATA:    []byte{8, 8, 8, 8},
		})
	}

	// function that creates a response Message from a request Message
	response := m.Message{
		Header:    header,
		Questions: questions,
		Answers:   answers,
	}

	return response
}
