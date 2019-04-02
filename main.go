package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/reecerussell/ReeceRussellGo/Database"
	"github.com/reecerussell/ReeceRussellGo/Experience"
	"github.com/reecerussell/ReeceRussellGo/Home"

	"github.com/gorilla/mux"
	"github.com/reecerussell/ReeceRussellGo/Project"
)

func main() {
	fmt.Println("Reece Russell API")

	router := mux.NewRouter()

	database := Database.Database{}

	InitControllers(database, router)

	// HTTP Server

	port := "8888"

	var productionPort = os.Getenv("ASPNETCORE_PORT")
	if productionPort != "" {
		port = productionPort
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

// InitControllers ... Initialises controllers and functions for http router
func InitControllers(database Database.Database, router *mux.Router) {
	homeController := Home.HomeController{}
	homeController.Init(database, router)

	projectController := Project.Controller{}
	projectController.Init(database, router)

	experienceController := Experience.Controller{}
	experienceController.Init(database, router)
}
