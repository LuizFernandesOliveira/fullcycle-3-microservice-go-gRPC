# Full Cycle 3.0
## Comunicação entre microservices - gRPC
Este projeto foi criado durante o curso de comunicações entre micro
services focado no gRPC

#### Processa o proto
```shell
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

#### Usando evans para simular as requests
```shell
evans -r repl
```