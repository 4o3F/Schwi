package main

import (
	"github.com/CardinalDevLab/Schwi-Backend/database"
	"github.com/CardinalDevLab/Schwi-Backend/handler"
	"github.com/CardinalDevLab/Schwi-Backend/log"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
)



func main() {
	database.InitDatabase()
	database.LoadConfig()
	log.SentryInit()
	handler.InitSession()
	h := RegisterHandlers()
	if database.UseTLS {
		http.ListenAndServeTLS(":21005", "./data/tls/full_chain.pem", "./data/tls/private.key", handler.SessionManager.LoadAndSave(h))
	} else {
		http.ListenAndServe(":21005", handler.SessionManager.LoadAndSave(h))
	}
}

func RegisterHandlers() http.Handler {
	router := httprouter.New()
	router.GET("/apistatus", handler.ApiStatus)
	router.POST("/user/register", handler.Register)
	router.POST("/user/login", handler.Login)
	router.GET("/user/getavatar/:uid", handler.GetAvatar)
	//router.GET("/user/logged", handler.Logged)

	c := cors.New(cors.Options{
		AllowedOrigins: database.CORSDomain,
		AllowCredentials: true,
		Debug: true,
	})

	h := c.Handler(router)
	return h
}
