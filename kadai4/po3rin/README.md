# Gopher道場課題 #4

## description

GET /omikuzi --> 今日の運勢を占う(大吉/中吉/小吉/吉/凶/大凶)
* 1/1 ~ 1/3 のみ大吉になります。

## Quick Start

### test

```bash
go test -v -count=1 ./handler/
```
### sever start

```bash
$ go run main.go
```

### request API

```bash
$ curl localhost:8080/omikuzi
{"code":0,"result":"大吉"}
```

# Issue
なぜかtestだけrandが固定される。
=> testがchacheされてるだけでした。。