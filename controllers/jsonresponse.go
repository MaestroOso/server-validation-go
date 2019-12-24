package controllers

import (
  "net/http" //Http handling
)

func JsonResponse( w http.ResponseWriter, httpStatus int, jsondata []byte ) {
  w.WriteHeader( httpStatus )
  w.Header().Set("Content-Type", "application/json")
  w.Write( jsondata )
}
