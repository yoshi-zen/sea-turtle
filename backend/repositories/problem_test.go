package repositories_test

import (
	"testing"

	"github.com/yoshi-zen/sea-turtle/backend/models"
	"github.com/yoshi-zen/sea-turtle/backend/repositories/testdata"
)

var (
	ProblemNumPerPage = 2
)

func TestSelectProblemList(t *testing.T) {
	// 40個とか用意するの大変だから1ページ2問の想定でテスト
	func start(page int)int{
		return (page-1)*ProblemNumPerPage
	}
	func end(page int)int{
		return page*ProblemNumPerPage
	}

	tests := []struct {
		title     string
		page     int
		expected []models.Problem
	}{
		{
			title: "1",
			page: 1,
			expected: testdata.ProblemTestData[start(page):end(page)],
		},
		{
			title: "2",
			page: 2,
			expected: ,
		},
	}
}
