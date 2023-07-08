package controllers

import (
	"encoding/json"
	"fmt"
	"iamajraj/sql-practice/db"
	"iamajraj/sql-practice/models"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var user models.User
	err = json.Unmarshal(b, &user)

	if err != nil {
		panic(err)
	}

	db.DB.Create(&user)

	fmt.Fprint(w, "User has been created")
}
