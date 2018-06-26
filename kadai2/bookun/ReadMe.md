## io.Reader と io.Writer の使われ方

### 標準パッケージでどのように使われているか
一例として、fmtパッケージのFscanfメソッドの引数では io.Readerが、Fprintf メソッドの引数ではio.Writerが使用されている。
この利点について次節で説明する。

### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
io.Readerとio.Writerを引数として持つメソッドというのは、ReadとWriteメソッドを実装しているデータ型であれば引数にいれてよいということである。File型でもRead()およびWrite()を実装していて、それはbufferや標準出力などでも同じである。
つまり、例えばio.Writerを引数として持つメソッドには引数としてos.StdoutやFileやbufferを与えても良いということになる。
これが可能なことでプログラマーは〇〇から入力を受けるや〇〇に対して出力をする場合に、fmt.Fscanfやfmt.Fprintfを用いればよいことになり、対象が異なるデータ型であっても大きくコードを変更しなくて良い。




## 課題1のテストを作ってみる
* テストのしやすさを考えてリファクタリングしてみる
* テストのカバレッジを取ってみる
* テーブル駆動テストを行う
* テストヘルパーを作ってみる

## build 

`go build -o convert-cli cli.go`

## 使い方

`./convert-cli dir_path`
dir_path 内のjpg画像がpngに変換され、dir_path内にpngファイルが追加で作成される

`./convert-cli -s png -d jpeg dir_path`
dir_path 内のsオプションで指定したフォーマットの画像がdオプションで指定されたフォーマットに変換され、dir_path内に変換後のファイルが追加で作成される

## 補足
対応している画像の形式は jpeg, jpg, png のみ。

## 課題2-1
kadai2.md

