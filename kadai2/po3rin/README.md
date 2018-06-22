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

全パターンは Table Driven Tests を実装した。
convert する全6パターンが非常に見やすくなった。
