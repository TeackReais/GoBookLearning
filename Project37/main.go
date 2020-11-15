package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgreSQLSettings struct {
	dbname                    string
	user                      string
	password                  string
	host                      string
	port                      int
	sslmode                   string
	fallback_application_name string
	connect_timeout           int
	sslcert                   string
	sslkey                    string
	sslrootcert               string
}

//SSL默认关闭
func PSQLSettingsStr(Settings PostgreSQLSettings) (string, error) {
	var SettingsStr string
	var ssl bool = false
	if Settings.dbname == "" || Settings.user == "" || Settings.password == "" {
		return "", errors.New("信息不全。")
	}
	SettingsStr = fmt.Sprintf("dbname=%s user=%s password=%s", Settings.dbname, Settings.user, Settings.password)
	if Settings.host != "" {
		SettingsStr = SettingsStr + fmt.Sprintf(" host=%s", Settings.host)
	}
	if Settings.port != 0 {
		SettingsStr = SettingsStr + fmt.Sprintf(" port=%d", Settings.port)
	}
	if Settings.sslmode != "" && Settings.sslmode != "disable" {
		SettingsStr = SettingsStr + fmt.Sprintf(" sslmode=%s", Settings.sslmode)
		ssl = true
	} else {
		SettingsStr = SettingsStr + fmt.Sprint(" sslmode=disable")
	}
	if Settings.fallback_application_name != "" {
		SettingsStr = SettingsStr + fmt.Sprintf(" fallback_application_name=%s", Settings.fallback_application_name)
	}
	if Settings.connect_timeout != 0 {
		SettingsStr = SettingsStr + fmt.Sprintf(" connect_timeout=%d", Settings.connect_timeout)
	}
	if ssl {
		SettingsStr = SettingsStr + fmt.Sprintf(" sslcert=%s", Settings.sslcert)
		SettingsStr = SettingsStr + fmt.Sprintf(" sslkey=%s", Settings.sslkey)
		SettingsStr = SettingsStr + fmt.Sprintf(" sslrootcert=%s", Settings.sslrootcert)
	}
	return SettingsStr, nil
}

type users struct {
	id       int
	account  string
	password string
}

func main() {
	var settings PostgreSQLSettings = PostgreSQLSettings{
		dbname:   "mydb",
		user:     "dbadmin",
		password: "20010925abcd",
	}
	settingsStr, _ := PSQLSettingsStr(settings)
	db, err := sql.Open("postgres", settingsStr)
	err = db.Ping()
	fmt.Println(err)
	row := db.QueryRow("SELECT id,account,password FROM users where id=2;")
	var user1 users
	err = row.Scan(&user1.id, &user1.account, &user1.password)
	fmt.Println(user1)
	if err == sql.ErrNoRows {
		fmt.Println("未查询到相应的用户")
	}
}
