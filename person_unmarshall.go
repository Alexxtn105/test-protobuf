package main

import (
	"fmt"
	"log"
	"os"

	"test-protobuf/models"

	"google.golang.org/protobuf/proto"
)

func main() {
	data, err := os.ReadFile("tmp/retrieved_person.bin")
	if err != nil {
		log.Fatalf("Не удалось проистать файл: %v", err)
	}

	person := &models.Person{}
	err = proto.Unmarshal(data, person)
	if err != nil {
		log.Fatalf("Сбой анмаршалинга: %v", err)
	}

	fmt.Println("Результат десериализации:", person)
}
