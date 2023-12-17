package main

func createResponse(request Header) Header {
	// function that creates a response Header from a request Header
	response := Header{
		ID:      request.ID,
		QR:      true,
		OPCODE:  request.OPCODE,
		AA:      false,
		TC:      false,
		RD:      request.RD,
		RA:      false,
		Z:       0,
		RCODE:   0,
		QDCOUNT: request.QDCOUNT,
		ANCOUNT: 1,
		NSCOUNT: 0,
		ARCOUNT: 0,
	}
	return response
}
