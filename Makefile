run:
	@templ generate
	@sqlc generate -f model/sqlc.yaml
	@air
	@rm data.sqlite3
