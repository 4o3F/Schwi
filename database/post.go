package database

import (
	"database/sql"
	"github.com/CardinalDevLab/Schwi-Backend/def"
	"time"
)

func AddPost(title string, content string, sort string, tag string, uid int) (*def.Post, error) {
	cstZone := time.FixedZone("CST", 8*3600)
	ctime := time.Now().In(cstZone).Format("2020-01-01 00:00")
	statement, err := PostDatabase.Prepare("INSERT INTO posts (title,content,sort,tag,time,uid) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = statement.Exec(title, content, sort, tag, ctime, uid)
	if err != nil {
		return nil, err
	}
	res := &def.Post{Title: title, Content: content, Sort: sort, Tag: tag, Time: ctime, Uid: uid}
	defer statement.Close()

	return res, nil
}

func UpdatePost(pid int, title string, content string, sort string, tag string, time string) (*def.Post, error) {
	statement, err := PostDatabase.Prepare("UPDATE posts SET title=?,content=?,sort=?,tag=?,time=? WHERE id =?")
	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(&title, &content, &sort, &tag, &time, &pid)
	if err!= nil {
		return nil, err
	}
	res := &def.Post{Pid: pid, Title: title, Content: content, Sort: sort, Tag: tag, Time: time}
	defer statement.Close()
	return res, nil
}

func DeletePost(pid int) error {
	statement, err := PostDatabase.Prepare("DELETE FROM posts WHERE id=?")
	if err != nil {
		return err
	}

	_, err = statement.Exec(pid)
	if err != nil {
		return err
	}
	defer statement.Close()

	return nil
}

func GetPost(id int) (*def.Post, error) {
	statementt, err := PostDatabase.Prepare(`SELECT posts.pid,posts.title,posts.content,posts.sort,posts.tag,posts.time,users.id,users.username 
FROM posts INNER JOIN users ON posts.uid = users.uid WHERE posts.pid = ?`)
	if err != nil {
		return nil, err
	}
	var pid, uid int
	var title, content, status, sort, tag, ctime, username string

	err = statementt.QueryRow(id).Scan(&pid, &title, &content, &status, &sort, &tag, &ctime, &uid, &username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer statementt.Close()

	res := &def.Post{Pid: pid, Title: title, Content: content, Sort: sort, Tag: tag, Time: ctime, Uid: uid, Username: username}

	return res, nil
}
