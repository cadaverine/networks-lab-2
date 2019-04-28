package main

import (
	"fmt"
	"net"
	"os"

	color "gopkg.in/gookit/color.v1"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s :port", os.Args[0])
		os.Exit(1)
	}

	port := os.Args[1]

	// ожидание ввода текста
	terminal := color.Style{color.FgCyan, color.OpBold}
	terminal.Printf("UDP server is listening on: localhost:%s\n", port)
	terminal.Println("Waiting request...")

	address, err := net.ResolveUDPAddr("udp", ":"+port)
	checkError(err)

	connection, err := net.ListenUDP("udp", address)
	checkError(err)

	buffer := []byte{}

	for {
		_, clientAddress, err := connection.ReadFromUDP(buffer)
		checkError(err)

		fmt.Println(buffer)

		connection.WriteToUDP(buffer, clientAddress)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
