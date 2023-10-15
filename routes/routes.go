package routes

import (
	"alura-go-personalities/controllers"
	"alura-go-personalities/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/personalities", controllers.FindAll).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.FindById).Methods("Get")
	r.HandleFunc("/personalities", controllers.Create).Methods("Post")
	r.HandleFunc("/personalities/{id}", controllers.Delete).Methods("Delete")
	r.HandleFunc("/personalities/{id}", controllers.Update).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
