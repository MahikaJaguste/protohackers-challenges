package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var address = "127.0.0.1:80"

func testConnection(threadId int) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Thread %d: %s\n", threadId, conn.LocalAddr().String())

	data := "Hello world"
	n, err := conn.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Thread %d: %d bytes written\n", threadId, n)

	dataReturned := make([]byte, 100)
	n, err = conn.Read(dataReturned)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Thread %d: %d bytes read: %s\n", threadId, n, string(dataReturned[:]))
}

func main() {
	fmt.Println("Testing echo server")
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			testConnection(i)
		}()
	}
	wg.Wait()
}
