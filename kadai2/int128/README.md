# kadai2

See also https://github.com/gopherdojo/dojo2/tree/kadai1-int128/kadai1/int128.


## `io.Reader` と `io.Writer`

`io.Reader` と `io.Writer` はストリームの読み書きを行うためのインタフェースで、Javaにおける `InputStream` や `OutputStream` に相当する。

### 標準パッケージにおける利用

Go 1.10では `io.Reader` と `io.Writer` は以下のように定義されている。

```go
package io

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
```

Go 1.10の標準パッケージでは16個の構造体が `Read([]byte)` メソッドを実装している。
また、15個の構造体が `Write([]byte)` メソッドを実装している。
（テストコードおよび `internal` パッケージを除く）

具体的には以下のメソッドが存在する。

```go
% make showReaderWriterImplements
./bufio/bufio.go                  func (b *Reader) Read(p []byte) (n int, err error) {
./bufio/bufio.go                  func (b *Writer) Write(p []byte) (nn int, err error) {
./crypto/cipher/io.go             func (r StreamReader) Read(dst []byte) (n int, err error) {
./crypto/cipher/io.go             func (w StreamWriter) Write(src []byte) (n int, err error) {
./crypto/tls/conn.go              func (c *Conn) Write(b []byte) (int, error) {
./crypto/tls/conn.go              func (c *Conn) Read(b []byte) (n int, err error) {
./compress/flate/deflate.go       func (w *Writer) Write(data []byte) (n int, err error) {
./compress/gzip/gzip.go           func (z *Writer) Write(p []byte) (int, error) {
./compress/gzip/gunzip.go         func (z *Reader) Read(p []byte) (n int, err error) {
./compress/zlib/writer.go         func (z *Writer) Write(p []byte) (n int, err error) {
./strings/reader.go               func (r *Reader) Read(b []byte) (n int, err error) {
./strings/builder.go              func (b *Builder) Write(p []byte) (int, error) {
./net/net.go                      func (v *Buffers) Read(p []byte) (n int, err error) {
./net/http/httptest/recorder.go   func (rw *ResponseRecorder) Write(buf []byte) (int, error) {
./archive/tar/writer.go           func (tw *Writer) Write(b []byte) (int, error) {
./archive/tar/reader.go           func (tr *Reader) Read(b []byte) (int, error) {
./bytes/buffer.go                 func (b *Buffer) Write(p []byte) (n int, err error) {
./bytes/buffer.go                 func (b *Buffer) Read(p []byte) (n int, err error) {
./bytes/reader.go                 func (r *Reader) Read(b []byte) (n int, err error) {
./io/io.go                        func (l *LimitedReader) Read(p []byte) (n int, err error) {
./io/io.go                        func (s *SectionReader) Read(p []byte) (n int, err error) {
./io/pipe.go                      func (r *PipeReader) Read(data []byte) (n int, err error) {
./io/pipe.go                      func (w *PipeWriter) Write(data []byte) (n int, err error) {
./math/rand/rand.go               func (r *Rand) Read(p []byte) (n int, err error) {
./log/syslog/syslog.go            func (w *Writer) Write(b []byte) (int, error) {
./mime/multipart/multipart.go     func (p *Part) Read(d []byte) (n int, err error) {
./mime/quotedprintable/writer.go  func (w *Writer) Write(p []byte) (n int, err error) {
./mime/quotedprintable/reader.go  func (r *Reader) Read(p []byte) (n int, err error) {
./os/file.go                      func (f *File) Read(b []byte) (n int, err error) {
./os/file.go                      func (f *File) Write(b []byte) (n int, err error) {
./text/tabwriter/tabwriter.go     func (b *Writer) Write(buf []byte) (n int, err error) {
```

メソッドの役割をまとめるとおおよそ以下のようになる。

- ファイルの読み書き
- ネットワーク通信
- 暗号化、復号
- ファイルの圧縮、展開（ZIP/TAR）
- MIMEエンコード、デコード
- バイト配列や文字列の処理
- 行指向やトークン分割の処理

このように、Goの標準パッケージでは入出力に関わるインタフェースが抽象化されていることが分かる。

### 抽象化の利点

入出力に関わるインタフェースを抽象化することで、コードをシンプルに保ちながら拡張性を持たせることができる。

例えば、 `image/jpeg` パッケージでは以下のメソッドが定義されている。

```go
func Decode(r io.Reader) (image.Image, error) {}
```

ローカルにあるJPEGファイルを読み込みたい場合は `os.Open()` の戻り値を渡せばよい。

```go
func Example_io_Reader_File() {
	f, err := os.Open("image.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("size=%+v", img.Bounds())
	// Output: size=(0,0)-(1000,750)
}
```

また、リモートにあるJPEGファイルを読み込みたい場合は `http.Get()` の戻り値を渡せばよい。

```go
func Example_io_Reader_HTTP() {
	resp, err := http.Get("https://upload.wikimedia.org/wikipedia/commons/b/b2/JPEG_compression_Example.jpg")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("size=%+v", img.Bounds())
	// Output: size=(0,0)-(1000,750)
}
```

もちろん、独自に定義した型を渡すこともできる。

```go
type DummyReader struct{}

func (r *DummyReader) Read(p []byte) (int, error) {
	return 0, io.EOF
}

func Example_io_Reader_DummyReader() {
	jpeg.Decode(&DummyReader{})
}
```

このように、インタフェースによる抽象化を行うことで、JPEGデータがローカルにある場合でもリモートにある場合でも同じメソッドを使うことができる。

もし、インタフェースが使えない場合は、以下のように具象型ごとに関数を定義することになる。

```go
func DecodeFile(f *os.File) (image.Image, error) {}
func DecodeHTTPResponseBody(r /* レスポンスボディ型 */) (image.Image, error) {}
func DecodeZIPFile(f *zip.File) (image.Image, error) {}
```

これでは具象型が増えるたびに関数を定義する必要があり、冗長なコードが増えてしまう。
また、標準パッケージの外側で独自に定義した型を受け取ることができない問題がある。


## kadai1のリファクタリングとテスト

前回の課題1でほとんどのテストコードを書いていたため、課題2では `main_test.go` を追加しました。


## 課題2

> io.Readerとio.Writerについて調べてみよう
>
> - 標準パッケージでどのように使われているか
> - io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
>
> 1回目の宿題のテストを作ってみて下さい
>
> - テストのしやすさを考えてリファクタリングしてみる
> - テストのカバレッジを取ってみる
> - テーブル駆動テストを行う
> - テストヘルパーを作ってみる
