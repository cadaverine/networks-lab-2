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
	terminal := color.Style{color.FgCyan}
	terminal.Printf("UDP server is listening on: localhost:%s\n", port)
	terminal.Printf("Waiting request...\n\n")

	address, err := net.ResolveUDPAddr("udp", ":"+port)
	checkError(err)

	connection, err := net.ListenUDP("udp", address)
	checkError(err)

	buffer := make([]byte, 512)

	for {
		length, clientAddress, err := connection.ReadFromUDP(buffer)
		checkError(err)

		message := string(buffer[0:length-1])

		terminal.Printf("Client address: %s\n", clientAddress)
		terminal.Printf("Client message: %s\n", message)

		reversed := reverseString(message)

		connection.WriteToUDP([]byte(reversed), clientAddress)

		terminal.Printf("Reversed message ('%s') was sended back.\n\n", reversed)
	}
}

func reverseString(str string) string {
	length := len(str)
	runes := make([]rune, length)

	for i, char := range str {
		runes[length - (i + 1)] = char
	}

	return string(runes)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
