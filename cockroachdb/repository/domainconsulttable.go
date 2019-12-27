package repository

import (
  "database/sql"
  _ "github.com/lib/pq"
  "entities"
  "log"
  "time"
  "cockroachdb/models"
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

func SearchDomainConsultByDomainId( db *sql.DB, id_domain int ) ( []models.DomainConsultDbModel , error ) {
  log.Println( "Query to Search DomainConsult row" )

  // Execute queries
  rows, err := db.Query(`SELECT id, id_domain, consult_time, ssl_grade, title, logo, is_down FROM DOMAINCONSULT where id_domain = $1`, id_domain )

  if err != nil {
      return []models.DomainConsultDbModel{}, err
  }

  defer rows.Close()
  var domainconsults []models.DomainConsultDbModel

  for rows.Next() {
      var id, id_domain int
      var consult_time time.Time
      var ssl_grade, title, logo string
      var is_down bool
      dataerror := rows.Scan( &id, &id_domain, &consult_time, &ssl_grade, &title, &logo, &is_down )

      if dataerror != nil {
          return []models.DomainConsultDbModel{}, err
      }

      domainconsult := models.DomainConsultDbModel{ id, id_domain, consult_time, ssl_grade, title, logo, is_down }
      domainconsults = append(domainconsults, domainconsult)
  }

  return domainconsults, nil

}
