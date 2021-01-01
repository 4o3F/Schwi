package database

func GetDanmu(vid int) (string, error) {
	statement, _ := DanmuDatabase.Prepare(`SELECT danmu FROM danmu WHERE id = ?`)

	var danmuJson string
	err = statement.QueryRow(vid).Scan(&danmuJson)
	if err != nil {
		return "", err
	}
	return danmuJson,nil
}

func SendDanmu(vid int, danmudata string) (error) {
	statement, err := DanmuDatabase.Prepare(`UPDATE danmu SET danmu = ? WHERE id = ?`)
	if err != nil {
		return err
	}

	_, err = statement.Exec(&danmudata, &vid)
	if err != nil {
		return err
	}
	return nil
}