package main

import (
	"bytes"           // provides functionality for manipulating byte slices
	"encoding/binary" // used for encoding and decoding numbers in binary format, supporting different byte orders (order in which bytes are stored (Endian))
	"fmt"             //
	"net"             // provides portable interface for network input and output
)

func main() {
	// Listening on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error encountered while starting the server: ", err)
		return
	}
	defer listener.Close()
	// it is important to close the listener to release resources from memory

	fmt.Println("Server is listening on the port 8080!")

	// the following loop is to continuously listen till the connection is closed from connection handler function
	for {
		// listener.Accept() blocks [pauses execution till a certain event occurs]...
		// ...and makes further code to execute till client initiates a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error excepting the connection: ", err)
			continue
		}
		fmt.Println("Client connected!")

		// the following function will handle the listener's i/o...
		// ...and will close connection when required or asked by the user
		// the following is a goroutine i.e. it will enable server to serve multiple connection concurrently
		go handleConnection(conn)
	}
}

// Will serialise the number into bytes
func handleConnection(conn net.Conn) {
	// for this example we are going to close the connection at the end of sending of the number
	defer conn.Close()
	var number int32 = 562945

	// buffer will be created to hold the byte value of the number
	// The buffer acts as an in-memory byte slice (something like an ArrayList) that grows as needed.
	buf := new(bytes.Buffer)

	// writing number to the buffer using littleEndian byte order
	// byte order is the order in which the bytes are stored
	err := binary.Write(buf, binary.LittleEndian, number)

	if err != nil {
		fmt.Println("Error writing to the buffer: ", err)
		return
	}

	// binary representation of number will returned as slice of bytes using buf.Bytes()
	// conn.Write will send the argument over the connection
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("Error sending the data: ", err)
		return
	}
	fmt.Println("Sent number to client")
}
