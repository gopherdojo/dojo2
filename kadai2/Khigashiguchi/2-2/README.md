# 課題2-2: 【TRY】テストを書いてみよう
## 課題内容
- 1回目の宿題のテストを作ってみて下さい
    - [x] テストのしやすさを考えてリファクタリングしてみる
    - [x] テストのカバレッジを取ってみる
    - [x] テーブル駆動テストを行う
    - [x] テストヘルパーを作ってみる

### day1課題のレビューでの留意点チェックリスト
- [x] goimportsで標準・サードパーティを見やすくするように
- [x] マジックナンバーがある場合はconst定義
- [x] MustCompileは初期化・init関数時のみ
- [x] cliのエラー出力はos.Stderrに出すほうがいい
- [x] 標準入出力エラー出力は、io.Writer型をもたせるとGOOD
    - https://github.com/gopherdojo/dojo2/pull/6#discussion-diff-197047150R26
- [x] panicはしない、log.Fatalはos.Exitが呼ばれるのでmain関数以外では極力使わない。
- [x] エラー処理を忘れていないことをチェックするように`go errcheck`をかけておこう
- [x] 引数定義、同じ型であれば、in, out Format みたいに省略できます
- [x] 型公開・フィールド非公開はあんまり意味がない
- [x] エラー処理はboolでfalseとかではなく、errorを返す
- [x] ファイルの存在チェック、`os.Stat`では不十分、`os.IsNotExist`がベター

### リファクタリング大項目
- 大枠でパッケージを分け・テストを書く
    - [x] コマンドオプションのパース：option
        - [x] parse
        - [x] 拡張子のチェック等バリデーション
    - [x] 拡張子に応じてEncode/Decodeする：format
    - [x] ファイル読み・書き：file
- [ ] 最終型のmainパッケージに対するテストを書く
- [x] day1課題のレビューでの留意点チェックリストに対応する

### errcheck結果
- fmt.Fprintf・deferで実行するやつらが指摘された、fmt.Fprintf系はエラーを返す系だから共通で一箇所でエラーハンドリングしたほうが良さそう

```
-> % errcheck ./...
cli.go:23:14:   fmt.Fprintf(os.Stderr, "parse errors caused. %s.\n", err)
cli.go:28:14:   fmt.Fprintf(os.Stderr, "Failed to get decoder")
cli.go:33:14:   fmt.Fprintf(os.Stderr, "Failed to get encoder")
conversion/conversion.go:32:16: defer sf.Close()
conversion/conversion.go:41:16: defer wf.Close()
format/format_test.go:20:22:    defer jpegFile.Close()
format/format_test.go:21:17:    defer os.Remove(jpegFile.Name())
```

### テストカバレッジ結果
2018/7/1 18:38 計測

```
-> % make test
go test ./... -cover
?       github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2    [no test files]
ok      github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/conversion 0.008s  coverage: 17.6% of statements
ok      github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/format     0.013s  coverage: 25.0% of statements
ok      github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/options    0.008s  coverage: 94.4% of statements
```

### 雑記メモ
- os.Argsは型ではなくパッケージ変数なんですね
    - `var Args []string`
- flags.NewFlagSet
    - ContinueOnErrorの指定で、panicではなくerrorを返す、mainの中でやらないのでこの指定が有効そう
    - `// Return a descriptive error.`
- flags.Usageは`func()`定義だから関数型宣言？
    - `Usage is the function called when an error occurs while parsing flags.`
    - parseのエラー時に何をするか
- 構造体同士をphpのノリでまるっと比較しようとしたけど、メモリ番地違うなどでフィールドごとにやったほうがいい感じがあった
- slice型同士の比較もループで回してアサーションって感じなのかな

### 参考
- day1のint128さんのコード勉強させてもらいました
    - https://github.com/gopherdojo/dojo2/pull/7
- それぞれのテストパターンの参考
    - https://qiita.com/nirasan/items/b357f0ad9172ab9fa19b
