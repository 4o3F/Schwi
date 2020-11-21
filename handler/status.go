package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ApiStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	sendMsg(w, 200, "API is up")
}
