pwdify:
	@go build -o pwdify cmd/pwdify/*

run: pwdify
	@./pwdify $(ARGS)

.PHONY: build