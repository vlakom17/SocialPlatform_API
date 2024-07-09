package app

import (
	routes "api/internal/presentation/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func App() {
	router := mux.NewRouter()

	routes.InitExternalRoutes(router)
	routes.InitEventsRoutes(router)
	routes.InitRoomRoutes(router)
	routes.InitUserRoutes(router)

	log.Println("API Gateway сервис запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
