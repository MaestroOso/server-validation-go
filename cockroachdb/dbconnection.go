package cockroachdb

import (
  "database/sql"
  _ "github.com/lib/pq"
  "log"
)

func OpenConnection( ) ( *sql.DB, error ) {
    db, err := sql.Open( "postgres", "postgresql://" + DbUser + "@" + DbUrl + DbPort + "/" + DbDatabaseName + "?sslmode=disable" )
    log.Println("Opening connection using postgresql://" + DbUser + "@" + DbUrl + DbPort + "/" + DbDatabaseName + "?sslmode=disable")
    return db, err
}
