package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

// Repository 接口
type Repository interface {
	QueryIds() ([]string, error)
	QueryScoreSum(paperID string) (score float64, err error)
	UpdateUserExamPaper(id string, score float64) error
}

// ExamRepository sql.DB
type ExamRepository struct {
	DB *sql.DB
}

// QueryIds 查询考试未完成Id
func (repo *ExamRepository) QueryIds() ([]string, error) {
	query := `select id from UserExamPaper where IsOver=0 and EndTime is not null and EndTime<DATEADD(s,10,GETDATE())`
	rows, err := repo.DB.Query(query)
	if err != nil {
		//log.Fatalln(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	var result []string
	for rows.Next() {
		var id []byte
		err = rows.Scan(&id)
		if err != nil {
			//log.Fatalln(err)
			return nil, err
		}
		result = append(result, fixUUID(id))
	}
	return result, nil
}

// QueryScoreSum 根据paperId查询未完成考试的总分
func (repo *ExamRepository) QueryScoreSum(paperID string) (score float64, err error) {
	query := `select ISNUll(sum(Score),0) from UserProblemScore WHERE UserExamPaperID = ?`
	err = repo.DB.QueryRow(query, paperID).Scan(&score)
	return
}

// UpdateUserExamPaper 修改试卷成绩
func (repo *ExamRepository) UpdateUserExamPaper(id string, score float64) error {
	query := `UPDATE UserExamPaper SET IsOver=1,SubmitTime=GETDATE(),TotalScores=? WHERE ID=?`
	stmt, err := repo.DB.Prepare(query)
	if err != nil {
		return nil
	}
	res, err := stmt.Exec(score, id)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		return err
	}
	fmt.Printf("id=%s,rowCnt=%d \n", id, rowsAfected)
	return nil
}

// fixUuid 修正mssql下GUID
func fixUUID(b []byte) string {
	b[0], b[1], b[2], b[3] = b[3], b[2], b[1], b[0]
	b[4], b[5] = b[5], b[4]
	b[6], b[7] = b[7], b[6]
	fixid, _ := uuid.FromBytes(b)
	return fixid.String()
}
