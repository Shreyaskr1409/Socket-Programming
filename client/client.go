package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to the server: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to the server.")

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading from the server; ", err)
		return
	}

	fmt.Println("Recieved from the server: ", string(buf[:n]))
}
