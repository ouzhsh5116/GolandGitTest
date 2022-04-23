package dao

import (
	"book/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB() (err error) {
	addr := "root:123456@tcp(127.0.0.1:3306)/book_hw"
	db, err = sqlx.Connect("mysql", addr)
	if err != nil {
		return err
	}
	// 最大连接
	db.SetMaxOpenConns(100)
	// 最大空闲
	db.SetMaxIdleConns(16)
	fmt.Println("数据库连接成功")
	return
}

// QueryAllBook Load
// 查询所有数据
func QueryAllBook() (bookList []*model.Book, err error) {
	sqlStr := "select id ,title,price from book"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("查询失败")
		return
	}
	return
}

// Del 删除数据
func Del(id int) (err error) {
	sqlStr := "delete from book where id = ?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除失败")
		return
	}
	return
}

// Add 添加单条书木
func Add(title string, price int64) (err error) {
	sqlStr := "insert into book(title,price) values(?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("插入失败")
		return
	}
	return
}
