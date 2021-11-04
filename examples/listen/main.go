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

	chosenPort := strings.Split(l1.Addr().String(), ":")[1]

	log.Printf("Chosen port number is %s by listener 1", chosenPort)

	// Listen on same port number
	l2, err := reuseable.Listen("tcp", "127.0.0.1:"+chosenPort)
	if err != nil {
		log.Fatalf("unable to start listener 2: %v", err)
	}

	// If err not exists. We have two listener on the same port.
	log.Printf("Listener 1 Address: %s\n", l1.Addr().String())
	log.Printf("Listener 2 Address: %s\n", l2.Addr().String())

	l1.Close()
	l2.Close()
}
