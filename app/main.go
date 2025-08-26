package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err.Error())
	}
	for {
		_, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break // client closed connection
			}
			// handle error, possibly break
			break
		}

		// process netData, e.g., parse PING
		response := "+PONG\r\n"
		_, err = conn.Write([]byte(response))
		if err != nil {
			// error writing back, likely connection closed
			break
		}
	}

}
