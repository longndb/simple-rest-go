dep:
	go mod download -x
build:
	go build -v cmd/main.go