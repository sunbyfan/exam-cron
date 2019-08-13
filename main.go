package main

import (
	"cron/service"
	"log"

	"github.com/robfig/cron"
)

func main() {
	log.Println("任务开始...")
	c := cron.New()
	//spec := "0 */5 * * * ?" //5分执行一次
	spec:="*/10 * * * * ?"//10s执行一次
	c.AddFunc(spec, func() {
		log.Println("运行计算分数的任务...")
		//mssql.NewSqlServer()
		s:=service.New()
		s.Task()
	})
	c.Start()

	select {}
}
