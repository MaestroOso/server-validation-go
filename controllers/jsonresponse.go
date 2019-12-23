package controllers

import (
  "net/http" //Http handling
  //"encoding/json" //Json parsing
  "io"
)

func JsonResponse( w http.ResponseWriter, httpStatus int, jsondata string ) {
  w.WriteHeader( httpStatus )
  io.WriteString( w, jsondata );
}
