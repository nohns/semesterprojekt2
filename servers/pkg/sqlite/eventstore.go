package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/nohns/semesterprojekt2/pkg/event"
	"github.com/nohns/semesterprojekt2/pkg/eventsource"
)

type EventStore struct {
	db *sql.DB
}

func NewEventSource(db *sql.DB) *EventStore {
	return &EventStore{
		db: db,
	}
}

func (es *EventStore) Put(ctx context.Context, evts ...*event.Event) error {
	// Start DB transaction
	tx, err := es.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Put events
	for _, evt := range evts {
		if err := es.putTx(tx, evt); err != nil {
			return err
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (es *EventStore) putTx(tx *sql.Tx, event *event.Event) error {
	params := []any{}
	format := []any{}
	if event.AggregateId != "" {
		params = append(params, event.AggregateId)
		format = append(format, "(SELECT IFNULL(MAX(aggregate_version), 0) + 1 FROM events WHERE aggregate_id = ?)")
		params = append(params, event.AggregateId)
	} else {
		params = append(params, nil)
		format = append(format, "?")
		params = append(params, nil)
	}
	params = append(params, event.Type, event.Data)

	// The SELECT IFNULL subquery is a replacement for AUTO_INCREMENT for non primary id columns in SQLite.
	// Build SQL query
	sql := fmt.Sprintf(`
		INSERT INTO events (
			aggregate_id, 
			aggregate_version, 
			version, 
			type, 
			data
		) 
		VALUES (
			?, 
			%s, 
			(
				SELECT IFNULL(MAX(version), 0) + 1 
				FROM events
			), 
			?, 
			?
		)
		RETURNING aggregate_version, version, at;`, format...)

	// Prepare SQL statement
	stmt, err := tx.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert event into database, and return the given versions and timestamp
	if err := stmt.QueryRow(params...).Scan(&event.AggregateVersion, &event.Version, &event.At); err != nil {
		return err
	}

	return nil
}

func (es *EventStore) Get(ctx context.Context, aggregateId event.CompositeID) (eventsource.Cursor, error) {
	return es.Range(ctx, aggregateId, 0, 0)
}

func (es *EventStore) Range(ctx context.Context, aggregateId event.CompositeID, fromVersion, toVersion int) (eventsource.Cursor, error) {
	// Start DB transaction
	tx, err := es.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Only add version range constrains if they are set
	params := []interface{}{aggregateId}
	conditions := []string{"aggregate_id = ?"}
	if fromVersion > 0 {
		conditions = append(conditions, "version >= ?")
		params = append(params, fromVersion)
	}
	if toVersion > 0 {
		conditions = append(conditions, "version <= ?")
		params = append(params, toVersion)
	}
	// Build SQL query
	sql := fmt.Sprintf(`
		SELECT 
			aggregate_id, 
			version,
			aggregate_version, 
			type, 
			at,
			data 
		FROM events 
		WHERE 
			%s 
		ORDER BY version ASC`, strings.Join(conditions, " \nAND "))

	rows, err := tx.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	// Return cursor to iterate over events. Avoids loading all events into memory.
	return &evtSrcCursor{
		rows: rows,
	}, nil
}

func (es *EventStore) Play(ctx context.Context) (eventsource.Cursor, error) {
	// Start DB transaction
	tx, err := es.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query(`
		SELECT 
			aggregate_id, 
			version, 
			aggregate_version,
			type, 
			at,
			data 
		FROM events 
		ORDER BY version ASC`)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	// Return cursor to iterate over events. Avoids loading all events into memory.
	return &evtSrcCursor{
		rows: rows,
	}, nil
}

type evtSrcCursor struct {
	rows *sql.Rows
}

func (esc *evtSrcCursor) Next() bool {
	return esc.rows.Next()
}

func (esc *evtSrcCursor) Event() (*event.Event, error) {
	var aggregateId event.CompositeID
	var version int
	var aggregateVersion int
	var eventType string
	var data []byte
	var at time.Time
	err := esc.rows.Scan(&aggregateId, &version, &aggregateVersion, &eventType, &at, &data)
	if err != nil {
		return nil, err
	}

	return &event.Event{
		AggregateId:      aggregateId,
		Version:          version,
		AggregateVersion: aggregateVersion,
		Type:             eventType,
		At:               at,
		Data:             data,
	}, nil
}

func (esc *evtSrcCursor) Close() error {
	return esc.rows.Close()
}
