package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M20221205120000 = gormigrate.Migration{
	ID: "20221205120000",
	Migrate: func(tx *gorm.DB) error {
		err := tx.Exec(`
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
			
			CREATE TABLE tickets (
				uuid uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
				created_at timestamp without time zone NOT NULL,
				updated_at timestamp without time zone NOT NULL,
				order_id integer NOT NULL,
				vat numeric NOT NULL,
				total numeric NOT NULL
			);
				
			CREATE TABLE products (
				uuid uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
				created_at timestamp without time zone NOT NULL,
				updated_at timestamp without time zone NOT NULL,
				product_id character varying(64) NOT NULL,
				name character varying(128) NOT NULL,
				price numeric NOT NULL,
				ticket_uuid  uuid NOT NULL,
				CONSTRAINT fk_ticket_uuid FOREIGN KEY (ticket_uuid) REFERENCES tickets(uuid) ON DELETE CASCADE DEFERRABLE
			);
		`).Error
		return err
	},
	Rollback: func(tx *gorm.DB) error {
		err := tx.Exec(`
			DROP TABLE products;
			DROP TABLE tickets;
		`).Error
		return err
	},
}
