GO=go
PORT=9090
IP=10.12.9.1

build:
	go build -o bin/server ./server
	go build -o bin/client_cli ./client_cli
	go build -o bin/client_gui ./client_gui

run-server: build
	./bin/server $(PORT)

run-client: build
	./bin/client_cli $(IP) $(PORT)

run-client-gui: build
	./bin/client_gui $(IP) $(PORT)

clean:
	rm -rf bin/

lint:
	go vet ./...

.PHONY: build run-server run-client run-client-gui clean lint