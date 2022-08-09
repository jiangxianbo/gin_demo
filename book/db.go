package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	addr := "root:jiang123@tcp(127.0.0.1)/test"
	db, err = sqlx.Connect("mysql", addr)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20) // 与数据库建立连接的最大数目
	db.SetMaxIdleConns(10) // 连接池中的最大闲置连接数
	return
}

func queryAllBook() (bookList []*Book, err error) {
	sqlStr := "select id, title, price from book"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("查询失败")
		return
	}
	return
}

func insertBook(title string, price int) (err error) {
	sqlStr := "insert into book(title, price) values (?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("插入失败")
		return
	}
	return
}

func deleteBook(id int) (err error) {
	sqlStr := "delete from book where id = ?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除失败")
		return
	}
	return
}
