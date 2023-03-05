package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/nohns/semesterprojekt2/pkg/event"
	"github.com/nohns/semesterprojekt2/pkg/eventsource"
	"github.com/nohns/semesterprojekt2/pkg/sqlite"
	"github.com/nohns/semesterprojekt2/proto/gen/go/pi/events/v1"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Open database
	db, err := sql.Open("sqlite3", "file:bridge.db")
	if err != nil {
		log.Printf("error opening database: %v", err)
		return
	}
	if err := sqlite.Migrate(db); err != nil {
		log.Printf("error migrating database: %v", err)
		return
	}
	defer db.Close()

	// Create event store
	var store eventsource.EventStore = sqlite.NewEventSource(db)

	// Try store random DoorUnlocked event
	evt, err := event.FromMessage("door2", &events.PhoneAuthorized{
		UserId: "martin",
	})
	if err != nil {
		log.Printf("error creating event: %v", err)
		return
	}
	if err := store.Put(context.Background(), evt); err != nil {
		log.Printf("error persisting event: %v", err)
		return
	}
	log.Printf("event persisted. Got version %d at %v", evt.Version, evt.At)

	// Try to read all events
	cursor, err := store.Play(context.Background())
	if err != nil {
		log.Printf("error reading events: %v", err)
		return
	}
	defer cursor.Close()

	for cursor.Next() {
		evt, err := cursor.Event()
		if err != nil {
			log.Printf("error reading event: %v", err)
			return
		}
		log.Printf("event: %v (id = %s), version: %d (aggr. ver = %d), at: %v", evt.Type, evt.AggregateId, evt.Version, evt.AggregateVersion, evt.At)

		switch evt.Type {
		case "PhoneAuthorized":
			var msg events.PhoneAuthorized
			if err := evt.Unmarshal(&msg); err != nil {
				log.Printf("error unmarshaling event data: %v", err)
				return
			}
			log.Printf(" -> for user: %s", msg.UserId)
		}
	}

}
