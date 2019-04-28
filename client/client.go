package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	color "gopkg.in/gookit/color.v1"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	// подключение к серверу
	address, err := net.ResolveUDPAddr("udp", os.Args[1])
	checkError(err)

	connection, err := net.DialUDP("udp", nil, address)
	checkError(err)

	defer connection.Close()

	// ожидание ввода текста
	terminal := color.Style{color.FgCyan}
	terminal.Println("UDP client is started.")

	for {
		terminal.Println("\nPlease type a text:")

		message, _ := bufio.NewReader(os.Stdin).ReadBytes('\n')

		connection.Write(message)

		bytesFromServer := make([]byte, 512)

		length, err := connection.Read(bytesFromServer)
		checkError(err)

		stringFromServer := string(bytesFromServer[0:length])

		terminal.Println("\nServer answer:")
		fmt.Println(stringFromServer)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
