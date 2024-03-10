.PHONY:  run stop  wire build swag 

build: ${BINARY_DIR} ## Compile the code, build Executable File
    $(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

run: ## Start application
    $(GOCMD) run ./cmd/api

wire: ## Generate wire_gen.go
    cd pkg/di && wire

swag: ## Generate swagger docs
    swag init -g pkg/api/handler/admin.go -o ./cmd/api/docs

run :
	docker compose up
stop:
	docker compose down
