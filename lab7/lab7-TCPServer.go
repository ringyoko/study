package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup
var stop = make(chan struct{})

func handleConnection(conn net.Conn) {
	defer wg.Done()
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		select {
		case <-stop:
			fmt.Println("Server is shutting down")
			return
		default:
			// Чтение сообщения от клиента
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Client disconnected:", conn.RemoteAddr())
				return
			}

			fmt.Printf("Received message from %s: %s", conn.RemoteAddr(), message)

			_, err = conn.Write([]byte("Message received\n"))
		}
	}
}

func startServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting TCP server: %s", err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-stop:
				fmt.Println("Server stopped accepting new connections")
				return
			default:
				log.Println("Error accepting connection:", err)
			}
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr())
		wg.Add(1)
		go handleConnection(conn)
	}
}

func main() {
	port := "9000"

	// Обработка системных сигналов для graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go startServer(port)

	<-sigs
	fmt.Println("\nShutting down server...")

	// Остановка всех новых подключений и завершение активных
	close(stop)
	wg.Wait()
	fmt.Println("Server gracefully stopped")
}
