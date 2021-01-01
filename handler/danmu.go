package handler

import (
	"database/sql"
	"github.com/CardinalDevLab/Schwi-Backend/database"
	"github.com/CardinalDevLab/Schwi-Backend/def"
	"github.com/CardinalDevLab/Schwi-Backend/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetDanmu(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	vidstr := p.ByName("vid")
	vid, err := strconv.Atoi(vidstr)
	if err != nil {
		sendMsg(w, 401, "vid not found")
		return
	}
	DanmuJSON, err := database.GetDanmu(vid)
	if err != nil {
		if err == sql.ErrNoRows {
			sendMsg(w, 401, "vid not found")
			return
		} else {
			utils.ErrorHandler(err)
			sendMsg(w, 500, "database error")
			return
		}
	}
	Danmu := []def.DanmuGet{}
	err = jsoniter.Unmarshal([]byte(DanmuJSON), &Danmu)
	if err != nil {
		utils.ErrorHandler(err)
		sendMsg(w, 500, "database error")
		return
	}
	getDanmuResponse(w, 200, Danmu)
}

func SendDanmu(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	request, _ := ioutil.ReadAll(r.Body)
	newDanmu := def.Danmu{}
	err := jsoniter.Unmarshal(request, &newDanmu)
	if err != nil {
		sendMsg(w, 401, "request json error")
		return
	}
	vidstr := p.ByName("vid")
	vid, err := strconv.Atoi(vidstr)
	if err != nil {
		sendMsg(w, 401, "vid not found")
		return
	}
	originDanmuJSON,err := database.GetDanmu(vid)
	if err != nil {
		utils.ErrorHandler(err)
		sendMsg(w, 500, "database error")
		return
	}

	originDanmu := []def.Danmu{}
	err = jsoniter.Unmarshal([]byte(originDanmuJSON), &originDanmu)
	if err != nil {
		utils.ErrorHandler(err)
		sendMsg(w, 500, "databse error")
		return
	}

	writeDanmu := append(originDanmu, newDanmu)
	writeDanmuJSON,err := jsoniter.Marshal(writeDanmu)
	if err != nil {
		utils.ErrorHandler(err)
		sendMsg(w, 500, "databse error")
		return
	}

	err = database.SendDanmu(vid, string(writeDanmuJSON))
	if err != nil {
		utils.ErrorHandler(err)
		sendMsg(w, 500, "databse error")
		return
	}
	sendMsg(w, 200, "send success")
}