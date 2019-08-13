package dao

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	_queryIds_sql            = "select id from UserExamPaper where IsOver=0 and EndTime is not null and EndTime<DATEADD(s,10,GETDATE())"
	_querySum_sql            = "select ISNUll(sum(Score),0) from UserProblemScore WHERE UserExamPaperID = ?"
	_updateUserExamPAper_sql = "UPDATE UserExamPaper SET IsOver=1,SubmitTime=GETDATE(),TotalScores=? WHERE ID=?"
)

type QueryRow struct {
	QRow *sql.Rows
}

//查询未完成Id
func (dbw *DbWorker) QueryIds() *QueryRow {
	rows, err := dbw.Db.Query(_queryIds_sql)
	checkErr(err)
	return &QueryRow{QRow: rows}
}

//求和
func (dbw *DbWorker) QueryRowSum(PaperId string) (score float64) {
	err := dbw.Db.QueryRow(_querySum_sql, PaperId).Scan(&score)
	checkErr(err)
	//fmt.Println(score)
	return
}

//更新
func (dbw *DbWorker) UpdateUserExamPaper(score float64, id string) {
	stmt, err := dbw.Db.Prepare(_updateUserExamPAper_sql)
	checkErr(err)
	res, err := stmt.Exec(score, id)
	checkErr(err)
	rowCnt, err := res.RowsAffected()
	checkErr(err)
	//fmt.Println("id")
	fmt.Printf("id=%s,rowCnt=%d \n", id, rowCnt)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
