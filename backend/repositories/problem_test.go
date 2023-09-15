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

func TestSelectProblemList(t *testing.T) {
	// 40個とか用意するの大変だから1ページ2問の想定でテスト

	start := func(page int) int {
		return (page - 1) * ProblemNumPerPage
	}
	end := func(page int) int {
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
