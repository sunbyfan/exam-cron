package dao

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type DbWorker struct {
	//Dsn string
	Db  *sql.DB
}


func NewSqlServer() *DbWorker {
	host := "localhost"
	user := "sa"
	DBName :="db"
	password :="123456"
	connString:= fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		host,
		user,
		password,
		1433,
		DBName)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	//dbw.Db = conn
	return &DbWorker{Db: conn}
}

func (dbw *DbWorker) Close(){
	if dbw.Db != nil {
		dbw.Db.Close()
	}
}