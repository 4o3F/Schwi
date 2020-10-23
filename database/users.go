package database

import (
	"database/sql"
	"github.com/CardinalDevLab/Schwi-Backend/def"
	"github.com/CardinalDevLab/Schwi-Backend/utils"
)

func CreateUser(name string, password string, email string, level int) error {
	password = utils.PasswordCrypto(password)
	statement, error := Database.Prepare("INSERT INTO users (name,password,email,level) VALUES (?,?,?,?)")
	if error != nil {
		return error
	}

	_, error = statement.Exec(name, password, email, level)
	if error != nil {
		return error
	}
	defer statement.Close()
	return nil
}

func GetUser(uid int, email string) (*def.User, error) {
	var query string
	if uid != 0 {
		query += `SELECT uid,name,password,email,level FROM users WHERE uid = ?`
	} else if email != "" {
		query += `SELECT uid,name,password,email,level FROM users WHERE email = ?`
	}
	statement, _ := Database.Prepare(query)
	var level int
	var password, name string
	if uid != 0 {
		err = statement.QueryRow(uid).Scan(&uid, &name, &password, &email, &level)
	} else if email != "" {
		err = statement.QueryRow(email).Scan(&uid, &name, &password, &email, &level)
	}
	defer statement.Close()
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &def.User{Uid: uid, Name: name, Password: password, Email: email, Level: level}

	return res, nil
}
