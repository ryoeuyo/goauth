BIN_DIR := bin/goauthbin

build:
	go build -o $(BIN_DIR) ./cmd/goauth

run:
	GIN_MODE=release ./$(BIN_DIR)
