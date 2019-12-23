package main

import (
  "fmt" //Input Output Library
  "net/http" //Http handling
  //"encoding/json" //Json parsing
  "github.com/gorilla/mux"
  "os"
  "io"
  "./controllers"
  "./properties"
)

func initRouter( ) ( bool ){
  // Router variable
  router := mux.NewRouter()

  //Define the routes
  router.HandleFunc( "/serverInfo/{server}", GetServerInformation ).Methods( "GET" )
  http.Handle("/", router)

  //Start the router
  error := http.ListenAndServe( port, router );
  if error != nil {
    return false
  }
  return true
}

func main() {
  fmt.Printf( "Attempting to run application on port 8080" )
  status := initRouter( )
  if status == false {
    fmt.Printf( "Error on startup of Application" )
    os.Exit( 1 )
  }

}
