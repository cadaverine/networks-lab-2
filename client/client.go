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

	// ожидание ввода текста
	terminal := color.Style{color.FgCyan, color.OpBold}
	terminal.Println("UDP client is started.")
	terminal.Println("Please type a text:")

	stringToSend, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// подключение к серверу
	address := os.Args[1]

	connection, err := net.Dial("udp", address)
	checkError(err)

	defer connection.Close()

	fmt.Fprintf(connection, stringToSend)

	bytesFromServer := []byte{}

	length, err := connection.Read(bytesFromServer)
	checkError(err)

	stringFromServer := string(bytesFromServer[0:length])

	terminal.Println("Server answer:")
	terminal.Println(stringFromServer)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
