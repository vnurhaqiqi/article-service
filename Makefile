run:
	go run cmd/server/main.go

wire:
	wire gen internal/app/wire.go

swagger:
	swag init -g cmd/server/main.go

.PHONY: run wire swagger