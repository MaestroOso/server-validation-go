package services

import (
  "entities"
  "whois"
  "fmt"
  "ssllabs"
  "utils"
  "cockroachdb"
  "database/sql"
)

func GetServerInformationService( domain string ) ( entities.DomainModel, error ) {
    var model entities.DomainModel

    //Get Information from SSL labs
    ssllabsresponse, ssllabserror := ssllabs.SslLabs( domain );

    if ssllabserror != nil {
      return entities.DomainModel{}, ssllabserror
    }

    fmt.Println( ssllabsresponse )
    lowestGrade := "A+"
    lowestGradeValue := utils.GradeValue( lowestGrade )

    for i:=0; i<len( ssllabsresponse.Endpoints ); i++ {

      //Get the WhoIs information for each of the servers
      fmt.Println("Searching for", ssllabsresponse.Endpoints[i].IpAddress)
      whoisresponse, whoiserror := whois.WhoIs( ssllabsresponse.Endpoints[i].IpAddress )

      if whoiserror != nil {
        return entities.DomainModel{}, whoiserror
      }

      fmt.Println( whoisresponse )

      if utils.GradeValue ( ssllabsresponse.Endpoints[i].Grade ) < lowestGradeValue {
        lowestGradeValue = utils.GradeValue ( ssllabsresponse.Endpoints[i].Grade )
        lowestGrade = ssllabsresponse.Endpoints[i].Grade
      }
      newEntity := entities.MakeServerInformationModel( ssllabsresponse.Endpoints[i].IpAddress, ssllabsresponse.Endpoints[i].Grade,
            whoisresponse.Country, "")

      model.Servers = append( model.Servers, newEntity )
    }

    model.SslGrade = lowestGrade;

    fmt.Println("The resulting entity is", model)

    //Get Information from Web Page
    logo, title, webpageerror := utils.GetWebpageInfo( domain )
    if webpageerror != nil {
      return entities.DomainModel{}, webpageerror
    }

    model.Logo = logo
    model.Title = title

    return model, nil
}

func GetHistory( db *sql.DB ) ( []cockroachdb.DomainDbModel, error ){
  domains, err := cockroachdb.GetDomains( db )

  if err != nil {
    return []cockroachdb.DomainDbModel{}, err
  }

  return domains, nil
}
