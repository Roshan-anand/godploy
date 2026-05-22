package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Roshan-anand/godploy/internal/config"
	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
)

// get temp dir for testing
func getTempDir() (string, error) {
	p, err := os.MkdirTemp("", "godploy_test_*")
	if err != nil {
		return "", err
	}

	badgerDir := fmt.Sprintf("%s/badger", p)
	if err := os.Mkdir(badgerDir, os.FileMode(0755)); err != nil {
		return "", err
	}

	return badgerDir, nil
}

func AddLogs(db *config.BadgerDB, dID uuid.UUID, logs []string) {
	txn := db.Pool.NewTransaction(true)

	for i, log := range logs {
		key := fmt.Sprintf("%s_%06d", dID.String(), i)
		if err := txn.Set([]byte(key), []byte(log)); err == badger.ErrTxnTooBig {
			_ = txn.Commit()
			txn = db.Pool.NewTransaction(true)
			_ = txn.Set([]byte(key), []byte(log))
		}
	}
	_ = txn.Commit()
}

// get all logs of a deployment by deployment id
func StreamAllLogsByDeploymentID(db *config.BadgerDB, dID uuid.UUID) ([]string, error) {
	prefix := []byte(dID.String() + "_")

	var logs []string

	err := db.Pool.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		opts.Prefix = prefix

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			// k := item.Key()
			if err := item.Value(func(val []byte) error {
				logs = append(logs, string(val))
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	})

	return logs, err
}

func main() {

	badgerTempPath, err := getTempDir()

	badger, err := config.InitBadgerDB(badgerTempPath)
	if err != nil {
		log.Fatalf("failed to initialize badger db: %v", err)
	}

	logs := []string{}
	for i := range 100 {
		logs = append(logs, fmt.Sprintf("log %d", i))
	}

	dID := uuid.New()
	AddLogs(badger, dID, logs)
	streamedLogs, err := StreamAllLogsByDeploymentID(badger, dID)
	if err != nil {
		log.Fatalf("failed to stream logs: %v", err)
	}

	fmt.Printf("add logs :\n")
	for _, log := range logs {
		fmt.Printf("%s\n", log)
	}

	fmt.Printf("----------------------------\n")
	fmt.Printf("streamed logs :\n")
	for _, log := range streamedLogs {
		fmt.Printf("%s\n", log)
	}
}
