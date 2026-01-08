package imgconv

import (
	"bytes"
	"image"
	"io"
	"os"
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
		name         string
		inputPath    string
		outputFormat Format
		wantErr      bool
	}{
		{"test: jpeg->png", filepath.Join("..", "testdata", "test1.jpeg"), Format(".png"), false},
		{"test: png->png", filepath.Join("..", "testdata", "test2.png"), Format(".png"), false},
		{"test: gif->png", filepath.Join("..", "testdata", "test3.gif"), Format(".png"), false},
		{"test: broken", filepath.Join("..", "testdata", "brokenImage.png"), Format(".png"), true},
		{"test: txt", filepath.Join("..", "testdata", "test.txt"), Format(".png"), true},

		{"test: jpeg->jpeg", filepath.Join("..", "testdata", "test4.jpeg"), Format(".jpeg"), false},
		{"test: png->jpeg", filepath.Join("..", "testdata", "test5.png"), Format(".jpeg"), false},
		{"test: gif->jpeg", filepath.Join("..", "testdata", "test6.gif"), Format(".jpeg"), false},
		{"test: broken", filepath.Join("..", "testdata", "brokenImage.png"), Format(".jpeg"), true},
		{"test: txt", filepath.Join("..", "testdata", "test.txt"), Format(".jpeg"), true},

		{"test: jpeg->gif", filepath.Join("..", "testdata", "test7.jpeg"), Format(".gif"), false},
		{"test: png->gif", filepath.Join("..", "testdata", "test8.png"), Format(".gif"), false},
		{"test: gif->gif", filepath.Join("..", "testdata", "test9.gif"), Format(".gif"), false},
		{"test: broken", filepath.Join("..", "testdata", "brokenImage.png"), Format(".gif"), true},
		{"test: txt", filepath.Join("..", "testdata", "test.txt"), Format(".gif"), true},

		{"test: jpeg->nil", filepath.Join("..", "testdata", "test10.jpeg"), Format(".txt"), true},
		{"test: png->nil", filepath.Join("..", "testdata", "test11.png"), Format(".txt"), true},
		{"test: gif->nil", filepath.Join("..", "testdata", "test12.gif"), Format(".txt"), true},
		{"test: broken", filepath.Join("..", "testdata", "brokenImage.png"), Format(".txt"), true},
		{"test: txt", filepath.Join("..", "testdata", "test.txt"), Format(".txt"), true},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			in := openTestImage(t, test.inputPath)
			defer in.Close()
			inConfig, _, inErr := image.DecodeConfig(in)
			in2 := openTestImage(t, test.inputPath)
			defer in2.Close()

			var out bytes.Buffer
			err := Convert(in2, &out, test.outputFormat)
			if (err != nil) != test.wantErr {
				t.Fatalf("Convert(input, output, %q) err=%v, wantErr=%v", test.outputFormat, err, test.wantErr)
			}
			if test.wantErr {
				return
			}
			if inErr != nil {
				t.Fatalf("DecodeConfig(input) failed: %v", inErr)
			}
			outConfig, outFormat, err := image.DecodeConfig(&out)
			if err != nil {
				t.Fatalf("DecodeConfig(output) failed: %v", err)
			}
			if Format("."+outFormat) != test.outputFormat {
				t.Fatalf("output format = %q, want %q", outFormat, test.outputFormat)
			}
			if inConfig.Height != outConfig.Height || inConfig.Width != outConfig.Width {
				t.Fatalf("size doesn't match: input=(%d, %d), output=(%d, %d)",
					inConfig.Width, inConfig.Height, outConfig.Width, outConfig.Height)
			}
		})
	}
}

/*
ヘルパー関数は「テストの前提条件」を作る関数
- テストデータをopenする
- 失敗したらテストを中断する
→ 「このテストが、この画像が読めることを前提とするという前提条件の宣言」

r := openTestImage(...)（ケースごとに）

w := new buffer

関数呼び出し

w の中身を image.Decode 等で検証
*/
func openTestImage(t *testing.T, inputPath string) io.ReadCloser {
	t.Helper()
	r, err := os.Open(inputPath)
	if err != nil {
		t.Fatalf("error: %s\n", err)
	}
	return r
}

func TestParseFormat(t *testing.T) {
	var tests = []struct {
		name    string
		input   string
		want    Format
		wantErr bool
	}{
		{"test: jpg", "jpg", Format(".jpg"), false},
		{"test: jpeg", "jpeg", Format(".jpg"), false},
		{"test: png", "png", Format(".png"), false},
		{"test: gif", "gif", Format(".gif"), false},
		{"test: (empty string)", "", Format(""), true},
		{"test: txt", "txt", Format(""), true},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := ParseFormat(test.input)
			if err != nil && test.wantErr == false {
				t.Fatalf("ParseFormat(%q) err=%v, wantErr=%v", test.input, err, test.wantErr)
			}
			if test.wantErr {
				return
			}
			if got != test.want {
				t.Fatalf("ParseFormat(%q) got=%q, want=%q", test.input, got, test.want)
			}
		})
	}
}
