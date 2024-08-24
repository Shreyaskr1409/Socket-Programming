package main

import (
	"encoding/json"
	"testing"
)

type User struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	IsStudent bool     `json:"isStudent"`
	Courses   []string `json:"courses"`
}

func BenchmarkJSONSerialization(b *testing.B) {
	user := User{
		Name:      "Alice",
		Age:       25,
		IsStudent: false,
		Courses:   []string{"Math", "Science"},
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(user)
	}
}

func BenchmarkJSONDeserialization(b *testing.B) {
	jsonData := []byte(`{"name":"Alice","age":25,"isStudent":false,"courses":["Math","Science"]}`)

	var user User

	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(jsonData, &user)
	}
}
