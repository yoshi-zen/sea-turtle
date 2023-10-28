package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/yoshi-zen/sea-turtle/backend/models"
	"github.com/yoshi-zen/sea-turtle/backend/myerrors"
	"github.com/yoshi-zen/sea-turtle/backend/services"
)

type ProblemController struct {
	db *sql.DB
}

func NewProblemController(db *sql.DB) *ProblemController {
	return &ProblemController{db: db}
}

func (c *ProblemController) GetProblemListHandler(w http.ResponseWriter, req *http.Request) {
	var page int
	q := req.URL.Query()
	p, ok := q["page"]
	if ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = myerrors.BadParameter.Wrap(err, "query parameter must be number")
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		err := myerrors.BadParameter.Wrap(nil, "query parameter must have page parameter")
		return
	}

	problemList, err := services.GetProblemListService(c.db, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(problemList)
}

func (c *ProblemController) GetProblemDetailHandler(w http.ResponseWriter, req *http.Request) {
	var ID int
	q := req.URL.Query()
	id, ok := q["id"]
	if ok && len(id) > 0 {
		var err error
		ID, err = strconv.Atoi(id[0])
		if err != nil {
			err = myerrors.BadParameter.Wrap(err, "query parameter must be number")
			fmt.Println(err)
			return
		}
	}

	problem, err := services.GetProblemDetailService(c.db, ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(problem)
}

func (c *ProblemController) PostProblemHandler(w http.ResponseWriter, req *http.Request) {
	var problem models.Problem
	json.NewDecoder(req.Body).Decode(&problem)

	fmt.Println(problem)

	newProblem, err := services.PostProblemService(c.db, problem)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(newProblem)
}
