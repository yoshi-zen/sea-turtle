package services

import (
	"database/sql"

	"github.com/yoshi-zen/sea-turtle/backend/models"
	"github.com/yoshi-zen/sea-turtle/backend/myerrors"
	"github.com/yoshi-zen/sea-turtle/backend/repositories"
)

type ProblemServicer struct {
}

/*
一個も問題が入手できないときのエラー処理をどうするか
ここでリストのサイズ取得して処理しよう
*/
func GetProblemListService(db *sql.DB, page int) ([]models.Problem, error) {
	problemList, err := repositories.SelectProblemList(db, page)
	if err != nil {
		return nil, err
	}

	if len(problemList) == 0 {
		err := myerrors.NoData.Wrap(NoData, "0 problem found")
		return nil, err
	}

	return problemList, nil
}

func GetProblemDetailService(db *sql.DB, ID int) (models.Problem, error) {
	problem, err := repositories.SelectProblemDetail(db, ID)
	if err != nil {
		return models.Problem{}, err
	}

	return problem, nil
}

func PostProblemService(db *sql.DB, problem models.Problem) (models.Problem, error) {
	// この変数名は要検討
	newProblem, err := repositories.InsertProblem(db, problem)
	if err != nil {
		return models.Problem{}, err
	}

	return newProblem, nil
}
