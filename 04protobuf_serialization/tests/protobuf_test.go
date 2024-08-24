package main

import (
	"google.golang.org/protobuf/proto"
	"testing"
	"tests/protobuf_data"
)

func BenchmarkProtobufSerialization(b *testing.B) {
	user := &protobuf_data.User{
		Name:      "Shreyas",
		Age:       18,
		IsStudent: true,
		Courses:   []string{"Analog Electronics", "Digital Electronics"},
	}

	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(user)
	}
}

func BenchmarkProtobufDeserialization(b *testing.B) {
	protobufData := []byte(`[10 7 83 104 114 101 121 97 115 16 18 24 1 34 18 65 110 97 108 111 103 32 69 108 101 99 116 114 111 110 105 99 115 34 19 68 105 103 105 116 97 108 32 69 108 101 99 116 114 111 110 105 99 115]`)
	var user protobuf_data.User
	for i := 0; i < b.N; i++ {
		_ = proto.Unmarshal(protobufData, &user)
	}
}
