package models

import "time"

type Problem struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	ProblemStatement string    `json:"problem_statement"`
	Answer           string    `json:"answer"`
	CreatedAt        time.Time `json:"created_at"`
}

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	SolvedList []int     `json:"solved_list"`
	CreatedAt  time.Time `json:"created_at"`
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Hash     string
	Uuid     string
}
