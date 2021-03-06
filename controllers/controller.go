package controllers

import (
  "net/http" //Http handling
  "encoding/json" //Json parsing
  "github.com/go-chi/chi"
  "services"
  "log"
)

func GetServerInformation ( w http.ResponseWriter, r *http.Request ) {
  log.Println("Controller -> GetServerInformation -> The request", r)

	domain := chi.URLParam( r, "server" )

	if domain == "" {
		http.Error( w, "domain is empty", http.StatusInternalServerError )
    log.Println("Domain error: Domain is Empty" )
    return
	}

  //Model with data to be returned
  response, error := services.GetServerInformationService( domain )

  if error != nil {
    http.Error( w, error.Error(), http.StatusInternalServerError )
    log.Println("GetServerInformationService error:", error.Error() )
    return
  }

  data, parseerror := json.Marshal( response );

  if parseerror != nil {
    http.Error( w, parseerror.Error(), http.StatusInternalServerError )
    log.Println("Json Parsing error:", parseerror.Error() )
    return
  }

  JsonResponse( w, http.StatusOK, data);
}

func GetHistory ( w http.ResponseWriter, r *http.Request ) {
  log.Println("Controller -> GetHistory -> The request", r)

  //Model with data to be returned
  response, error := services.GetHistory( )

  if error != nil {
    http.Error( w, error.Error(), http.StatusInternalServerError )
    log.Println("GetHistory error:", error.Error() )
    return
  }

  log.Println("Created response model", response)
  data, parseerror := json.Marshal( response );

  if parseerror != nil {
    http.Error( w, parseerror.Error(), http.StatusInternalServerError )
    log.Println("Json Parsing error:", parseerror.Error() )
    return
  }

  JsonResponse( w, http.StatusOK, data);

}
