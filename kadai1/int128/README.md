# kadai1

`kadai1` is a command to convert image files.


## Getting Started

```
Usage: kadai1 FILE or DIRECTORY...
  -from string
    	Source image format: auto, jpg, png, gif (default "jpg")
  -gif-colors int
    	GIF number of colors (default 256)
  -jpeg-quality int
    	JPEG quality (default 75)
  -png-compression string
    	PNG compression level: default, no, best-speed, best-compression (default "default")
  -to string
    	Destination image format: jpg, png, gif (default "png")
```

The command skips any invalid format files, for example:

```
$ kadai1 main.go photo.jpg
2018/06/19 10:36:41 main.go -> main.png
2018/06/19 10:36:41 Skipped main.go: Error while decoding file main.go: invalid JPEG format: missing SOI marker
2018/06/19 10:36:41 photo.jpg -> photo.png
```


### Examples

To convert files in the folder from JPEG to PNG:

```sh
kadai1 my-photos/
```

To convert files from PNG to 16-colors GIF:

```sh
kadai1 -from png -to gif -gif-colors 16 my-photo.png
```

To reduce size of the JPEG file:

```sh
kadai1 -to jpg -jpeg-quality 25 
```


### Development

Build:

```sh
make
```

Test:

```sh
make test
```

Show GoDoc:

```sh
make doc
```


## 課題

> 次の仕様を満たすコマンドを作って下さい
> - ディレクトリを指定する
> - 指定したディレクトリ以下のJPGファイルをPNGに変換
> - ディレクトリ以下は再帰的に処理する
> - 変換前と変換後の画像形式を指定できる
>
> 以下を満たすように開発してください
> - mainパッケージと分離する
> - 自作パッケージと標準パッケージと準標準パッケージのみ使う
> - 準標準パッケージ：golang.org/x以下のパッケージ
> - ユーザ定義型を作ってみる
> - GoDocを生成してみる
