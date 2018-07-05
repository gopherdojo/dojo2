# Quick start

you can exex test file server via below command

```bash
$ cd server
$ go run main.go
```

exec splitt loader

```bash
$ cd loader
$ go run main.go
```

##　反省
まだキャンセルの実装ができていない。
errgroupによるエラー処理を追加したが、全てのgorutineがちゃんととまているのかテストする術を調査中
はじめてHeadメソッドを使ったので勉強になった。
