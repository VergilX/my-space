db := "dump.db"

init:
	@sqlite3 -line $(db) ".read ./sqlc/schema.sql"