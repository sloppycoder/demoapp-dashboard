build:
	go build -o bin/server cmd/server.go

test: 
	go test -coverprofile=coverage.out

all: protoc build test

clean:
	rm -f bin/server
	rm -f coverage.out
