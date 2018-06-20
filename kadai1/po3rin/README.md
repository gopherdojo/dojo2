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
ユーザー定義型として構造体を定義したが適切な使い方がイマイチつかめない。関数に レシーバで渡す or 引数で渡すを、どちらを使うかを適切に判断したい。

全ての変換パターンのテストを書きたかった為、下記のようにレジーバに渡す構造体をfor文で回して作りたかったが上手くできず。テストを実行すると "gopher.pnggif" 的な名前のものが出てくる。レシーバに渡した値がどっか別のところで書き換えられている？？

```go
package extension

import (
	"os/exec"
	"testing"
)

func TestConvert(t *testing.T) {
	type item [3]string
	type any [6]item
	a := any{
		item{"jpg","png", "../images/gopher.jpg"}, 
		item{"png", "jpg", "../images/gopher.png"},
		item{"jpg", "gif", "../images/gopher.png"},
		item{"jpg", "png", "../images/sub/gopher-sub.jpg"},
		item{"png", "gif", "../images/sub/gopher-sub.png"},
		item{"gif", "jpg", "../images/sub/gopher-sub.png"},
	}
	for i := 0; i < len(a); i++{
		arg := Arg{
			From: a[i][0],
			To:   a[i][1],
			Path: a[i][2],
		}
		t.Logf("%s convert %s to %s",  a[i][2], a[i][0], a[i][1])
		err := arg.Convert()
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
    }
    // // clean up
	// exec.Command("rm", "../images/gopher.png")
}
```