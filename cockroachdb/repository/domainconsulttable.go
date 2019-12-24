package repository

import (
  "database/sql"
  _ "github.com/lib/pq"
  "entities"
  "log"
  "time"
)

func CreateDomainConsult( db *sql.DB,  model entities.DomainModel, id_domain int ) ( bool, error ) {
  log.Println( "Query to Create DomainConsult row" )
  // Execute queries
  _, err := db.Exec( "INSERT INTO DOMAINCONSULT(id_domain, consult_time, ssl_grade, title, logo, is_down) VALUES($1,$2,$3,$4,$5,$6)", id_domain, time.Now(), model.SslGrade, model.Title, model.Logo, model.Is_down );

  if err != nil {
    return false, err
  }
  return true, nil
}
