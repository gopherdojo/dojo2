# 課題1 - 画像変換コマンドを作ろう -

以下の仕様を満たすコマンドを作る。

- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる

また以下を満たすように実装すること

- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
    - 準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる

# 実行手順

以下のコマンドでビルドする。

```
# ビルド
go build -o myconverter

# 出力用フォルダを作成する
mkdir outputs
```

ディレクトリは`-target`で指定し、`-from`で変更前の拡張子、`-to`で変更後の拡張子を指定する。変換されたファイルはoutputフォルダに出力される。

```
# images配下のjpgファイルをpngへ変換する
$ ./myconverter -target=images -from=jpg -to=png
Convert Success from images/dir/sub_lena.jpg to outputs/sub_lena.png
Convert Success from images/lena.jpg to outputs/lena.png

# images配下のpngファイルをjpgへ変換する
$ ./myconverter -target=images -from=png -to=jpg
Convert Success from images/dir/sub_icon.png to outputs/sub_icon.jpg
Convert Success from images/icon.png to outputs/icon.jpg
```

デフォルトでは`target`はカレントディレクトリが指定され、`jpg`から`png`への変換を実行する。
```
$ ./myconverter 
Convert Success from current_rena.jpg to outputs/current_rena.png
Convert Success from images/dir/sub_lena.jpg to outputs/sub_lena.png
Convert Success from images/lena.jpg to outputs/lena.png
```

`jpg`から`png`、`png`から`jpg`への変換のみサポートしており、その他の変換には失敗する。
```
$ ./myconverter -target=images -from=gif -to=jpg
[ERROR]No supprot to convert from gif to jpg
```

ヘルプコマンド
```
# ヘルプコマンドの実行例
$ ./myconverter -h
Usage of ./myconverter:
  -from string
    	from file extention (default "jpg")
  -target string
    	target filepath. (default ".")
  -to string
    	to file extention (default "png")
```

godocを生成してみる
```
$godoc -html ./converter > godoc-converter.html
```

---

# 所感

- 以前入れ忘れたVSCodeのGoプラグインを入れた結果、GoLintによってコンパイル実行前に型エラーに気づけたり、GoDocの記入漏れに気づくことができた
- まだGoに書きなれてないので文法をググりながらやったので時間がかかってしまった
- GoDocの生成のときにgithub.comから参照しようとしたがブランチを参照する方法が分からなかった
