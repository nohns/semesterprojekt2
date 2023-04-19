package sqlite

import (
	"database/sql"
)

type sqlite struct {
	db *sql.DB
}

func New() *sqlite {

	/* 	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s", conf.DBPath))
	   	if err != nil {
	   		log.Printf("error opening database: %v", err)
	   		return
	   	}
	   	if err := sqlite.Migrate(db); err != nil {
	   		log.Printf("error migrating database: %v", err)
	   		return
	   	} */

	// Create event store
	//var store eventsource.EventStore = sqlite.NewEventSource(db)

	//defer db.Close()

	return &sqlite{}
}
