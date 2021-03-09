

Установить пакет:
- go get -u github.com/golang/protobuf/protoc-gen-go

Сгенерить из proto:
- protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto

---

Для запуска в корне проекта выполнить команды:
- go run cmd/server/main.go 
- go run cmd/client/main.go 2 2

При запуске клиента передать слагаемые аргументы для сложения
  
---

Для тестирования сервера можно использовать Evans. Скачать https://github.com/ktr0731/evans/releases распаковать "Source code".
Для тестирования сервера необходимо запустить его:
-  go run cmd/server/main.go 
В корне проекта Evans выполнить команду:
- go run main.go /home/ff/go/src/grpc_my_example/api/proto/adder.proto --host localhost -p 8080
