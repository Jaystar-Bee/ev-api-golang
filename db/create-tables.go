package db

func createDatabaseTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME	
	)
`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Failed to create user table")
	}

	createEventString := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		created_at DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
`
	_, err = DB.Exec(createEventString)

	if err != nil {
		panic("Failed to create event table")
	}

}
