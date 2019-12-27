package controllers

import (
  "net/http" //Http handling
)

func JsonResponse( w http.ResponseWriter, httpStatus int, jsondata []byte ) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader( httpStatus )
  w.Write( jsondata )
}
