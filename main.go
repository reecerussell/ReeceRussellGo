package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	"github.com/reecerussell/ReeceRussellGo/Authentication"
	"github.com/reecerussell/ReeceRussellGo/Database"
	"github.com/reecerussell/ReeceRussellGo/Education"
	"github.com/reecerussell/ReeceRussellGo/Experience"
	"github.com/reecerussell/ReeceRussellGo/Home"
	"github.com/reecerussell/ReeceRussellGo/Project"
	"github.com/reecerussell/ReeceRussellGo/Skills"
	"github.com/reecerussell/ReeceRussellGo/Statics"
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

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS()(router)))
}

// InitControllers ... Initialises controllers and functions for http router
func InitControllers(database Database.Database, router *mux.Router) {
	homeController := Home.Controller{}
	homeController.Init(database, router)

	authController := Authentication.Controller{}
	authController.Init(router)

	projectController := Project.Controller{}
	projectController.Init(database, router)

	experienceController := Experience.Controller{}
	experienceController.Init(database, router)

	educationController := Education.Controller{}
	educationController.Init(database, router)

	skillController := Skills.Controller{}
	skillController.Init(database, router)

	staticsController := Statics.Controller{}
	staticsController.Init(router)
}
