# GopherDojo Task #1 By po3rin

This CLI convert extension of images. .png, .jpeg, .jpg and .gif is only supported.

# Feature

Traveling the directory structure recursively.

## Quick Start

There are test-images in /images directory.
default behavior is to convert jpg to png in /images directory.

```
$ make
$ ./chext
```

## Flag

You can set arguments.

|  Flag  |  Description  | Default |
| ---- | ---- | --- |
|  -f  |  What images converted from  | jpg |
|  -t  |  What images converted to  | png |
|  -d  |  Designate directory that has images | . |

following is example command with flags. This command convert jpg to gif in images/sub directory.

```
./chext -f jpg -t gif -d images/sub
```

## Doc

following command show GoDoc

```
make doc
```

## 感想と課題

全パターンは Table Driven Tests を実装した。convert する全6パターンが非常に見やすくなった。しかし、mapにfor文を適用すると、書いた順番で繰り返されない仕様にハマった。for文で順番を担保するにはどうしたらいいのだろう。。

カバレッジは80%いくが他の20%はエラー処理。
エラー処理もカバーする必要がある？

サブテストとテーブル駆動テストのパターンを採用した場合、テストヘルパーを使っても呼び出し行が変わらないのであまり意味がない？？そもそも使い方が間違っているかもしれない。

## 宿題 io.Writer と io.Reader について
io.Writer io.Readerはそれぞれインターフェースであり、抽象化されている。故に振る舞いを与えられているだけなのでFile Buffer等の様々なものに対して実装されている。
ゆえに実装が使い回せるので、結果的にGo言語のシンプルさに貢献している。
