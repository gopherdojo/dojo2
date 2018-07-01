分割ダウンローダー
===

Download
---

```bash
$ go run main.go ${URL}
```

example

```bash
$ go run main.go https://github.com/sawadashota/gocmd/archive/v1.0.11.tar.gz
```

### Options

- `-p`: ダウンロードプロセス数（デフォルトはコア数）
- `-t`: タイムアウト（デフォルトは10秒）
- `-o`: 出力先（デフォルトはURLのファイル名）
