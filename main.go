package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/reecerussell/ReeceRussellGo/Database"

	"github.com/gorilla/mux"
	"github.com/reecerussell/ReeceRussellGo/Project"
)

func main() {
	fmt.Println("Reece Russell API")

	router := mux.NewRouter()

	dataBase := Database.Database{}
	dataBase.Init()

	projectController := Project.ProjectController{}
	projectController.Init(dataBase, router)

	port := "8888"

	var productionPort = os.Getenv("ASPNETCORE_PORT")
	if productionPort != "" {
		port = productionPort
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
