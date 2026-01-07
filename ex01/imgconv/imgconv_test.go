package imgconv

import (
	"path/filepath"
	"testing"
)

/*TODO
- test helperの理解
- test Parallellの理解
*/
/*
- inputPath
- outputForamt
- wantErr
- wantWidth / wantHeight
*/
func TestConvert(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		inputPath    string
		outputFormat Format
		wantErr      bool
	}{
		{filepath.Join("..", "testdata", "test.jpeg"), Format(".png"), false},
		{filepath.Join("..", "testdata", "test.png"), Format(".png"), false},
		{filepath.Join("..", "testdata", "test.gif"), Format(".png"), false},
		{filepath.Join("..", "testdata", "brokenImage.png"), Format(".png"), true},
		{filepath.Join("..", "testdata", "test.txt"), Format(".png"), true},
	}
	for _, test := range tests {
		err := CheckConvert(test)
		if err != nil {
			t.Errorf("error")
		}

	}
}

/*
-
*/
func testCheckConvert(t *testing.T) string {
	t.Helper()
	// tf, err :=
}

func TestParseFormat(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		input   string
		want    Format
		wantErr bool
	}{
		{"jpg", Format(".jpg"), false},
		{"jpeg", Format(".jpg"), false},
		{"png", Format(".png"), false},
		{"gif", Format(".gif"), false},
		{"", Format(""), true},
		{"txt", Format(""), true},
	}
	for _, test := range tests {
		if got, _ := ParseFormat(test.input); got != test.want {
			t.Errorf("ParseFormat(%q) = %v", test.input, got)
		}
	}
}
