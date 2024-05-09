package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":227")
	if err != nil {
		log.Println("couldn't bind the port")
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("couldn't establish the connection")
			return
		}

		go handle(c)
	}
}

func handle(c net.Conn) {
	defer c.Close()

	dst, err := net.Dial("tcp", "localhost:8082")
	if err != nil {
		log.Println("couldn't reach the destination host")
		return
	}

	go func() {
		if _, err := io.Copy(c, dst); err != nil {
			log.Println("couldn't send data to the client")
		}
	}()

	if _, err := io.Copy(dst, c); err != nil {
		log.Println("couldn't send data to the destination host")
	}
}

func handleBuffered(c net.Conn) {
	defer c.Close()

}
