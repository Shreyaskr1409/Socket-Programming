package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type User struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	IsStudent bool     `json:"isStudent"`
	Courses   []string `json:"courses"`
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error encountered while starting the server: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on the port 8080!")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error while accepting the connection: ", err)
			continue
		}
		fmt.Println("Client connected!")

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	user := User{
		Name:      "Shreyas",
		Age:       18,
		IsStudent: true,
		Courses:   []string{"Analog Electronics", "Digital Electronics"},
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error converting data into json")
		return
	}

	fmt.Println(string(jsonData))

	fmt.Println("----------------------------------")

	n, err := conn.Write(jsonData)
	if err != nil {
		fmt.Println("Error while establishing the connection: ", err)
		return
	}
	fmt.Println("Data sent successfully with ", n, " bytes sent")
}
