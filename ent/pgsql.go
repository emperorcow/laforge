package ent

import (
	"database/sql"
	"log"

	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib" //
)

// PGOpen Open new PostGres connection
func PGOpen(databaseURL string) *Client {
    db, err := sql.Open("pgx", databaseURL)
    if err != nil {
        log.Fatal(err)
    }

    // Create an ent.Driver from `db`.
    drv := entsql.OpenDB(dialect.Postgres, db)
    return NewClient(Driver(drv))
}
