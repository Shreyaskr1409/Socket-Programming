package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	IsStudent bool     `json:"isStudent"`
	Courses   []string `json:"courses"`
}

func main() {
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
}
