package db

func createDatabaseTables() {
	createEventString := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		created_at DATETIME NOT NULL,
		user_id INTEGER
	)
`
	_, err := DB.Exec(createEventString)

	if err != nil {
		panic("Failed to create event table")
	}
}
