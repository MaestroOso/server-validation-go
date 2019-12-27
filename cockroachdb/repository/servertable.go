package repository

import (
  "database/sql"
  _ "github.com/lib/pq"
  "entities"
  "log"
  "cockroachdb/models"
  "strconv"
)

func CreateServerConsult( db *sql.DB,  model entities.ServerInformationModel, id_domainconsult int ) ( bool, error ) {

  log.Println( "Query to Create Server row" )
  // Execute queries
  _, err := db.Exec( "INSERT INTO SERVER(id_domainconsult, address, ssl_grade, country, owner) VALUES($1,$2,$3,$4,$5)", id_domainconsult, model.Address, model.SslGrade, model.Country, model.Owner );

  if err != nil {
    return false, err
  }
  return true, nil
}

func SearchServerByDomainConsultId( db *sql.DB, id_domainconsult int ) ( []models.ServerDbModel , error ) {
  log.Println( "Query to Search Server row" )

  // Execute queries
  rows, err := db.Query("SELECT id, id_domainconsult, address, ssl_grade, country, owner FROM SERVER where id_domainconsult = " + strconv.Itoa( id_domainconsult ) )

  if err != nil {
      return []models.ServerDbModel{}, err
  }

  defer rows.Close()
  var servers []models.ServerDbModel

  for rows.Next() {
      var id, id_domainconsult int
      var address, ssl_grade, country, owner string
      dataerror := rows.Scan( &id, &id_domainconsult, &address, &ssl_grade, &country, &owner )

      if dataerror != nil {
          return []models.ServerDbModel{}, err
      }

      server := models.ServerDbModel{ id, id_domainconsult, address, ssl_grade, country, owner }
      servers = append(servers, server)
  }

  return servers, nil

}
