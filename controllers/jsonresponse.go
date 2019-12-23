package controllers

import (
  "net/http" //Http handling
)

func JsonResponse( w http.ResponseWriter, httpStatus int, jsondata []byte ) {
  w.WriteHeader( httpStatus )
  w.Write( jsondata )
}
