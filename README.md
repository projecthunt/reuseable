<p align="center">
    <img width="400" src="img/logo.png">
<p>

<p align="center">
	<a href="LICENSE">
		<img src="https://img.shields.io/badge/License-MIT-yellow.svg">
	</a>
	<a href="https://pkg.go.dev/github.com/projecthunt/reuseable">
		<img src="https://pkg.go.dev/badge/github.com/projecthunt/reuseable.svg" alt="Go Reference">
	</a>
	<a href="https://goreportcard.com/report/github.com/projecthunt/reuseable">
		<img src="https://goreportcard.com/badge/github.com/projecthunt/reuseable" >
	</a>
<p>

This library helps go developers to open sockets with SO_REUSEPORT and SO_REUSEADDR flags.

## Why ?
This flags will allow many processes to bind to the same port. In fact, any number of processes will be allowed to bind and the load will be spread across them.

With SO_REUSEPORT and SO_REUSEADDR each of the processes will have a separate socket descriptor. Therefore each will own a dedicated UDP or TCP receive buffer.

## Installation
```bash
go get github.com/projecthunt/reuseable
```

## Simple Example
```go
package main

import (
	"log"
	"strings"

	"github.com/projecthunt/reuseable"
)

func main() {
	l1, err := reuseable.Listen("tcp", "127.0.0.1:54651")
	if err != nil {
		log.Fatalf("unable to start listener 1: %v", err)
	}

	// Listen on same port number
	l2, err := reuseable.Listen("tcp", "127.0.0.1:54651")
	if err != nil {
		log.Fatalf("unable to start listener 2: %v", err)
	}

	// If err not exists. We have two listener on the same port and same ip.
	log.Printf("Listener 1 Address: %s\n", l1.Addr().String())
	log.Printf("Listener 2 Address: %s\n", l2.Addr().String())

	l1.Close()
	l2.Close()
}
```

The program output:
```bash
Listener 1 Address: 127.0.0.1:54651
Listener 2 Address: 127.0.0.1:54651
```

## LICENSE
This project is under [MIT License](LICENSE)