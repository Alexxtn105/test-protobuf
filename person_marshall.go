package main

import (
	"log"
	"os"
	"test-protobuf/models"

	"google.golang.org/protobuf/proto"
)

func main() {
	person := &models.Person{
		Name:  "John Doe",
		Id:    1234,
		Email: "johndoe@mail.ru",
		Phones: []*models.PhoneNumber{
			{Number: "123-456-7891", Type: models.PhoneType_MOBILE},
		},
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Сбой маршалинга: %v", err)
	}

	os.WriteFile("tmp/person.bin", data, 0644)
}
