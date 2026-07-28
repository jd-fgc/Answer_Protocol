GO=go

SRC = \
		server/look.go \
		server/models.go \
		server/move.go \
		server/players_charac.go \
		server/talk.go \
		server/user_input_loop.go

APP_EXECUTABLE=Answer_Protocol

build:
	GOARCH=amd64 GOOS=linux go build -o ${APP_EXECUTABLE} server/main.go ${SRC}

run: build
	./${APP_EXECUTABLE}

clean:
	go clean
	rm ${APP_EXECUTABLE}