package controllers

import (
	"encoding/json"
	"fmt"
	"iamajraj/sql-practice/db"
	"iamajraj/sql-practice/models"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)

	b, err := json.Marshal(users)

	fmt.Println(users)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
