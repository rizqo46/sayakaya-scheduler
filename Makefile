create-migration: ## Create new migration file. It takes parameter `file` as filename. Usage: `make create-migration file=add_column_time`
	migrate create -ext sql -dir migrations -seq $(file)
	
up-migration:
	migrate -path migrations -database "$(db_url)" up