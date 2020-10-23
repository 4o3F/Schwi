package main

import (
	"github.com/CardinalDevLab/Schwi-Backend/database"
	"github.com/CardinalDevLab/Schwi-Backend/handler"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	database.InitDatabase()
	handler.InitSession()
	r := RegisterHandlers()
	http.ListenAndServe(":8084", handler.SessionManager.LoadAndSave(r))
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/user/register", handler.Register)
	router.GET("/user/login", handler.Login)
	//router.GET("/user/logged", handler.Logged)

	return router
}
