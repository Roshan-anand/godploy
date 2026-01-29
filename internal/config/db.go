package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Roshan-anand/godploy/internal/db"
	migration "github.com/Roshan-anand/godploy/sqlite"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

const (
	MAX_DB_OPEN_CONNECTIONS = 1
	MAX_DB_IDLE_CONNECTIONS = 1
	PING_TIMEOUT            = 5
)

var Pool_Close_Err = fmt.Errorf("DB pool close err")

// for migrating the database
func MigrateDb(db *sql.DB) error {
	mFs, err := migration.GetMigrationFS()
	if err != nil {
		return err
	}

	source, err := iofs.New(mFs, ".")
	if err != nil {
		return err
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		source,
		"sqlite3",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

// initialize and return a new database connection
func IntiDb() (*DataBase, error) {
	// TODO : replace path with config value
	dsn := "file:" + "test.db" +
		"?_pragma=journal_mode(WAL)" +
		"&_pragma=foreign_keys(ON)" +
		"&_pragma=busy_timeout(5000)" +
		"&_pragma=synchronous(NORMAL)"
	pool, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	pool.SetMaxOpenConns(MAX_DB_OPEN_CONNECTIONS)
	pool.SetMaxIdleConns(MAX_DB_IDLE_CONNECTIONS)

	// run migrations
	if err := MigrateDb(pool); err != nil {
		if cErr := pool.Close(); cErr != nil {
			return nil, errors.Join(err, cErr)
		}
		return nil, fmt.Errorf("Migration error : %w", err)
	}

	// ping the database to ensure connection is established
	ctx, cancle := context.WithTimeout(context.Background(), PING_TIMEOUT*time.Second)
	defer cancle()

	if err := pool.PingContext(ctx); err != nil {
		if cErr := pool.Close(); cErr != nil {
			return nil, errors.Join(Pool_Close_Err, err, cErr)
		}
		return nil, errors.Join(Pool_Close_Err, err)
	}

	queries := db.New(pool) // get query instance from sqlc generated code

	fmt.Println("database connection established ...") //TODO : replace with proper logging
	return &DataBase{
		Pool:    pool,
		Queries: queries,
	}, nil
}

// close the database connection
func (s *Server) CloseDb() error {
	fmt.Println("closing database connection")
	if err := s.DB.Pool.Close(); err != nil {
		return errors.Join(Pool_Close_Err, err)
	}

	return nil
}
