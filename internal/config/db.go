package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Roshan-anand/godploy/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

// initialize and return a new database connection
func IntiDb() (*DataBase, error) {
	dsn := "file:" + "test.db" + "?cache=shared&_pragma=journal_mode(WAL)&_pragma=foreign_keys(ON)&_pragma=busy_timeout(5000)&_pragma=synchronous(NORMAL)"
	sqlDb, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxOpenConns(1)
	sqlDb.SetMaxIdleConns(1)

	// ping the database to ensure connection is established
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	if err := sqlDb.PingContext(ctx); err != nil {
		if cErr := sqlDb.Close(); cErr != nil {
			return nil, errors.Join(err, cErr)
		}
		return nil, err
	}

	queries := db.New(sqlDb) // get query instance from sqlc generated code

	fmt.Println("database connection established ...") //TODO : replace with proper logging
	return &DataBase{
		Pool:    sqlDb,
		Queries: queries,
	}, nil
}

func (s *Server) CloseDb() error {
	fmt.Println("closing database connection")
	if err := s.db.Pool.Close(); err != nil {
		return err
	}

	return nil
}
