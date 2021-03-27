

install package:
- `go get -u github.com/golang/protobuf/protoc-gen-go`

generate from .proto:
- `protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto`

---

execute commands at the root of the project:
- `go run cmd/server/main.go`
- `go run cmd/client/main.go 2 2`

When starting the client, pass the summand arguments (a, b)
  
---

You can use application "Evans" for testing the server: (download and execute) https://github.com/ktr0731/evans/releases "Source code".
For testing, in the root of Evans:
- `go run cmd/server/main.go` 
- `go run main.go /home/ff/go/src/grpc_my_example/api/proto/adder.proto --host localhost -p 8080`
