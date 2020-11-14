package main

import (
	"github.com/CardinalDevLab/Schwi-Backend/database"
	"github.com/CardinalDevLab/Schwi-Backend/handler"
	"github.com/julienschmidt/httprouter"
	"net/http"
)



func main() {
	database.InitDatabase()
	database.LoadConfig()
	handler.InitSession()
	r := RegisterHandlers()
	//http.ListenAndServeTLS(":21005", "./data/tls/full_chain.pem", "./data/tls/private.key", handler.SessionManager.LoadAndSave(r))
	http.ListenAndServe(":21005", handler.SessionManager.LoadAndSave(r))
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user/register", handler.Register)
	router.POST("/user/login", handler.Login)
	router.GET("/user/getavatar/:uid", handler.GetAvatar)
	//router.GET("/user/logged", handler.Logged)

	return router
}
