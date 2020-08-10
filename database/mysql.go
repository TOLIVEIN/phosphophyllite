package database

import (
	"database/sql"
	"log"

	//...
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//InitMysql ...
func InitMysql() {
	println("Init Mysql......")
	if db == nil {
		db, _ = sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/phos")
		CreateTableWithUser()
	}
}

//CreateTableWithUser ...
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
	)`

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
