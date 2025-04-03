package incompatible

import (
	"go/token"
	"reflect"
	"testing"
)

func TestAnalyseMod(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name          string
		path          string
		expectResults []Result
	}{
		{
			name: "contains incompatible",
			path: "testdata/incompatible/go.mod",
			expectResults: []Result{
				{
					Start:  token.Position{Filename: "testdata/incompatible/go.mod", Line: 6, Column: 2},
					End:    token.Position{Filename: "testdata/incompatible/go.mod", Line: 6, Column: 47},
					Reason: "github.com/google/martian@v2.1.0+incompatible is not a modules compatible version",
				},
			},
		},
		{
			name:          "no requires",
			path:          "testdata/norequire/go.mod",
			expectResults: nil,
		},
		{
			name:          "indirect incompatible",
			path:          "testdata/indirectincompatible/go.mod",
			expectResults: nil,
		},
		{
			name: "multiple incompatible",
			path: "testdata/multipleincompatible/go.mod",
			expectResults: []Result{
				{
					Start:  token.Position{Filename: "testdata/multipleincompatible/go.mod", Line: 6, Column: 2},
					End:    token.Position{Filename: "testdata/multipleincompatible/go.mod", Line: 6, Column: 47},
					Reason: "github.com/google/martian@v2.1.0+incompatible is not a modules compatible version",
				},
				{
					Start:  token.Position{Filename: "testdata/multipleincompatible/go.mod", Line: 9, Column: 2},
					End:    token.Position{Filename: "testdata/multipleincompatible/go.mod", Line: 9, Column: 42},
					Reason: "example.com/me/mymod@v4.1.0+incompatible is not a modules compatible version",
				},
			},
		},
		{
			name:          "no incompatible",
			path:          "testdata/noincompatible/go.mod",
			expectResults: nil,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			results, err := AnalyseMod(test.path)
			if err != nil {
				t.Logf("expected => no error, got => %s", err.Error())
				t.Fail()
			}

			if !reflect.DeepEqual(results, test.expectResults) {
				t.Logf("expected => %s, got => %s", test.expectResults, results)
				t.Fail()
			}
		})
	}
}
