# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN := $(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH ="$$PATH:$(LOCAL_BIN)" protoc

# Путь до GOBIN
GOBIN := C:\Users\User\Development\sdk\go\1.22.5\bin

# Путь до protobuf файлов
PROTO_PATH := $(CURDIR)/api

# Путь до сгенерированных .pb.go файлов
PKG_PROTO_PATH :=$(CURDIR)/pkg

# Устанавливаем необходимые файлы
.bin-deps: export GOBIN := C:\Users\User\Development\sdk\go\1.22.5\bin
.bin-deps:
	$(info Installing binary dependencies...)

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Генерация .go файлов с помощью protoc
.protoc-generate:
	protoc --proto_path=. \
	--go_out=./pkg --go_opt paths=source_relative \
	--go-grpc_out=./pkg --go-grpc_opt paths=source_relative \
	./api/notes/service.proto \
	./api/notes/messages.proto

# go mod tidy
.tidy: 
	go mod tidy

# Генерация кода из protobuf
generate: .bin-deps .protoc-generate .tidy

# Билд приложения
build:
	go build -o $(LOCAL_BIN) ./cmd/notes/client
	go build -o $(LOCAL_BIN) ./cmd/notes/server