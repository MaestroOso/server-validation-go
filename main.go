package main

import (
  "fmt"
  "net/http" //Http handling
  "os"
  "controllers"
  "properties"
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
  fmt.Printf( "Attempting to run application on port %v\n", properties.Port )
  status := initRouter( )
  if status == false {
    fmt.Printf( "Error on startup of Application" )
    os.Exit( 1 )
  }
}
