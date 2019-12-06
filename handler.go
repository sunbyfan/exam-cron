package main

// ExamHandler struct
type ExamHandler struct {
	repo Repository
}

// ExecTask 执行未计算分数的考试任务
func (e *ExamHandler) ExecTask() error {
	var (
		paperIDs []string
		err      error
	)
	paperIDs, err = e.repo.QueryIds()
	if err != nil {
		return err
	}
	for _, paperID := range paperIDs {
		score, err := e.repo.QueryScoreSum(paperID)
		if err != nil {
			return err
		}
		err = e.repo.UpdateUserExamPaper(paperID, score)
		if err != nil {
			return err
		}
	}
	return nil
}
