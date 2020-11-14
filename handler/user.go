package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/CardinalDevLab/Schwi-Backend/database"
	"github.com/CardinalDevLab/Schwi-Backend/def"
	"github.com/CardinalDevLab/Schwi-Backend/utils"
	"github.com/asarandi/identicon"
	"github.com/julienschmidt/httprouter"
	image2 "image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)


func Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	captcha, _ := utils.NewReCAPTCHA(database.RecaptchaKey, utils.RecaptchaV2, 10 * time.Second)

	req, _ := ioutil.ReadAll(r.Body)
	body := &def.User{}

	if err := json.Unmarshal(req, body); err != nil {
		sendMsg(w, 401, "wrong json")
		return
	}

	err := captcha.Verify(body.Recaptcha)
	if err != nil {
		sendMsg(w, 401, "recaptcha wrong")
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
		user, _ := database.GetUser(0, body.Email)
		err = generateAvatar(user.Uid)
		if err != nil {
			sendMsg(w, 500, "avatar generate error")
		} else {
			sendMsg(w, 200, "register success")
		}
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

	if err != nil || response == nil || len(response.Password) == 0 || password != response.Password {
		sendMsg(w, 401, "login detail error")
		return
	} else {
		name := response.Name
		uid := response.Uid
		experience := response.Experience
		email := response.Email
		responseStr, _ := json.Marshal(def.User{Uid: uid, Name: name, Email: email, Experience: experience})
		io.WriteString(w, string(responseStr))
		SessionManager.Put(r.Context(), "uid", string(uid))
		SessionManager.Put(r.Context(), "name", name)
		SessionManager.Put(r.Context(), "email", email)
		SessionManager.Put(r.Context(), "experience", experience)
	}
}

func GetAvatar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid := p.ByName("uid")
	filepath := "./data/avatar/" + uid + ".jpg"
	if utils.CheckFileExistance(filepath) {
		reader, _ := os.Open("./data/avatar/" + uid + ".jpg")

		defer reader.Close()
		image, _, err := image2.Decode(reader)
		if err != nil {
			fmt.Println(err)
		}
		buffer := new(bytes.Buffer)
		err = jpeg.Encode(buffer, image, nil)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(buffer.Bytes())
	} else {
		sendMsg(w, 404, "Not Found")
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

func generateAvatar(uid int) error {
	//md5 := []byte(utils.DoMD5(strconv.Itoa(uid)))
	filepath := "./data/avatar/" + strconv.Itoa(uid) + ".jpg"
	err := identicon.File([]byte(strconv.Itoa(uid)), 8, filepath)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}
