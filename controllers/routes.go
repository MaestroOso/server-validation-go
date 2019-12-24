package controllers

import (
    "net/http"
    "github.com/go-chi/chi"
    "log"
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

func InitializeRouter() ( *chi.Mux ) {

    router := chi.NewRouter()
    log.Println("Created Router")

    for i:=0; i<len(routes); i++ {
      log.Println("Adding route", routes[i].Url)
      switch routes[i].Method {
        case "GET":
          router.Get( routes[i].Url, routes[i].HandlerFunc )
        case "POST":
          router.Post( routes[i].Url, routes[i].HandlerFunc )
      }
    }

    return router
}
