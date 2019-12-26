package services

import (
  "entities"
  "log"
  "cockroachdb/repository"
  "database/sql"
)

func GetHistory( db *sql.DB ) ( entities.HistoryModel, error ){
  log.Println( "Initializing GetHistory" )
  //Get all looked up domains
  domains, err := repository.GetDomains( db )

  if err != nil {
    return entities.HistoryModel{}, err
  }

  var response entities.HistoryModel

  for i := 0; i < len(domains); i++ {

    responseHistoryDomain := entities.HistoryDomainModel{ domains[i].Host, []entities.HistoryDomainConsultModel{} }

    //Lookup for all associated DomainConsult rows related with the domains
    domainconsult, domainconsultdberror := repository.SearchDomainConsultByDomainId( db, domains[i].Id )

    if domainconsultdberror != nil {
      return entities.HistoryModel{}, domainconsultdberror
    }

    for j := 0; j < len(domainconsult); j++ {
        responseHistoryDomainConsult := entities.HistoryDomainConsultModel{ domainconsult[j].Consult_time, domainconsult[j].Ssl_grade, domainconsult[j].Title, domainconsult[j].Logo, domainconsult[j].Is_down, []entities.HistoryServerModel{} }

        responseHistoryDomain.Servers = append( responseHistoryDomain.Servers, responseHistoryDomainConsult );
    }


    response.Items = append( response.Items, responseHistoryDomain )
  }

  return response, nil
}
