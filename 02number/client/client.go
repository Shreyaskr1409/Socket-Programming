package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	// dialing will be done to localhost:8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to the server: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to the server successfully!")

	// buffer to store incoming data
	buf := make([]byte, 4) // 4 bytes is the size of int32

	// the following is to read data from the server
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from the server: ", err)
		return
	}

	var number int32
	// the buffer will be read knowing that it is coded in littleEndian byte order
	// the read data will be stored in number's value
	err = binary.Read(bytes.NewReader(buf), binary.LittleEndian, &number)
	if err != nil {
		fmt.Println("Error while reading the incoming binary data: ", err)
		return
	}

	fmt.Println("Recieved number from the server: ", number)
}
