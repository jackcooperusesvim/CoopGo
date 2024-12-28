run:
	@templ generate
	@go run cmd/main.go
	@sqlc generate -f model/sqlc.yaml
