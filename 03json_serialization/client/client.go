package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type User struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	IsStudent bool     `json:"isStudent"`
	Courses   []string `json:"courses"`
}

func main() {
	jsondata := `{"name":"Shreyas","age":18,"isStudent":true,"courses":["Analog Electronics", "Digital Electronics"]}`

	// Method 1
	var user User
	err := json.Unmarshal([]byte(jsondata), &user)
	if err != nil {
		fmt.Println("Error occured while deserializing json data")
		return
	}
	fmt.Println("Method 1")
	fmt.Printf("Name: %s, Age: %d, IsStudent: %t, Courses: %v\n", user.Name, user.Age, user.IsStudent, user.Courses)

	println("-----------------------------------")

	// Method 2
	var data map[string]interface{}
	// using interface, we can store values when we do not necessarrily know what data we are about to get
	err = json.Unmarshal([]byte(jsondata), &data)
	if err != nil {
		fmt.Println("Error occured while deserializing json data")
		return
	}
	fmt.Println("Name: ", data["name"])
	fmt.Println("Age: ", data["age"])
	fmt.Println("IsStudent: ", data["isStudent"])
	fmt.Println("Courses: ", data["courses"])

	println("-----------------------------------")

	Main_2()

	println("-----------------------------------")

	Main_3()
}

func Main_2() {
	// JSON data representing an array of users
	jsonData := `[
		{"name": "Alice", "age": 25, "isStudent": false},
		{"name": "Bob", "age": 22, "isStudent": true}
	]`

	var users []map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &users)
	if err != nil {
		fmt.Println("Error deserializing JSON:", err)
		return
	}

	// Loop through the slice of users
	for _, user := range users {
		fmt.Printf("Name: %s, Age: %v, IsStudent: %v\n", user["name"], user["age"], user["isStudent"])
	}
}

func Main_3() {
	conn, err := net.Dial("tcp", "localhost:8080")
	// could have used :8080 but it would dial to any port with number 8080 on the machine
	if err != nil {
		println("Connection established")
	}
	defer conn.Close()

	fmt.Println("Connection established")

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)

	if err != nil {
		if err == io.EOF {
			fmt.Println("Connection closed by the server")
		} else {
			fmt.Println("Error reading the data: ", err)
		}
		return
	}

	jsonData := buffer[:n]

	fmt.Println("Recieved JSON data: ", string(jsonData))

	var user User
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		fmt.Println("Error while unmarshalling the json data: ", err)
		return
	}

	fmt.Println("Recieved user: \n", user)
}
