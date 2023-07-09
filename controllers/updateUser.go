package controllers

import (
	"encoding/json"
	"iamajraj/sql-practice/db"
	"iamajraj/sql-practice/models"
	"iamajraj/sql-practice/utils"
	"io"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var user models.User

	if err := json.Unmarshal(b, &user); err != nil {
		panic(err)
	}

	if user.Name == "" || user.Age == 0 || user.Email == "" || user.ID == 0 {
		utils.SendError(w, "Please provide all the fields along with the ID", http.StatusBadRequest)
	}

	db.DB.Model(&user).Updates(&user)

	w.Write([]byte("Ok"))
}
