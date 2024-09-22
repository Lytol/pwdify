BINFILE=pwdify

build: build-template
	go build -o ${BINFILE} cmd/pwdify/*

run: pwdify
	@DEBUG=1 go run cmd/pwdify/* $(ARGS)

test:
	go test -v ./...

build-template:
	cd template && npm run build

.PHONY: build build-template run test