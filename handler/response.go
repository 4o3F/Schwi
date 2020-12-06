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
		Code   int      `json:"code"`
		Result *def.Post `json:"result"`
	}{sc, postResponse})

	io.WriteString(w, string(resStr))
}

func sendDanmuGetResponse(w http.ResponseWriter, code int, data []def.DanmuGetResponse) {
	w.WriteHeader(code)
	resStr, _ := jsoniter.Marshal(struct {
		Code int    `json:"code"`
		Data  []def.DanmuGetResponse `json:"data"`
	}{Code: 0, Data: data})

	io.WriteString(w, string(resStr))
}
func sendDanmuSendResponse(w http.ResponseWriter, code int, data *def.DanmuSendResponse) {
	w.WriteHeader(code)
	resStr, _ := jsoniter.Marshal(data)

	io.WriteString(w, string(resStr))
}