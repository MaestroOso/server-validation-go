package controllers

import (
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

type Route struct {
    Url string
    Method string
    HandlerFunc http.HandlerFunc
}

var routes = []Route {
    Route{
        "/serverInfo/{server}",
        "GET",
        GetServerInformation,
    },
    Route{
        "/history",
        "GET",
        GetHistory,
    },
}

func InitializeRouter() ( *mux.Router ) {

    router := mux.NewRouter()
    fmt.Println("Created Router")

    for i:=0; i<len(routes); i++ {
      fmt.Println("Adding route", routes[i].Url)
      router.HandleFunc( routes[i].Url, routes[i].HandlerFunc ).Methods( routes[i].Method )
    }

    return router
}
