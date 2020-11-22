package handler

import (
	"encoding/json"
	"github.com/CardinalDevLab/Schwi-Backend/database"
	"github.com/CardinalDevLab/Schwi-Backend/def"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func AddPost(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	uid := SessionManager.Get(r.Context(), "uid")
	if uid == nil {
		sendMsg(w, 403, "not logged in")
		return
	}

	request, _ := ioutil.ReadAll(r.Body)
	postbody := &def.Post{}

	if err := json.Unmarshal(request, postbody); err != nil {
		sendMsg(w, 401, "wrong json")
		return
	}

	response, err := database.AddPost(postbody.Title, postbody.Content, postbody.Sort, postbody.Tag, postbody.Uid)
	if err != nil {
		sendMsg(w, 401, "database error")
		return
	} else {
		sendPostResponse(w, response, 200)
	}

}