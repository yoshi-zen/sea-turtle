package models

import "time"

type Problem struct {
	ID        int
	Title     string
	Contents  string
	Answer    string
	CreatedAt time.Time
}

type User struct {
	ID         int
	Name       string
	SolvedList []int
	CreatedAt  time.Time
}
