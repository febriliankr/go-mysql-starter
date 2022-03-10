package routes

import (
	"net/http"

	"github.com/febriliankr/go-mysql-starter/handler"
)

func InitializeRoutes() {

	http.HandleFunc("/api/v1/test", handler.TestInsert)
	http.HandleFunc("/api/v1/test/select", handler.TestSelect)
	http.Handle("/", http.FileServer(http.Dir("./public")))

}
