# 課題2-1

## 課題

* io.Readerとio.Writerについて調べてみよう
    * 標準パッケージでどのように使われているか
    * io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

## 回答

参照したコード

* https://github.com/golang/go/blob/master/src/io/io.go
* https://github.com/golang/go/blob/master/src/bytes/buffer.go
* https://github.com/golang/go/blob/master/src/bytes/reader.go
* https://github.com/golang/go/blob/master/src/archive/zip/reader.go
* https://github.com/golang/go/blob/master/src/archive/zip/writer.go

### どのように使われているか

* io.go
    * io.goRead、Write、Close、Seek等に対してそれぞれInterfaceが定義されている
    * ReadWriter、ReadCloser等複数の振る舞いに対しては型埋め込みが使われている
* bytes/buffer.go, bytes/reader.go
    * io.goでintefaceとして定義されたByteReaderのメソッド `ReadByte` の振る舞いをそれぞれ記述していた
* zip/reader.go, zip/writer.go
    * ...

### どういう利点があるのか

* ???

---

# 課題2-2

## 課題

* 1回目の宿題のテストを作ってみて下さい
    * テストのしやすさを考えてリファクタリングしてみる
    * テストのカバレッジを取ってみる
    * テーブル駆動テストを行う
    * テストヘルパーを作ってみる

## 回答

W.I.P...
