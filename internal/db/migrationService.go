package db

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"
	"sync"
)

type MigrationService interface {
	Up(path string) error
	Down(path string) error
}

var once sync.Once
var instance MigrationService

type migrationService struct {
	db *sql.DB
}

func (m migrationService) Up(path string) error {
	err := goose.SetDialect("postgres")

	if err != nil {
		return fmt.Errorf("can't migtate: %v", err)
	}

	if err := goose.Up(m.db, path); err != nil {
		log.Fatalf("Ошибка применения миграций: %v", err)
	}

	return err
}

func (m migrationService) Down(path string) error {
	err := goose.SetDialect("postgres")

	if err != nil {
		return fmt.Errorf("can't migtate: %v", err)
	}

	if err := goose.Down(m.db, path); err != nil {
		log.Fatalf("Ошибка применения миграций: %v", err)
	}

	return err
}

func NewMigrationService(db *sql.DB) MigrationService {
	once.Do(func() {
		instance = &migrationService{
			db: db,
		}
	})

	return instance
}
