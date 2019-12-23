package controllers

import (
    "net/http"
    "github.com/gorilla/mux"
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
}

func InitializeRouter() ( *mux.Router ) {

    router := mux.NewRouter()

    for i:=0; i<len(routes); i++ {
      router.HandleFunc( routes[i].Url, routes[i].HandlerFunc ).Methods( routes[i].Method )
    }

    return router
}
