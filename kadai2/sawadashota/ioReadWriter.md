`io.Reader`と`io.Writer`について調べてみよう
===

標準パッケージでどのように使われているか
---

* 関数の引数として
* 構造体の埋め込み
* 構造体のフィールドの型として
* インターフェースへの埋め込み

io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
---

### テストするときに、`ioutil.Discard`に差し替えることができる

テスト時に`io.Writer`として`ioutil.Discard`を使うことで、実際の何もしない`Write`関数を使うことができる

`http_test`パッケージの例

```go
func TestIssue10884_MaxBytesEOF(t *testing.T) {
	dst := ioutil.Discard
	_, err := io.Copy(dst, MaxBytesReader(
		responseWriterJustWriter{dst},
		ioutil.NopCloser(delayedEOFReader{strings.NewReader("12345")}),
		5))
	if err != nil {
		t.Fatal(err)
	}
}
```

`ioutil.Discard`の実装

```go
// Discard is an io.Writer on which all Write calls succeed
// without doing anything.
var Discard io.Writer = devNull(0)
```
