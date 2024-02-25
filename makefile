wire: ## Generate wire_gen.go
	cd pkg/di && wire

	swag: ## Generate swagger docs
		swag init -g pkg/api/handler/admin.go -o ./cmd/api/docs