package controllers

import (
  "net/http" //Http handling
  "encoding/json" //Json parsing
  "github.com/gorilla/mux"
  "services"
)

func GetServerInformation ( w http.ResponseWriter, r *http.Request ) {
	params := mux.Vars( r )

  if len(params) == 0 {
    http.Error( w, http.StatusText( http.StatusInternalServerError ), http.StatusInternalServerError )
  }

	domain := params["server"]

	if domain == "" {
		http.Error( w, http.StatusText( http.StatusInternalServerError ), http.StatusInternalServerError )
	}

  //Model with data to be returned
  response, error := services.GetServerInformationService( domain )

  if error != nil {
    http.Error( w, http.StatusText( http.StatusInternalServerError ), http.StatusInternalServerError )
  }

  data, parseerror := json.Marshal( response );

  if parseerror != nil {
    http.Error( w, http.StatusText( http.StatusInternalServerError ), http.StatusInternalServerError )
  }

  JsonResponse( w, http.StatusOK, data);
}
