package sqlite

import "database/sql"

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS events (
			aggregate_id TEXT,
			version INTEGER NOT NULL UNIQUE,
			aggregate_version INTEGER,
			type TEXT NOT NULL,
			at DATETIME DEFAULT CURRENT_TIMESTAMP,
			data BLOB NOT NULL
		);
		CREATE INDEX IF NOT EXISTS events_version ON events (version);
	`)
	if err != nil {
		return err
	}

	return nil
}
