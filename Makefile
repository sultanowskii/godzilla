NAME = godzilla

.PHONY: clean build install

default: build

clean:
	go clean

build:
	go build -o $(NAME) ./main.go

install:
	go install
