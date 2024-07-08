package app

import (
	"api/internal/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func App() {
	router := mux.NewRouter()
	routes.InitRoutes(router)

	log.Println("API Gateway сервис запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
