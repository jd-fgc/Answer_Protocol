SRC = \
	combat.go \
	item_usage.go \
	look.go \
	move.go \
	players_charac.go \
	quests.go \
	talk.go \
	user_input_loop.go \

NAME = Client


build:
	go build -o $(NAME) main.go $(SRC)