package main

import (
	"fmt"
	"net"
)

func main() {
	ports := make(chan int, 100)
	res := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, res)
	}

	go func() {
		for i := 0; i < 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-res
		if port != 0 {
			fmt.Printf("Port %d is open.\n", port)
		}
	}

	close(ports)
	close(res)
}

func worker(ports <-chan int, res chan<- int) {
	for p := range ports {
		c, err := net.Dial("tcp", "127.0.0.1:"+fmt.Sprintf("%d", p))
		if err != nil {
			res <- 0
			continue
		}

		c.Close()
		res <- p
	}
}
