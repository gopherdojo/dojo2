# kadai3-2
## 課題内容
- 分割ダウンロードを行う
    - Rangeアクセスを用いる
    - いくつかのゴルーチンでダウンロードしてマージする
    - エラー処理を工夫する
        - golang.org/x/sync/errgourpパッケージなどを使ってみる
    - キャンセルが発生した場合の実装を行う

### 講義中に紹介されたコード
- https://github.com/Code-Hex/pget

### 処理フロー
- [x] cliを作る（hoge.New())
- cli
    - [x] コマンド引数からURLを取得する
        - [ ] url形式かどうか等がもしかしたら必要
    - [x] コマンド引数バリデーション
- download
    - [ ] rangeアクセス
    - ...etc

### Better
- [ ] io.Reader型を使って、テストのときは外部URLではなくローカルファイルでテストできるようにしておくと良い

### 雑記メモ
- `GOMAXPROCS`をどこで挟み込んでいくか
- `context.WithTimeout`のはさみどころがポイントなんだろう

## 参考
- https://qiita.com/codehex/items/d0a500ac387d39a34401
- http://tnamao.hatenablog.com/entry/20100617/p1