package imgconverter

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

// assert用の関数を定義(ステータス確認用)
func assertStatus(t *testing.T, status int, expected int) {
	if status != expected {
		t.Errorf("Status expected %v to eq %v", status, expected)
	}
}

// assert用の関数を定義(ファイル存在確認用)
func assertFile(t *testing.T, path string, expectedFilename string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		if f.Name() == expectedFilename {
			return
		}
	}
	t.Error("Expected File Not Found.")
}

// assert用の関数を定義(ブーリアン確認用)
func assertBool(t *testing.T, output bool, expected bool) {
	if output != expected {
		t.Errorf("Output expected %v to eq %v", output, expected)
	}
}

// テスト用にファイルをコピーする関数
func copyFile(srcName string, dstName string) {
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}

// checkcheckUnacceptableFormat関数のテスト
// .jpg, .png, .gifを入力したときは、falseを返却
func TestCheckUnacceptableFormatWithAcceptableFormat(t *testing.T) {
	cases := []struct {name string; input string; expected bool}{
		{name: "jpg", input: ".jpg", expected: false},
		{name: "png", input: ".png", expected: false},
		{name: "gif", input: ".gif", expected: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := checkUnacceptableFormat(c.input); c.expected != actual {
				t.Errorf("Output expected %v to eq %v", actual, c.expected)
			}
		})
	}
}

// checkcheckUnacceptableFormat関数のテスト
// .jpg, .png, .gif以外を入力したときは、trueを返却
func TestCheckUnacceptableFormatWithUnacceptableFormat(t *testing.T) {
	cases := []struct {name string; input string; expected bool}{
		{name: "txt", input: ".txt", expected: true},
		{name: "doc", input: ".doc", expected: true},
		{name: "muga", input: "muga", expected: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := checkUnacceptableFormat(c.input); c.expected != actual {
				t.Errorf("Output expected %v to eq %v", actual, c.expected)
			}
		})
	}
}

// インプット JPGファイル 10枚　ディレクトリ0:ディレクトリ下ファイルなし
// 変換元 JPG 変換後 PNG
// 想定結果 JPGファイル 10枚がPNGに変換されている
func TestCliFromJpgToPngWith10pictures(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{OutStream: outStream, ErrStream: errStream}
	checkPath := "../testdata/testImg/jpg1/"

	args := strings.Split("imgConverter -src jpg -dst png "+checkPath, " ")

	// ファイルをコピーする
	for i := 1; i < 10; i++ {
		copyFile(checkPath+strconv.Itoa(i)+".jpg", checkPath+strconv.Itoa(i+1)+".jpg")
	}

	// cliを実行、ステータスを確認
	assertStatus(t, cli.Run(args), ExitCodeOK)

	// 変換後のファイルをチェックして削除する
	for i := 1; i <= 10; i++ {
		convertedFileName := strconv.Itoa(i) + ".png"

		assertFile(t, checkPath, convertedFileName)

		// 変換されたファイルを削除
		if err := os.Remove(checkPath + convertedFileName); err != nil {
			t.Fatal("Failed to remove the file")
		}
	}

	// コピーしたファイルをチェックして削除
	for i := 2; i <= 10; i++ {
		convertFileName := strconv.Itoa(i) + ".jpg"

		// 変換前のファイルを削除
		if err := os.Remove(checkPath + convertFileName); err != nil {
			t.Fatal("Failed to remove the file")
		}
	}
}

// インプット JPGファイル 10枚　ディレクトリ0:ディレクトリ下ファイルなし
// 変換元 JPG 変換後 PNG
// ※フラグの指定なし
// 想定結果 JPGファイル 10枚がPNGに変換されている
func TestCliFromJpgToPngWith10picturesWithNoFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{OutStream: outStream, ErrStream: errStream}
	checkPath := "../testdata/testImg/jpg1/"

	args := strings.Split("imgConverter "+checkPath, " ")

	// ファイルをコピーする
	for i := 1; i < 10; i++ {
		copyFile(checkPath+strconv.Itoa(i)+".jpg", checkPath+strconv.Itoa(i+1)+".jpg")
	}

	// cliを実行、ステータスを確認
	assertStatus(t, cli.Run(args), ExitCodeOK)

	// 変換後のファイルをチェックして削除する
	for i := 1; i <= 10; i++ {
		convertedFileName := strconv.Itoa(i) + ".png"

		assertFile(t, checkPath, convertedFileName)

		// 変換されたファイルを削除
		if err := os.Remove(checkPath + convertedFileName); err != nil {
			t.Fatal("Failed to remove the file")
		}
	}

	// コピーしたファイルをチェックして削除
	for i := 2; i <= 10; i++ {
		convertFileName := strconv.Itoa(i) + ".jpg"

		// 変換前のファイルを削除
		if err := os.Remove(checkPath + convertFileName); err != nil {
			t.Fatal("Failed to remove the file")
		}
	}
}

