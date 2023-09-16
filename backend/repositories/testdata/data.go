package testdata

import (
	"github.com/yoshi-zen/sea-turtle/backend/models"
)

var ProblemTestData = []models.Problem{
	{
		ID:               1,
		Title:            "test1",
		ProblemStatement: "赤くて甘い果物は?",
		Answer:           "りんご",
	},
	{
		ID:               2,
		Title:            "test2",
		ProblemStatement: "黄色で酸っぱい果物は?",
		Answer:           "レモン",
	},
	{
		ID:               3,
		Title:            "test3",
		ProblemStatement: "紫で甘い果物は?",
		Answer:           "ぶどう",
	},
}
