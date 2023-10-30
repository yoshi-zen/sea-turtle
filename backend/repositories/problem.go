package repositories

import (
	"database/sql"

	"github.com/yoshi-zen/sea-turtle/backend/models"
	"github.com/yoshi-zen/sea-turtle/backend/myerrors"

	_ "github.com/go-sql-driver/mysql"
)

const (
	/*
		テストするときは、ここも変更する必要がある.
		テストケースを増やしてこのままテストできるようにしたい
	*/
	ProblemNumPerPage = 20
)

// 1ページ何個取得するかを決めないといけない
// 一旦20問
func SelectProblemList(db *sql.DB, page int) ([]models.Problem, error) {
	const sqlStr = `
	select id, title, problem_statement, answer
	from problems
	limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, ProblemNumPerPage, (page-1)*ProblemNumPerPage)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return nil, err
	}
	defer rows.Close()

	problemList := make([]models.Problem, 0)

	for rows.Next() {
		var problem models.Problem
		err := rows.Scan(&problem.ID, &problem.Title, &problem.ProblemStatement, &problem.Answer)
		if err != nil {
			err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
			return nil, err
		}
		problemList = append(problemList, problem)
	}

	return problemList, nil
}

func SelectProblemDetail(db *sql.DB, ID int) (models.Problem, error) {
	const sqlStr = `
	select id, title, problem_statement, answer
	from problems
	where id = ?;
	`

	var problem models.Problem

	row := db.QueryRow(sqlStr, ID)
	if err := row.Err(); err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return models.Problem{}, err
	}

	err := row.Scan(&problem.ID, &problem.Title, &problem.ProblemStatement, &problem.Answer)
	if err != nil {
		err = myerrors.GetDataFailed.Wrap(err, "failed to get data")
		return models.Problem{}, err
	}

	return problem, nil
}

func InsertProblem(db *sql.DB, problem models.Problem) (models.Problem, error) {
	const sqlStr = `
	insert into problems (title, problem_statement, answer, created_at) values
	(?, ?, ?, now());
	`

	var newProblem models.Problem
	newProblem.Title, newProblem.ProblemStatement, newProblem.Answer = problem.Title, problem.ProblemStatement, problem.Answer

	result, err := db.Exec(sqlStr, problem.Title, problem.ProblemStatement, problem.Answer)
	if err != nil {
		err = myerrors.InsertDataFailed.Wrap(err, "failed to insert data")
		return models.Problem{}, err
	}

	id, _ := result.LastInsertId()
	newProblem.ID = int(id)

	return newProblem, nil
}
