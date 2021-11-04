package main

import (
	"log"
	"strings"

	"github.com/projecthunt/reuseable"
)

func main() {
	l1, err := reuseable.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("unable to start listener 1: %v", err)
	}

	// Listen different port number
	l2, err := reuseable.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("unable to start listener 2: %v", err)
	}

	listenerOneChosenPort := strings.Split(l1.Addr().String(), ":")[1]

	// Dial with Listener 2 using same port and ip address with Listener 1
	conn, err := reuseable.Dial("tcp", "127.0.0.1:"+listenerOneChosenPort, l2.Addr().String())
	if err != nil {
		log.Fatalf("unable to dial with listener 2: %v", err)
	}

	log.Printf("Listener 1 Address: %s\n", l1.Addr().String())
	log.Printf("Listener 2 Address: %s\n", l2.Addr().String())
	log.Printf("Conn Local Address: %s\n", conn.LocalAddr().String())
	log.Printf("Conn Remote Address: %s\n", conn.RemoteAddr().String())

	conn.Close()
	l1.Close()
	l2.Close()
}
