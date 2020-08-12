package database

import (
	"database/sql"
	"fmt"
	"log"

	//...
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//InitMysql ...
func InitMysql() {
	fmt.Println("Init Mysql......")
	if db == nil {
		db, _ = sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/phos")
		CreateTableWithUser()
		CreateTableWithArticle()
	}
}

//CreateTableWithUser ...
func CreateTableWithUser() {
	sql := `create table if not exists users(
		id int(4) primary key auto_increment not null,
		username varchar(64),
		password varchar(64),
		status int(4),
		createtime int(10)
	);`
	ModifyDB(sql)
}

//CreateTableWithArticle ...
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(50),
		author varchar(20),
		tags varchar(50),
		brief varchar(255),
		content longtext,
		createtime int(10)
	);`
	ModifyDB(sql)
}

//ModifyDB ...
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//QueryRowDB ...
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

//QueryDB ...
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
