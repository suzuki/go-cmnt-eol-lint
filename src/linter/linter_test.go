package linter

import (
	"context"
	"testing"

	"github.com/suzuki/go-cmnt-eol-lint/src/env"
)

func Test_Linter_LintFile(t *testing.T) {
	tests := map[string]struct {
		filename         string
		expectedComments []string
	}{
		"01_no_problem": {
			filename:         "01_no_problem.go",
			expectedComments: []string{},
		},
		"02_ng": {
			filename: "02_ng.go",
			expectedComments: []string{
				"Interface02 comment 1st line\nInterface02 comment 2nd line\n",
				"Struct02 is a structure for checking the comment eol\n",
				"Method is a method for checking the comment eol\n",
				"main02 is a main func for checking the comment eol\n",
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			path := getFixtureFilePath(tt.filename)

			ctx := context.Background()
			ctx = env.WithConfig(ctx)

			linter := New(env.GetConfig(ctx))
			results, err := linter.LintFile(path)
			if err != nil {
				t.Fatalf("LintFile could not lint. path=%q err=%q", path, err)
			}

			if len(results) != len(tt.expectedComments) {
				t.Fatalf("Length of results is not match. want=%d got=%d", len(tt.expectedComments), len(results))
			}

			for i, expectedComment := range tt.expectedComments {
				if results[i].Comment != expectedComment {
					t.Errorf("Comment is not match. want=%q got=%q", expectedComment, results[i].Comment)
				}
			}
		})
	}
}

func getFixtureFilePath(filename string) string {
	return "../../fixture/" + filename
}
