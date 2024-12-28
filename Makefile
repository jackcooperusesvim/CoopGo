run:
	@templ generate
	@sqlc generate -f model/sqlc.yaml
	go run cmd/main.go
