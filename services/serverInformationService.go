package services

import (
  "entities"
  "whois"
  "log"
  "ssllabs"
  "utils"
  "cockroachdb/repository"
  "database/sql"
)

func GetServerInformationService( db *sql.DB, domain string ) ( entities.DomainModel, error ) {
    log.Println("Initializing GetServerInformationService for", domain)
    var model entities.DomainModel

    //Get Information from SSL labs
    ssllabsresponse, ssllabserror := ssllabs.SslLabs( domain );

    if ssllabserror != nil {
      return entities.DomainModel{}, ssllabserror
    }

    lowestGrade := "A+"
    lowestGradeValue := utils.GradeValue( lowestGrade )

    for i:=0; i<len( ssllabsresponse.Endpoints ); i++ {

      //Get the WhoIs information for each of the servers
      whoisresponse, whoiserror := whois.WhoIs( ssllabsresponse.Endpoints[i].IpAddress )

      if whoiserror != nil {
        return entities.DomainModel{}, whoiserror
      }

      if utils.GradeValue ( ssllabsresponse.Endpoints[i].Grade ) < lowestGradeValue {
        lowestGradeValue = utils.GradeValue ( ssllabsresponse.Endpoints[i].Grade )
        lowestGrade = ssllabsresponse.Endpoints[i].Grade
      }
      newEntity := entities.MakeServerInformationModel( ssllabsresponse.Endpoints[i].IpAddress, ssllabsresponse.Endpoints[i].Grade,
            whoisresponse.Country, whoisresponse.Isp)

      model.Servers = append( model.Servers, newEntity )
    }

    model.SslGrade = lowestGrade;

    //Get Information from Web Page
    title, logo, webpageerror := utils.GetWebpageInfo( domain )
    if webpageerror != nil {
      return entities.DomainModel{}, webpageerror
    }

    model.Logo = logo
    model.Title = title

    //Query for previous ssl_grade if element exists
    model.PreviousSslGrade = repository.GetSslGradeFromDomain( db, domain )

    //Insert Data into Database
    insertResponse, dbinserterror := repository.CreateOrUpdateDomain( db, model, domain )

    if dbinserterror != nil || insertResponse == false {
      return entities.DomainModel{}, dbinserterror
    }

    return model, nil
}
