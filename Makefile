BINFILE=pwdify

build:
	go build -o ${BINFILE} cmd/pwdify/*

run: pwdify
	@DEBUG=1 go run cmd/pwdify/* $(ARGS)

test:
	go test -v ./...

.PHONY: build run test