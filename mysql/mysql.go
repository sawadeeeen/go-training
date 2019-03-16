package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func selectSQL(db *sql.DB) {
	rows, err := db.Query("select * from tbl")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id:%s\\tname:%s\\", id, name)
	}
}

func main() {
	// sql.Open("mysql", "user:password@tcp(host:port)/db_name)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ianBlog")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//	selectSQL(db)
}
