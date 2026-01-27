package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// initialize and return a new database connection
func IntiDb() (*sql.DB, error) {
	dsn := "file:" + "test.db" + "?cache=shared&_pragma=journal_mode(WAL)&_pragma=foreign_keys(ON)&_pragma=busy_timeout(5000)&_pragma=synchronous(NORMAL)"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	// ping the database to ensure connection is established
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	if err := db.PingContext(ctx); err != nil {
		if cErr := db.Close(); cErr != nil {
			return nil, errors.Join(err, cErr)
		}
		return nil, err
	}

	return db, nil
}

func (s *Server) CloseDb() error {
	fmt.Println("closing database connection")
	if err := s.db.Close(); err != nil {
		return err
	}

	return nil
}
