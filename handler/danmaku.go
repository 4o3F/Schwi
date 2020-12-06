package handler

import (
	"database/sql"
	"fmt"
	"github.com/CardinalDevLab/Schwi-Backend/database"
	"github.com/CardinalDevLab/Schwi-Backend/def"
	"github.com/CardinalDevLab/Schwi-Backend/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetDanmu(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	vid, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		sendMsg(w, 403, "wrong vid")
		return
	}
	danmu, err := database.GetDanmu(vid)
	if err == sql.ErrNoRows {
		sendMsg(w, 404, "video danmu not found")
		return
	} else if err != nil {
		utils.ErrorHandler(err)
		sendMsg(w, 500, "server error")
		return
	}
	sendDanmuGetResponse(w, 200, danmu)
}

func SendDanmu(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	request, _ := ioutil.ReadAll(r.Body)
	body := &def.DanmuSendResponse{}

	if err := jsoniter.Unmarshal(request, body); err != nil {
		fmt.Println(err)
		sendMsg(w, 401, "unmarshal error")
		return
	}

	err := database.WriteDanmu(body)
	if err != nil {
		utils.ErrorHandler(err)
		sendMsg(w, 500, "database error")
		return
	}
	date := time.Now().UnixNano() / 1e6
	ip := utils.RemoteIp(r)
	player, err := strconv.Atoi(body.Id)
	returnData := &def.DanmuSendResponse{Uid: body.Uid, Color: body.Color, Date: date, Ip: ip, Player: player, Text: body.Text, Time: body.Time, Type: body.Type}
	sendDanmuSendResponse(w, 200, returnData)
}
