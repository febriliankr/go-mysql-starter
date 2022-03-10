package handler

import (
	"log"
	"net/http"

	"github.com/febriliankr/go-mysql-starter/helper"
)

func TestInsert(w http.ResponseWriter, r *http.Request) {

	type User struct {
		Username, Email, Password string
	}

	user := User{
		Username: "Elvan",
		Email:    "elvanwiyarta@gmail.com",
		Password: "123123123",
	}

	db, err := helper.InitializeDB()
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, err)
		log.Fatalln(err)
	}
	_, err = db.Queryx("INSERT INTO users (username, email, password) VALUES (?,?,?)", user.Username, user.Email, user.Password)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, err)
		log.Fatalln(err)
	}

	helper.SendResponse(w, http.StatusOK, "Inserted!")
}
func TestSelect(w http.ResponseWriter, r *http.Request) {

	type User struct {
		Id, Username, Email, Password string
	}

	db, err := helper.InitializeDB()
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, err)
		log.Fatalln(err)
	}
	user := []User{}

	err = db.Select(&user, "select * from users")
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, err)
		log.Fatalln(err)
	}

	helper.SendResponse(w, http.StatusOK, user)
}
