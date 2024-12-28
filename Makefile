run:
	@templ generate
	@sqlc generate -f model/sqlc.yaml
	@rm data.sqlite3
