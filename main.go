package main

import (
  "net/http" //Http handling
  "os"
  "controllers"
  "properties"
  "log"
)

func initRouter( ) ( bool ){
  // Router variable
  router := controllers.InitializeRouter()
  http.Handle("/", router)

  //Start the router
  error := http.ListenAndServe( properties.Port, router );
  if error != nil {
    return false
  }
  return true
}

func main() {
  //Logging should be done on a file
  log.Printf( "Attempting to run application on port %v\n", properties.Port )
  status := initRouter( )
  if status == false {
    log.Printf( "Error on startup of Application" )
    os.Exit( 1 )
  }
}
