package database

import (
	"github.com/CardinalDevLab/Schwi-Backend/def"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

func GetDanmu(vid int) ([]def.DanmuGetResponse, error) {
	statement, _ := DanmuDatabase.Prepare(`SELECT danmu FROM danmu WHERE id = ?`)

	var danmuJson string
	err = statement.QueryRow(vid).Scan(&danmuJson)
	if err != nil {
		return nil, err
	}
	danmuSave := []def.DanmuSaveType{}
	err = jsoniter.Unmarshal([]byte(danmuJson), &danmuSave)

	danmu := []def.DanmuGetResponse{}

	for i := range danmuSave{
		temp := danmuSave[i]
		danmu = append(danmu, def.DanmuGetResponse{Type: temp.Type, Time: temp.Time, Text: temp.Text, Color: temp.Color, Uid: temp.Uid})
	}
	return danmu,nil
}

func WriteDanmu(danmu *def.DanmuSendResponse) (error) {
	originStatement, err := DanmuDatabase.Prepare(`SELECT danmu FROM danmu WHERE id = ?`)
	if err != nil {
		return err
	}

	var danmuJsonOrigin string
	err = originStatement.QueryRow(danmu.Id).Scan(&danmuJsonOrigin)
	if err != nil {
		return err
	}

	danmuSave := []def.DanmuSaveType{}
	err = jsoniter.Unmarshal([]byte(danmuJsonOrigin), &danmuSave)
	uid, _ := strconv.Atoi(danmu.Uid)
	danmuSave = append(danmuSave, def.DanmuSaveType{Uid: uid, Type: danmu.Type, Time: danmu.Time, Color: danmu.Color, Text: danmu.Text})

	danmuSaveJson, err := jsoniter.Marshal(danmuSave)


	updateStatement, err := DanmuDatabase.Prepare(`UPDATE danmu SET danmu = ? WHERE id = ?`)
	if err != nil {
		return err
	}

	_, err = updateStatement.Exec(&danmuSaveJson, &danmu.Id)
	if err != nil {
		return err
	}
	return nil
}