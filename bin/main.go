package main

import (
	"net/http"

	"gihub.com/krishnatrea/go-assin/database"
	"gihub.com/krishnatrea/go-assin/routes"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	database.SetupMongo()
	routes.SetUp(r)
	http.ListenAndServe("localhost:3000", r)
}
