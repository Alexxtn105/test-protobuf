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

Кодогенерация из proto-файла, результат будет в файле `./models/person.pb.go`:
```bash
protoc --go_out=. --go_opt=paths=source_relative models/person.proto
```