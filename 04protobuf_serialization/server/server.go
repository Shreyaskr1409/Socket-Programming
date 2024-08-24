package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"net"
	"server/protobuf_data"
)

// type User struct {
// 	Name      string   `json:"name"`
// 	Age       int      `json:"age"`
// 	IsStudent bool     `json:"isStudent"`
// 	Courses   []string `json:"courses"`
// }

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error occured while starting the server: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on the port 8080!")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error occured while accepting the error: ", err)
			continue
		}
		fmt.Println("Client connected!")

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	user := &protobuf_data.User{
		Name:      "Shreyas",
		Age:       18,
		IsStudent: true,
		Courses:   []string{"Analog Electronics", "Digital Electronics"},
	}

	protoData, err := proto.Marshal(user)
	if err != nil {
		fmt.Println("Error converting data into protobuf")
		return
	}

	n, err := conn.Write(protoData)
	if err != nil {
		fmt.Println("Failed to send data: ", err)
		return
	}

	fmt.Println(n, " bytes of Protobuf data sent!")
	// will continue work on this after setting up protobuf
}
