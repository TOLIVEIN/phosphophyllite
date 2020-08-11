package models

import (
	"fmt"
	"phosphophyllite/database"
)

//User ...
type User struct {
	ID         int
	Username   string
	Password   string
	Status     int
	Createtime int64
}

//InsertUser ...
func InsertUser(user User) (int64, error) {
	return database.ModifyDB("insert into users(username, password, status, createtime) values (?, ?, ?, ?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

//QueryUserWithCondition ...
func QueryUserWithCondition(condition string) int {
	sql := fmt.Sprintf("select id from users %s", condition)
	println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//QueryUserWithUsername ...
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWithCondition(sql)
}

//QueryUserWithParam ...
func QueryUserWithParam(username string, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWithCondition(sql)
}
