package certificatestore

import "database/sql"

func Migrate(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS certificate (id INTEGER PRIMARY KEY, certificate BLOB, privatekey BLOB)")
	if err != nil {
		return err
	}
	return nil

}
