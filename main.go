package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Launching echo server on port " + port)
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Error(Listen)! ", err)
	}
	defer ln.Close()

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatal("Error(Accept)! ", err)
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	bufr := bufio.NewReader(c)
	buf := make([]byte, 1024)

	for {
		readBytes, err := bufr.Read(buf)
		if err != nil {
			c.Close()
			return
		}
		message := string(buf[:readBytes])
		log.Print("Received " + message)
		c.Write([]byte(message))
	}
}
