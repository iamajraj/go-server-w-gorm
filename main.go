package main

import (
	"fmt"
	"iamajraj/sql-practice/controllers"
	_ "iamajraj/sql-practice/db"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.HomeHandler)
	r.Post("/create-user", controllers.CreateUser)
	r.Put("/update-user", controllers.UpdateUser)

	fmt.Println("Server has been started 🚀")
	http.ListenAndServe("127.0.0.1:3000", r)
}
