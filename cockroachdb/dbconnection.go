package cockroachdb

import (
  "database/sql"
  _ "github.com/lib/pq"
  "entities"
  "fmt"
  "time"
)

func OpenConnection( ) ( *sql.DB, error ) {
    db, err := sql.Open( "postgres", "postgresql://" + DbUser + "@" + DbUrl + DbPort + "/" + DbDatabaseName + "?sslmode=disable" )
    return db, err
}

func GetDomains( db *sql.DB ) ( []DomainDbModel, error ) {
    // Execute queries
    rows, err := db.Query("SELECT id, host, ssl_grade, is_down FROM Domain")

    if err != nil {
        return []DomainDbModel{}, err
    }

    defer rows.Close()
    var domains []DomainDbModel

    fmt.Println("Getting results");

    for rows.Next() {

        var id int
        var host, ssl_grade string
        var is_down bool
        dataerror := rows.Scan( &id, &host, &ssl_grade, &is_down )

        if dataerror != nil {
            return []DomainDbModel{}, err
        }

        domain := DomainDbModel{id, host, ssl_grade, is_down}
        fmt.Println("Domain", domain)
        domains = append(domains, domain)
        fmt.Println("Domains", domains)
    }

    return domains, nil
}

func CreateOrUpdateDomain( db *sql.DB,  model entities.DomainModel, domain string ) ( bool, error ) {
  //Validate if it is first time or domain already exists
  idDomain := CheckIfDomainExists( db, domain )
  if idDomain == -1 {
    //Domain has to be created
    domainIsAdded, domainerror := CreateDomain( db, model, domain )
    if domainerror != nil || domainIsAdded == false{
      return false, domainerror
    }
    idDomain = CheckIfDomainExists( db, domain )
  }

  //Create DomainConsult Row
  domainconsultres, domainconsulterr := CreateDomainConsult( db, model, idDomain )

  if domainconsulterr != nil || domainconsultres == false {
    return false, domainconsulterr
  }

  return true, nil
}

func CreateDomain( db *sql.DB,  model entities.DomainModel, domain string ) ( bool, error ) {
  // Execute queries
  _, err := db.Exec( "INSERT INTO DOMAIN(host, ssl_grade, is_down) VALUES($1,$2,$3)", domain, model.SslGrade, model.Is_down );

  if err != nil {
    return false, err
  }
  return true, nil
}

func CheckIfDomainExists( db *sql.DB,  domain string ) ( int ){

  // Prepare and query row
  // row := db.QueryRow( `SELECT id FROM Domain WHERE host = '$1'`, domain )

  /*THIS SHOULD NOT BE DONE BUT THE QUERY ON PREVIOUS ROW IS FAILING*/
  row := db.QueryRow( `SELECT id FROM Domain WHERE host = '` + domain + `'`)
  var id int
  err := row.Scan(&id)

  if err != nil {
    fmt.Println("The error is ", err.Error() )
    return -1
  }

  return id
}

func CreateDomainConsult( db *sql.DB,  model entities.DomainModel, id_domain int ) ( bool, error ) {
  fmt.Println( id_domain )
  // Execute queries
  _, err := db.Exec( "INSERT INTO DOMAINCONSULT(id_domain, consult_time, ssl_grade, title, logo, is_down) VALUES($1,$2,$3,$4,$5,$6)", id_domain, time.Now(), model.SslGrade, model.Title, model.Logo, model.Is_down );

  if err != nil {
    return false, err
  }
  return true, nil
}