// インプット PNGファイル 10枚　ディレクトリ0:ディレクトリ下ファイルなし
// 変換元 PNG 変換後 GIF
// 想定結果 PNGファイル 10枚がGIFに変換されている
func TestCliFromPngToGifWith10pictures(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{OutStream: outStream, ErrStream: errStream}
	checkPath := "../testdata/testImg/png1/"

	args := strings.Split("imgConverter -src png -dst gif "+checkPath, " ")

	// ファイルをコピーする
	for i := 1; i < 10; i++ {
		copyFile(checkPath+strconv.Itoa(i)+".png", checkPath+strconv.Itoa(i+1)+".png")
	}

	// cliを実行、ステータスを確認
	assertStatus(t, cli.Run(args), ExitCodeOK)

	// 変換後のファイルをチェックして削除する
	for i := 1; i <= 10; i++ {
		convertedFileName := strconv.Itoa(i) + ".gif"

		assertFile(t, checkPath, convertedFileName)

		// 変換されたファイルを削除
		if err := os.Remove(checkPath + convertedFileName); err != nil {
			t.Fatal("Failed to remove the file")
		}
	}

	// コピーしたファイルをチェックして削除
	for i := 2; i <= 10; i++ {
		convertFileName := strconv.Itoa(i) + ".png"

		// 変換前のファイルを削除
		if err := os.Remove(checkPath + convertFileName); err != nil {
			t.Fatal("Failed to remove the file")
		}
	}
}

// インプット GIFファイル 10枚　ディレクトリ0:ディレクトリ下ファイルなし
// 変換元 GIF 変換後 JPG
// 想定結果 GIFファイル 10枚がJPGに変換されている
func TestCliFromPngToGifWith10picture(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{OutStream: outStream, ErrStream: errStream}
	checkPath := "../testdata/testImg/png1/"

	args := strings.Split("imgConverter -src png -dst gif "+checkPath, " ")

	// ファイルをコピーする
	for i := 1; i < 10; i++ {
		copyFile(checkPath+strconv.Itoa(i)+".png", checkPath+strconv.Itoa(i+1)+".png")
	}

	// cliを実行、ステータスを確認
	assertStatus(t, cli.Run(args), ExitCodeOK)

	// 変換後のファイルをチェックして削除する
	for i := 1; i <= 10; i++ {
		convertedFileName := strconv.Itoa(i) + ".gif"

		assertFile(t, checkPath, convertedFileName)

		// 変換されたファイルを削除
		if err := os.Remove(checkPath + convertedFileName); err != nil {
			t.Fatal("Failed to remove the file")
		}
	}

	// コピーしたファイルをチェックして削除
	for i := 2; i <= 10; i++ {
		convertFileName := strconv.Itoa(i) + ".png"

		// 変換前のファイルを削除
		if err := os.Remove(checkPath + convertFileName); err != nil {
			t.Fatal("Failed to remove the file")
		}
	}
}

// インプット JPGファイル 10枚　ディレクトリ1コ:ディレクトリ下ファイル1枚
// 変換元 JPG 変換後 PNG
// 想定結果 JPGファイル 10枚がPNGに変換されている
// TODO: テストコードを書く

// インプット JPGファイル 10枚　ディレクトリ1コ:ディレクトリ下ファイル10枚
// 変換元 JPG 変換後 PNG
// 想定結果 JPGファイル 10枚がPNGに変換されている
// TODO: テストコードを書く

// インプット JPGファイル 1枚（正しいJPG形式ではなく、形式がおかしい）　ディレクトリ0
// 変換元 JPG 変換後 PNG
// 想定結果 エラー終了する
func TestCliErrorFileWith1pictures(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{OutStream: outStream, ErrStream: errStream}
	checkPath := "../testdata/testImg/error1/"

	args := strings.Split("imgConverter -src jpg -dst png "+checkPath, " ")

	// cliを実行、ステータスを確認
	assertStatus(t, cli.Run(args), ExitCodeNG)
}
