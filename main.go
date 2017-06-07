package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"./parser"
	"./response"
	//"github.com/johnlonganecker/web-server/parser"
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

	resp := response.Response{
		Version:      "HTTP/1.1",
		Status:       "200",
		ReasonPhrase: "OK",
		Body:         "Hello, World",
		Headers:      "",
	}

	//t := time.Now()
	//date := fmt.Printf("%d, %d ", t.Weekday().t.Day())
	//t.Format("Mon, 02 Jan 2017 12:00:00 GMT")

	resp.AddHeader("Content-Length", "12")
	resp.AddHeader("Content-Type", "text/plain; charset=iso-latin-1")
	//resp.AddHeader("Date", date)

	scanner := bufio.NewScanner(conn)
	parser.Parse(scanner)

	conn.Write([]byte(resp.String()))

	err := conn.Close()
	if err != nil {
		fmt.Println(err)
	}
}
