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

## 使い方
サンプルURL: http://i.imgur.com/z4d4kWk.jpg


## 実装内容
### 処理フロー
- [x] cliを作る（hoge.New())
- cli
    - [x] コマンド引数からURLを取得する
        - [ ] url形式かどうか等がもしかしたら必要
    - [x] コマンド引数バリデーション
- request header
    - [x] headerリクエスト
- download
    - [ ] rangeアクセス
    - ...etc

### Better Todo
- [ ] io.Reader型を使って、テストのときは外部URLではなくローカルファイルでテストできるようにしておくと良い

### 雑記メモ
- `GOMAXPROCS`をどこで挟み込んでいくか
- `context.WithTimeout`のはさみどころがポイントなんだろう
- `http.header`を雑にした標準出力したときの結果

```
-> % go run kadai3-2/Khigashiguchi/sget/cmd/sget/main.go http://i.imgur.com/z4d4kWk.jpg
&{200 OK 200 HTTP/1.1 1 1 map[Access-Control-Allow-Origin:[*] X-Served-By:[cache-iad2135-IAD,cache-sjc3636-SJC] Access-Control-Allow-Methods:[GET, OPTIONS] Accept-Ranges:[bytes] X-Cache:[HIT, HIT] Etag:["18c50e1bbe972bdf9c7d9b8f6f019959"] Content-Length:[146515] X-Cache-Hits:[1, 1] Server:[cat factory 1.0] Date:[Mon, 02 Jul 2018 16:47:11 GMT] Connection:[keep-alive] Content-Type:[image/jpeg] Cache-Control:[public, max-age=31536000] Age:[208894] X-Timer:[S1530550031.455210,VS0,VE3] Last-Modified:[Thu, 02 Feb 2017 11:15:53 GMT] X-Amz-Storage-Class:[STANDARD_IA]] {} 146515 [] false false map[] 0xc4200f4000 <nil>}
```

## 参考
- https://qiita.com/codehex/items/d0a500ac387d39a34401
- http://tnamao.hatenablog.com/entry/20100617/p1
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Range_requests