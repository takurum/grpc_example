

Установить пакет:
- `go get -u github.com/golang/protobuf/protoc-gen-go`

Сгенерить из proto:
- `protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto`

---

Для запуска в корне проекта выполнить команды:
- `go run cmd/server/main.go`
- `go run cmd/client/main.go 2 2`

При запуске клиента передать слагаемые аргументы
  
---

Для тестирования сервера можно использовать Evans. Скачать и распаковать https://github.com/ktr0731/evans/releases "Source code".
Для тестирования необходимо:
- в корне этого проекте выполнить `go run cmd/server/main.go` 
- в корне проекта Evans выполнить `go run main.go /home/ff/go/src/grpc_my_example/api/proto/adder.proto --host localhost -p 8080`
