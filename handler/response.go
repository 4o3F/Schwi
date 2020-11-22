package handler

import (
	"encoding/json"
	"github.com/CardinalDevLab/Schwi-Backend/def"
	"io"
	"net/http"
)

func sendMsg(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	resStr, _ := json.Marshal(struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{Code: code, Msg: msg})

	io.WriteString(w, string(resStr))
}

func sendPostResponse(w http.ResponseWriter, postResponse *def.Post, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(struct {
		Code   int      `json:"code"`
		Result *def.Post `json:"result"`
	}{sc, postResponse})

	io.WriteString(w, string(resStr))
}
