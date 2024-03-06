package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SQLinit() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/test240214?charset=utf8")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("sql.DB 結構已建立")

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("資料庫連線成功")
	DBCreate := `
	CREATE TABLE IF NOT EXISTS user
	(
		id INT NOT NULL UNIQUE,
		name VARCHAR(20),
		password  VARCHAR(30)
	);`

	_, err = db.Exec(DBCreate)
	if err != nil {
		panic(err)
	}
	fmt.Println("表格 user 已建立")
}
