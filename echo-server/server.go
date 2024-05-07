package main

import (
	"bufio"
	"log"
	"net"
	"unicode/utf8"
)

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Println("couldn't bind port")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("couldn't establish connection")
			return
		}

		go echo(conn)
	}
}

func echo(c net.Conn) {
	defer c.Close()

	r := bufio.NewReader(c)
	s, err := r.ReadString('\n')
	if err != nil {
		log.Println("couldn't read data from the client")
		return
	}

	log.Printf("Read %d bytes from the client: %s", utf8.RuneCountInString(s), s)

	log.Println("Writing data...")

	w := bufio.NewWriter(c)
	if _, err := w.WriteString(s); err != nil {
		log.Println("couldn't write data to the client")
		return
	}

	w.Flush()
}
