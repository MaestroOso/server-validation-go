package services

import (
  "entities"
  "log"
  "cockroachdb/repository"
  "cockroachdb"
)

func GetHistory( ) ( entities.HistoryModel, error ) {
  log.Println( "Initializing GetHistory" )

  //Db Connection
  db, dberror := cockroachdb.OpenConnection();

  if dberror != nil {
    log.Println("Database error:", dberror.Error() )
    return entities.HistoryModel{}, dberror
  }

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

        server, serverdberror := repository.SearchServerByDomainConsultId( db, domainconsult[j].Id )

        if serverdberror != nil {
          return entities.HistoryModel{}, domainconsultdberror
        }

        for k := 0; k < len(server); k++ {
          responseServer := entities.HistoryServerModel{ server[k].Address, server[k].Ssl_grade, server[k].Country, server[k].Owner }

          responseHistoryDomainConsult.Server = append( responseHistoryDomainConsult.Server, responseServer )
        }

        responseHistoryDomain.Servers = append( responseHistoryDomain.Servers, responseHistoryDomainConsult );
    }


    response.Items = append( response.Items, responseHistoryDomain )
  }

  return response, nil
}
