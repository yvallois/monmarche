package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"log"
	"test_ticket/common/database/migrations"
)

func RunMigration() {
	m := gormigrate.New(DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		&migrations.M20221205120000,
	})
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
