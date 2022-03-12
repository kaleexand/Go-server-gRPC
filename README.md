# Go-server-gRPC
gRPC: мини сервер на Go


Сложение двух чисел

adder.proto - описание внешнего интерфейса
evans - утилита для подключения к серверу

Запуск:
1. Установка Protocol Buffers (можно использовать brew install procobuf) 

https://developers.google.com/protocol-buffers

3. Установка пакетов Go для работы с gRPC

go get -u google.golang.org/grpc

go get -u github.com/golang/protobuf/protoc-gen-go

https://grpc.io/docs/languages/go/quickstart/

3. protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto
4. Запуск сервера

go build -v ./cmd/server

./server

7. Во втором окне с помощью утилиты evans подключиться к серверу

evans api/proto/adder.proto -p 8080

Call Add

7.Введите x, y
