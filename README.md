# Тестовый проект с использованием protobuf
По материалам https://www.youtube.com/watch?v=xQRr2cO8Pg0

Преимущества protobuf:
 - скорость (это двоичный формат)
 - эффективность
 - безопасность схемы и типов
 - прямая и обратная совместимость

## Установка 

В системе должен быть установлен [компилятор `protoc`](https://github.com/protocolbuffers/protobuf/releases). 

Для установки пакета `protobuf` для Go выполнить команду:
```bash
go get google.golang.org/protobuf@latest 
```

Установка плагина для кодогенерации из определений:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## Кодогенерация

Кодогенерация из proto-файла, результат будет в файле `./models/person.pb.go`:
```bash
protoc --go_out=. --go_opt=paths=source_relative models/person.proto
```

## Проверка работоспособности

Сперва маршалим объект, который пошлем на http-сервер:
```bash
go run person_marshall.go
```
В результате будет создан файл `./tmp/person.bin`

Запуск сервера:
```bash
go run main.go
```

Отправка запроса на сервер (в командной строке bash):
```bash
curl -X POST --data-binary @tmp/person.bin http://localhost:8080/add
```
В ответ должна прийти строка:
```
Person added: name:"John Doe"  id:1234  email:"johndoe@mail.ru"  phones:{number:"123-456-7891"}
```


Получение добавленных данных:
```bash
curl -X GET http://localhost:8080/get?id=1234 --output tmp/retrieved_person.bin
```
Результатом выполнения GET-запроса станет создание файла `./tmp/retrieved_person.bin`

Для проверки полученного с сервера результат осуществим анмаршалинг принятого файла:
```bash
go run person_unmarshall.go
```
В результате должна быть строка:
```
Результат десериализации: name:"John Doe"  id:1234  email:"johndoe@mail.ru"  phones:{number:"123-456-7891"}
```