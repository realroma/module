package main

import (
	"net/http"
)

type App struct	{
	serv *http.Server
//	bd *pgxpool.Pool
}

func (App *App) Start() {
	server.db, _ = pgxpool.Connect()
	server.serv = &http.Server{
		Addr: ":8080",
		Handler: routes,
	}
	server.serv.ListenAndServe()
}
func (App *App) Shutdown() {}
