package repositories_test

import (
	"testing"

	"github.com/yoshi-zen/sea-turtle/backend/models"
	"github.com/yoshi-zen/sea-turtle/backend/repositories"
	"github.com/yoshi-zen/sea-turtle/backend/repositories/testdata"
)

var (
	ProblemNumPerPage = 2
)

// 40個とか用意するの大変だから1ページ2問の想定でテスト
func TestSelectProblemList(t *testing.T) {
	// 指定したページで表示される問題がテストデータの配列のどこからどこまでかを計算する
	start := func(page int) int {
		return (page - 1) * ProblemNumPerPage
	}
	end := func(page int) int {
		// テストデータの数をはみ出してしまう場合をケア
		res := min(page*ProblemNumPerPage, len(testdata.ProblemTestData))
		return res
	}

	tests := []struct {
		title    string
		page     int
		expected []models.Problem
	}{
		{
			title:    "1",
			page:     1,
			expected: testdata.ProblemTestData[start(1):end(1)],
		},
		{
			title:    "2",
			page:     2,
			expected: testdata.ProblemTestData[start(2):end(2)],
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			gotList, err := repositories.SelectProblemList(db, test.page)
			if err != nil {
				t.Fatal(err)
			}

			for index, got := range gotList {
				if got != test.expected[index] {
					t.Errorf("problem list is expected %+v but got %+v\n", test.expected, got)
				}
			}
		})
	}
}

func TestSelectProblemDetail(t *testing.T) {
	tests := []struct {
		title    string
		ID       int
		expected models.Problem
	}{
		{
			title:    "1",
			ID:       1,
			expected: testdata.ProblemTestData[0],
		},
		{
			title:    "2",
			ID:       2,
			expected: testdata.ProblemTestData[1],
		},
		{
			title:    "3",
			ID:       3,
			expected: testdata.ProblemTestData[2],
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			got, err := repositories.SelectProblemDetail(db, test.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got != test.expected {
				t.Errorf("got problem is expected %+v but got %+v\n", test.expected, got)
			}
		})
	}
}

// インサートできるかをテストする
// インサートして、それを取得して一致するか判定
func TestInsertProblem(t *testing.T) {
	tests := []struct {
		title   string
		problem models.Problem
	}{
		{
			title: "1",
			problem: models.Problem{
				Title:            "test1",
				ProblemStatement: "tes1tes1",
				Answer:           "tes1tes1tes1",
			},
		},
		{
			title: "2",
			problem: models.Problem{
				Title:            "test2",
				ProblemStatement: "tes2tes2",
				Answer:           "tes2tes2tes2",
			},
		},
		{
			title: "3",
			problem: models.Problem{
				Title:            "test3",
				ProblemStatement: "tes3tes3",
				Answer:           "tes3tes3tes3",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			problem, err := repositories.InsertProblem(db, test.problem)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repositories.SelectProblemDetail(db, problem.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got != problem {
				t.Errorf("inserted problem is expected %+v but got %+v\n", problem, got)
			}
		})
	}
}
