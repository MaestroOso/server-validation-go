package controllers

import (
  "net/http" //Http handling
  "encoding/json" //Json parsing
  "github.com/gorilla/mux"
  "services"
  "cockroachdb"
  "fmt"
)

func GetServerInformation ( w http.ResponseWriter, r *http.Request ) {
  //Init db connection
  db, dberror := cockroachdb.OpenConnection();

  if dberror != nil {
    http.Error( w, dberror.Error(), http.StatusInternalServerError )
    return
  }

	params := mux.Vars( r )

  if len(params) == 0 {
    http.Error( w, "params header is not present in request", http.StatusInternalServerError )
    return
  }

	domain := params["server"]

	if domain == "" {
		http.Error( w, "domain is empty", http.StatusInternalServerError )
	}

  //Model with data to be returned
  response, error := services.GetServerInformationService( db, domain )

  if error != nil {
    http.Error( w, error.Error(), http.StatusInternalServerError )
  }

  data, parseerror := json.Marshal( response );

  if parseerror != nil {
    http.Error( w, parseerror.Error(), http.StatusInternalServerError )
  }

  JsonResponse( w, http.StatusOK, data);
}

func GetHistory ( w http.ResponseWriter, r *http.Request ) {
  db, dberror := cockroachdb.OpenConnection();

  if dberror != nil {
    http.Error( w, dberror.Error(), http.StatusInternalServerError )
    return
  }

  //Model with data to be returned
  response, error := services.GetHistory( db )

  if error != nil {
    http.Error( w, error.Error(), http.StatusInternalServerError )
    return
  }

  fmt.Println("Response", response)
  data, parseerror := json.Marshal( response );

  if parseerror != nil {
    http.Error( w, parseerror.Error(), http.StatusInternalServerError )
    return
  }

  fmt.Println("Parsed data is", data)
  JsonResponse( w, http.StatusOK, data);

}
