package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//获取数据库连接
func dbConn() *sql.DB {
	dbDriver := "mysql"
	dbUser := "username"
	dbPass := "password"
	dbName := "dbName"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type User struct {
	Id    int
	Name  string
	Email string
}

//CRUD：Query,Scan,Prepare,Exec
func Query() {
	db := dbConn()
	defer db.Close()
	//带参数查询
	result, err := db.Query("select id,name,email from user where id = ?", 1)
	//不带参数查询
	//result, err := db.Query("select id,name,email from user")
	if err != nil {
		panic(err.Error())
	}

	//将数据封装到结构体
	user := User{}
	var users []User
	for result.Next() {
		var id int
		var name, email string
		//注意：Scan的字段数必须与查询语句查出的字段数一致
		err := result.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.Name = name
		user.Email = email
		users = append(users, user)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}

func Add() {
	db := dbConn()
	defer db.Close()
	unix := time.Now().Unix()
	name := unix
	email := unix
	stmt, err := db.Prepare("insert into user(name,email) values(?,?)")
	if err != nil {
		panic(err.Error())
	}
	result, err := stmt.Exec(name, email)
	if err != nil {
		panic(err.Error())
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(lastInsertId)
}

func Update() {
	db := dbConn()
	defer db.Close()
	id := 1
	name := "testAdd2"
	email := "testAdd2@testAdd.com"
	stmt, err := db.Prepare("update user set name=?,email=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	result, err := stmt.Exec(name, email, id)
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffected)
}

func Delete(id int) {
	db := dbConn()
	defer db.Close()
	stmt, err := db.Prepare("delete from user where id=?")
	if err != nil {
		panic(err.Error())
	}
	result, err := stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffected)
}

func main() {
	Query()
	Add()
	Update()
	Delete(1)
}
