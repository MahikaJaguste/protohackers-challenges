package main

import (
	"fmt"
	"log"
	"net"
)

var CHUNK_SIZE = 1024
var port = "80"

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Connection received from %s\n", conn.RemoteAddr().String())

	data := make([]byte, 0, CHUNK_SIZE)
	dataChunk := make([]byte, CHUNK_SIZE)
	n := 0
	for {
		_n, err := conn.Read(dataChunk)
		if err != nil {
			log.Fatal(err)
		}
		n += _n
		data = append(data, dataChunk...)
		if _n < CHUNK_SIZE {
			break
		}
	}

	data = data[:n]
	fmt.Printf("%d bytes read\n", n)

	n, err := conn.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes written\n", n)

}

func main() {

	fmt.Println("Server starting ...")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server running !!!")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
