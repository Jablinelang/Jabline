BINARY_NAME=jabline
INSTALL_PATH=/usr/local/bin
MAIN_PATH=./cmd/jabline

.PHONY: all build install clean

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

install: build
	sudo mv $(BINARY_NAME) $(INSTALL_PATH)

clean:
	rm -f $(BINARY_NAME)

all: install
