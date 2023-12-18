package main

import m "github.com/codecrafters-io/dns-server-starter-go/app/message"

func createResponse(request m.Message) m.Message {
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
		RCODE:   0,
		QDCOUNT: request.Header.QDCOUNT,
		ANCOUNT: request.Header.ANCOUNT + 1,
		NSCOUNT: 0,
		ARCOUNT: 0,
	}

	// function that creates a response Question from a request Question
	question := m.Question{
		Name:  request.Question.Name,
		Type:  request.Question.Type,
		Class: request.Question.Class,
	}

	// function that creates a response Answer from a request Answer
	answer := m.Answer{
		Name:     request.Question.Name,
		Type:     request.Question.Type,
		Class:    request.Question.Class,
		TTL:      60,
		RDLENGTH: 4,
		RDATA:    []byte{8, 8, 8, 8},
	}

	// function that creates a response Message from a request Message
	response := m.Message{
		Header:   header,
		Question: question,
		Answer:   answer,
	}

	return response
}
