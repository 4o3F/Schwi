package handler

import (
	"encoding/json"
	"github.com/CardinalDevLab/Schwi-Backend/database"
	"github.com/CardinalDevLab/Schwi-Backend/def"
	"github.com/CardinalDevLab/Schwi-Backend/utils"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	body := &def.User{}
	if err := json.Unmarshal(req, body); err != nil {
		sendMsg(w, 401, "wrong json")
		return
	}
	res, _ := database.GetUser(0, body.Email)
	if res != nil {
		sendMsg(w, 401, "email exist")
		return
	}

	if err := database.CreateUser(body.Name, body.Password, body.Email, 0); err != nil {
		sendMsg(w, 500, "database error")
		return
	} else {
		sendMsg(w, 200, "register success")
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request, _ := ioutil.ReadAll(r.Body)
	body := &def.User{}

	if err := json.Unmarshal(request, body); err != nil {
		sendMsg(w, 401, "unmarshal error")
		return
	}

	response, err := database.GetUser(0, body.Email)
	password := utils.PasswordCrypto(body.Password)

	if err != nil || len(response.Password) == 0 || password != response.Password {
		sendMsg(w, 401, "login detail error")
		return
	} else {
		name := response.Name
		uid := response.Uid
		experience := response.Uid
		email := response.Email
		responseStr, _ := json.Marshal(def.User{Uid: uid, Name: name, Email: email, Experience: experience})
		io.WriteString(w, string(responseStr))
		SessionManager.Put(r.Context(), string(uid), name)
	}
}

//func Logged(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
//	request, _ := ioutil.ReadAll(r.Body)
//	body := &def.User{}
//
//	if err := json.Unmarshal(request, body); err != nil {
//		sendMsg(w, 401, "unmarshal error")
//		return
//	}
//	fmt.Println("test1")
//	logged := SessionManager.GetString(r.Context(), string(body.Uid))
//	fmt.Println(logged)
//}
