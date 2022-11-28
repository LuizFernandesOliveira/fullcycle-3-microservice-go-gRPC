# Full Cycle 3.0
## Comunicação entre microservices - gRPC
Este projeto foi criado durante o curso de comunicações entre sistemas focado no gRPC

#### Processa o proto
```shell
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

#### Starting microsservice
```shell
go run cmd/grpcServer/main.go
```

#### Usando evans para simular as requests
```shell
evans -r repl
```