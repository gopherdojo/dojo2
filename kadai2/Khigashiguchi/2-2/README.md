# 課題2-2: 【TRY】テストを書いてみよう
## 課題内容
- 1回目の宿題のテストを作ってみて下さい
    - テストのしやすさを考えてリファクタリングしてみる
    - テストのカバレッジを取ってみる
    - テーブル駆動テストを行う
    - テストヘルパーを作ってみる

### day1課題のレビューでの留意点チェックリスト
- [ ] goimportsで標準・サードパーティを見やすくするように
- [ ] マジックナンバーがある場合はconst定義
- [ ] MustCompileは初期化・init関数時のみ
- [ ] cliのエラー出力はos.Stderrに出すほうがいい
- [ ] 標準入出力エラー出力は、io.Writer型をもたせるとGOOD
    - https://github.com/gopherdojo/dojo2/pull/6#discussion-diff-197047150R26
- [ ] panicはしない、log.Fatalはos.Exitが呼ばれるのでmain関数以外では極力使わない。
- [ ] エラー処理を忘れていないことをチェックするように`go errcheck`をかけておこう
- [ ] 引数定義、同じ型であれば、in, out Format みたいに省略できます
- [ ] 型公開・フィールド非公開はあんまり意味がない
- [ ] エラー処理はboolでfalseとかではなく、errorを返す
- [ ] ファイルの存在チェック、`os.Stat`では不十分、`os.IsNotExist`がベター

### リファクタリング大項目
- 大枠でパッケージを分け・テストを書く
    - [ ] コマンドオプションのパース：option
        - parse
        - 拡張子フォーマットのチェック等バリデーション
    - [ ] フォーマットに応じてDecodeする：format
    - [ ] ファイル読み・書き：file
- [ ] 最終型のmainパッケージに対するテストを書く
- [ ] day1課題のレビューでの留意点チェックリストに対応する

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
