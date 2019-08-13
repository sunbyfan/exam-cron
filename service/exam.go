package service

import (
	"cron/dao"
	"github.com/satori/go.uuid"
	"log"
)


func (s *Service) Task(){
	var (
		q *dao.QueryRow
	)
	q =s.dao.QueryIds()
	rows:=q.QRow
	defer rows.Close()
	for rows.Next() {
		var id []byte
		err := rows.Scan(&id)
		if err!=nil{
			log.Fatal(err)
		}
		b := id
		b[0], b[1], b[2], b[3] = b[3], b[2], b[1], b[0]
		b[4], b[5] = b[5], b[4]
		b[6], b[7] = b[7], b[6]
		uid, _ := uuid.FromBytes(b)
		uuidStr := uid.String()
		score := s.dao.QueryRowSum(uuidStr)
		s.dao.UpdateUserExamPaper(score, uuidStr)
	}
}