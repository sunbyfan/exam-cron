package main

import (
	"log"

	"github.com/robfig/cron"
)

func main() {

	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer db.Close()

	repo := &ExamRepository{db}

	handler := &ExamHandler{repo}

	log.Println("任务开始...")
	c := cron.New()
	//spec := "0 */5 * * * ?" //5分执行一次
	spec := "*/10 * * * * ?" //10s执行一次
	c.AddFunc(spec, func() {
		log.Println("运行计算分数的任务...")
		//mssql.NewSqlServer()
		err = handler.ExecTask()
		if err != nil {
			log.Fatalln(err)
		}
	})
	c.Start()

	select {}
}
