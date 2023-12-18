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
		ANCOUNT: 1,
		NSCOUNT: 0,
		ARCOUNT: 0,
	}

	// function that creates a response Body from a request Body
	body := m.Question{
		Name:  request.Body.Name,
		Type:  request.Body.Type,
		Class: request.Body.Class,
	}

	// function that creates a response Message from a request Message
	response := m.Message{
		Header: header,
		Body:   body,
	}

	return response
}
