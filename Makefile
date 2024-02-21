.PHONE: run
run:
	go run cmd/bad/main.go

.PHONY: build
build:
	go build -o bad cmd/bad/main.go