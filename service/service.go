package service

import(
	"cron/dao"
)

type Service struct {
	 dao *dao.DbWorker
}

func New()(s *Service){
	s=&Service{
		dao:dao.NewSqlServer(),
	}
	return s
}