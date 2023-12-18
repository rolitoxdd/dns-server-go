package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"

	m "github.com/codecrafters-io/dns-server-starter-go/app/message"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	udpAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		receivedData, err := m.BufToMessage(buf[:size])
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received %d bytes from %s\n", size, source)

		response := createResponse(receivedData)

		_, err = udpConn.WriteToUDP(response.ToBuf(), source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
