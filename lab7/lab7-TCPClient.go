package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Enter a message to send to the server:")
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	// Отправка сообщения на сервер
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	// Ожидание ответа от сервера
	serverResponse, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Server response:", serverResponse)

}
