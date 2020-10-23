package handler

import (
	"encoding/json"
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
