package handler

import (
	"github.com/CardinalDevLab/Schwi-Backend/def"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

func sendMsg(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	resStr, _ := jsoniter.Marshal(struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{Code: code, Msg: msg})

	io.WriteString(w, string(resStr))
}

func sendPostResponse(w http.ResponseWriter, postResponse *def.Post, sc int) {
	w.WriteHeader(sc)
	resStr, _ := jsoniter.Marshal(struct {
		Code   int       `json:"code"`
		Result *def.Post `json:"result"`
	}{sc, postResponse})

	io.WriteString(w, string(resStr))
}

func getDanmuResponse(w http.ResponseWriter, code int, data []def.DanmuGet) {
	w.WriteHeader(code)
	resStr, _ := jsoniter.Marshal(struct {
		Code int         `json:"code"`
		Data []def.DanmuGet `json:"data"`
	}{Code: code, Data: data})

	io.WriteString(w, string(resStr))
}
