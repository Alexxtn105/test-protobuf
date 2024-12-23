package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
	"strconv"
	"test-protobuf/models"
)

var persons = make(map[int32]*models.Person)

func main() {
	http.HandleFunc("/add", addPersonHandler)
	http.HandleFunc("/get", getPersonHandler)

	fmt.Println("Сервер запущен на порту 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func addPersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Поддерживается только метод POST", http.StatusMethodNotAllowed)
		return
	}

	// читаем тело запроса, которое в формате protobuf
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Сбой чтения тела запроса", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// анмаршалим из protobuf в нужную нам структуру типа models.Person
	person := models.Person{}
	if err = proto.Unmarshal(body, &person); err != nil {
		http.Error(w, "Сбой анмаршалинга protobuf", http.StatusBadRequest)
		return
	}

	// добавляем структуру в мапу
	persons[person.Id] = &person

	_, err = fmt.Fprintf(w, "Person added: %v", &person)
	if err != nil {
		http.Error(w, "не удалось отформатировать строку выводы", http.StatusInternalServerError)
		return
	}
}

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Поддерживается только метод GET", http.StatusMethodNotAllowed)
		return
	}

	// извлекаем ИД из тела запроса
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Отсутствует параметр id", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	id := int32(idInt)

	person, ok := persons[id]
	if !ok {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
	//if person == nil {
	//	http.Error(w, "Person not found", http.StatusNotFound)
	//	return
	//}

	// маршалим найденное в формат protobuf
	data, err := proto.Marshal(person)
	if err != nil {
		http.Error(w, "Не удалось осуществить маршалинг Protobuf", http.StatusInternalServerError)
		return

	}
	// устанавливаем заголовок ответа в тип "octet-stream"
	w.Header().Set("Content-Type", "application/octet-stream")

	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Не удалось записать данные в ответ", http.StatusInternalServerError)
		return
	}

}
