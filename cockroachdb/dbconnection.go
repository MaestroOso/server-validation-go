package cockroachdb

import (
  "database/sql"
  _ "github.com/lib/pq"
)

func OpenConnection( ) ( *sql.DB, error ) {
    db, err := sql.Open( "postgres", "postgresql://" + DbUser + "@" + DbUrl + DbPort + "/" + DbDatabaseName + "?sslmode=disable" )
    return db, err
}

func GetDomains( db *sql.DB ) ( []DomainDbModel, error ) {
    // Execute queries
    rows, err := db.Query("SELECT * FROM Domain")

    if err != nil {
        return []DomainDbModel{}, err
    }

    defer rows.Close()
    var domains []DomainDbModel

    for rows.Next() {
        var domain DomainDbModel
        dataerror := rows.Scan(&domain)

        if dataerror != nil {
            return []DomainDbModel{}, err
        }

        domains = append(domains, domain)
    }

    return domains, nil
}
