package main

import (
	"client/protobuf_data"
	"fmt"
	"google.golang.org/protobuf/proto"
	"net"
)

//type User struct {
//	Name      string   `json:"name"`
//	Age       int      `json:"age"`
//	IsStudent bool     `json:"isStudent"`
//	Courses   []string `json:"courses"`
//}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed to connect to the server: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to the server")

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Failed to read data from the connection: ", err)
	}

	var user protobuf_data.User
	protodata := buf[:n]
	err = proto.Unmarshal(protodata, &user)
	if err != nil {
		fmt.Println("Failed to unmarshal the protobuf data into current struct: ", err)
		return
	}

	// fmt.Println("Recieved user: ", user)
	fmt.Printf("Recieved user: %+v\n", user)

}
