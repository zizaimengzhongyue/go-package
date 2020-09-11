package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MySQL() {
	db, err := sql.Open("mysql", "root:123456@/test")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	result, err := db.Exec("INSERT test (uid, name) Values(?, ?)", 10, "赵钱孙李")
	if err != nil {
		panic(err)
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(lastID)

	rows, err := db.Query("SELECT * FROM test")
	defer rows.Close()
	var id, uid int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &uid, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, uid, name)
	}

	result, err = db.Exec("DELETE FROM test WHERE id = ?", 1)
	if err != nil {
		panic(err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(affected)

	result, err = db.Exec("UPDATE test SET name = '新名字' WHERE ID = 3")
	if err != nil {
		panic(err)
	}
	affected, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(affected)

	rows, err = db.Query("SELECT * FROM test")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &uid, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, uid, name)
	}
}

func main() {
	MySQL()
}
