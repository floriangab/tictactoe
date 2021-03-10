APP=Tictactoe

SRC=./src/
SRC_FILES=$(wildcard $(SRC)*.go)

.PHONY: build
build: clean
		go build -o ${APP} ${SRC_FILES}

.PHONY: run
run:
		go run -race $(SRC_FILES)

.PHONY: clean
clean:
		go clean

fclean: clean
		rm -rf ${APP}
		rm -rf audio
		rm -f *.wav