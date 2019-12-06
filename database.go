package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

// CreateConnection 数据库链接
func CreateConnection() (*sql.DB, error) {
	username := "sa"
	password := "pm123456"
	host := "172.16.10.26"
	dbname := "PMSoft_NewExamSystem_TZ"
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", host, username, password, 1433, dbname)
	return sql.Open("mssql", connString)
}
