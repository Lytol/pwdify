BINFILE=pwdify

all: setup build

setup:
	go mod tidy && cd template && npm install

build: build-template
	go build -o ${BINFILE} cmd/pwdify/main.go

run:
	@DEBUG=1 go run cmd/pwdify/main.go $(ARGS)

test:
	go test -v ./...

build-template:
	cd template && npm run build

install: build
	sudo mv ${BINFILE} /usr/local/bin/

.PHONY: all setup build build-template run test install