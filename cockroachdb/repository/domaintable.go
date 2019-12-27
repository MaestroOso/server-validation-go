package repository

import (
  "database/sql"
  _ "github.com/lib/pq"
  "entities"
  "log"
  "cockroachdb/models"
)

func GetDomains( db *sql.DB ) ( []models.DomainDbModel, error ) {
    log.Println( "Query to get all domains" )
    // Execute queries
    rows, err := db.Query("SELECT id, host, ssl_grade, is_down FROM Domain")

    if err != nil {
        return []models.DomainDbModel{}, err
    }

    defer rows.Close()
    var domains []models.DomainDbModel

    for rows.Next() {

        var id int
        var host, ssl_grade string
        var is_down bool
        dataerror := rows.Scan( &id, &host, &ssl_grade, &is_down )

        if dataerror != nil {
            return []models.DomainDbModel{}, err
        }

        domain := models.DomainDbModel{id, host, ssl_grade, is_down}
        domains = append(domains, domain)
    }

    return domains, nil
}

func CreateOrUpdateDomain( db *sql.DB,  model entities.DomainModel, domain string ) ( bool, error ) {
  log.Println( "Query to Create or Update Domain" )
  //Validate if it is first time or domain already exists
  idDomain := CheckIfDomainExists( db, domain )
  if idDomain == -1 {
    //Domain has to be created
    domainIsAdded, domainerror := CreateDomain( db, model, domain )
    if domainerror != nil || domainIsAdded == false{
      return false, domainerror
    }
    idDomain = CheckIfDomainExists( db, domain )
  } else {
    domainIsUpdated, domainerror := UpdateDomain( db, model, domain )
    if domainerror != nil || domainIsUpdated == false {
      return false, domainerror
    }
  }

  //Create DomainConsult Row
  domainconsultres, domainconsulterr := CreateDomainConsult( db, model, idDomain )

  if domainconsulterr != nil || domainconsultres == false {
    return false, domainconsulterr
  }

  // To Do: add call to function to insert into servers table

  return true, nil
}

func CreateDomain( db *sql.DB,  model entities.DomainModel, domain string ) ( bool, error ) {
  log.Println( "Query to Create Domain" )
  // Execute queries
  _, err := db.Exec( "INSERT INTO DOMAIN(host, ssl_grade, is_down) VALUES($1,$2,$3)", domain, model.SslGrade, model.Is_down );

  if err != nil {
    return false, err
  }
  return true, nil
}

func CheckIfDomainExists( db *sql.DB,  domain string ) ( int ){
  log.Println( "Query to Select id from Domain" )

  // Prepare and query row
  row := db.QueryRow( `SELECT id FROM Domain WHERE host = $1`, domain)
  var id int
  err := row.Scan(&id)

  if err != nil {
    log.Println("The error is ", err.Error() )
    return -1
  }

  return id
}

func UpdateDomain( db *sql.DB,  model entities.DomainModel, domain string ) ( bool, error ) {
  log.Println( "Query to Update Domain" )
  // Execute queries
  _, err := db.Exec( `UPDATE DOMAIN SET(ssl_grade, is_down) = ($1,$2) WHERE host = $3`, model.SslGrade, model.Is_down, domain );

  if err != nil {
    return false, err
  }
  return true, nil
}

func GetSslGradeFromDomain( db *sql.DB,  domain string ) ( string ){
  log.Println( "Query to Select ssl_grade from Domain" )
  // Prepare and query row
  row := db.QueryRow( `SELECT ssl_grade FROM Domain WHERE host = $1`, domain )

  var grade string
  err := row.Scan(&grade)

  if err != nil {
    log.Println("The error is ", err.Error() )
    return ""
  }

  return grade
}
