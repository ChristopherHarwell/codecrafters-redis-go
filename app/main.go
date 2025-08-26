package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	// conn, err := l.Accept()
	// if err != nil {
	// 	fmt.Println("Error accepting connection: ", err.Error())
	// 	os.Exit(1)
	// }

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue // handle next connection
		}
		go func(c net.Conn) {
			defer c.Close()
			// Handle the connection here, e.g., read/write loop
			// ...
			conn.Write([]byte("+PONG\r\n"))
			resp, err := conn.Write([]byte("+PONG\r\n"))

			if err != nil {
				fmt.Println("Error sending response", err.Error())
				os.Exit(1)
			}

			fmt.Println(resp)
		}(conn)
	}

}
