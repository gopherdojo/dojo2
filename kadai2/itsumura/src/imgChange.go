/*
ディレクトリを指定する
指定したディレクトリ以下のJPGファイルをPNGに変換
ディレクトリ以下は再帰的に処理する
変換前と変換後の画像形式を指定できる

mainパッケージと分離する
自作パッケージと標準パッケージと準標準パッケージのみ使う
準標準パッケージ：golang.org/x以下のパッケージ
ユーザ定義型を作ってみる
GoDocを生成してみる
*/

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"./changefunc"
)

//エラーハンドリング
func ec(err *error) {
	if *err != nil {
		fmt.Println(*err)
		panic(*err)
	}
}

//ユーザー定義型
type fileName string

//拡張子を表示
func (fileName fileName) Extension() string {
	return strings.Split(string(fileName), ".")[1]
}

func main() {

	for {
		fmt.Println("以下の中から画像を変換するディレクトリを選んでください")

		var fileNameSlice []string
		files, _ := ioutil.ReadDir("./") //カレントディレクトリをls
		for _, file := range files {
			fileNameSlice = append(fileNameSlice, file.Name())
		}
		fmt.Println(fileNameSlice) //カレントディレクトリのファイル一覧を表示

		var input string
		fmt.Scan(&input) //標準入力から文字列取得

		//検索
		var err error
		results, err := searchDir(input)
		//エラーが起きたらもう一度
		if err != nil {
			fmt.Println("正しく入力してください")
			continue
		}

		fmt.Println(results) //選択したディレクトリ以下のファイルを全表示
		check := true

		for _, result := range results {
			filename := fileName(result)

			if filename.Extension() != "jpg" &&
				filename.Extension() != "jpeg" &&
				filename.Extension() != "png" {
				check = false
			}
			/*newFileNameSlice := strings.Split(result, ".")

			if newFileNameSlice[len(newFileNameSlice)-1] != "jpg" &&
				newFileNameSlice[len(newFileNameSlice)-1] != "jpeg" &&
				newFileNameSlice[len(newFileNameSlice)-1] != "png" {
				check = false
			}
			*/
		}
		if check == false {
			fmt.Println("変換可能な画像ファイルがありません")
			continue
		}

		//resultに入ってるパスの画像を変換
		input = ""
		for {
			fmt.Println("変換先のファイル形式を選んで下さい")
			fmt.Println("jpg or png")

			fmt.Scan(&input) //標準入力から文字列取得

			if input != "jpg" &&
				input != "png" {
				fmt.Println("正しく入力してください")
				continue
			} else {
				break
			}
		}
		fmt.Println(input + "を選択") //変換ファイル形式

		//画像変換処理
		success := 0
		if input == "jpg" {
			success = changefunc.PngToJpg(results, input)
		} else if input == "png" {
			success = changefunc.JpgToPng(results, input)
		}

		if success > 0 {
			fmt.Println("完了")
			break
		} else {
			continue
		}
	}
}

func searchDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			var path []string
			path, err = searchDir(filepath.Join(dir, file.Name()))
			paths = append(paths, path...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths, err
}
