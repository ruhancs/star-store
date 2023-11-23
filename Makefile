run: build
	@./store

build:
	@go build -o store ./cmd

migrate:
	migrate create -ext=cql -dir=sql/migrations -seq clients
	migrate create -ext=cql -dir=sql/migrations -seq create_products_table
	migrate create -ext=cql -dir=sql/migrations -seq create_carts_table
	migrate create -ext=cql -dir=sql/migrations -seq create_cart_items_table
	migrate create -ext=cql -dir=sql/migrations -seq create_transactions_table

migrate_up:
	migrate -path sql/migrations -database "cassandra://localhost:9042/store" up