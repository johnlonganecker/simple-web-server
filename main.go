package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go HandleRequest(conn)
	}

}

func HandleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	startLine := "HTTP/1.1 200 OK"
	entityBody := "Hello, World"
	headers := Header("Content-Length", "5")

	response := startLine + "\n" + headers + "\n\n" + entityBody

	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Printf("%s", buf)

	conn.Write([]byte(response))

	conn.Close()
}

func Header(key string, value string) string {
	return key + ": " + value
}
