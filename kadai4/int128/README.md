# kadai4

> - JSON形式でおみくじの結果を返す
> - 結果は毎回ランダムに変わるようにする
> - 正月（1/1-1/3）だけ大吉にする
> - ハンドラのテストを書いてみる

## 実行例

```
% go run main.go
2018/07/09 10:24:36 Listening on :8000
2018/07/09 10:24:57 GET /
2018/07/09 10:25:09 GET /api/omikuji
2018/07/09 10:25:14 GET /api/omikuji
2018/07/09 10:25:16 GET /api/omikuji
```

```
% curl -v http://localhost:8000/api/omikuji
> GET /api/omikuji HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Mon, 09 Jul 2018 01:25:09 GMT
< Content-Length: 32
<
{"description":"凶","value":1}
```

```
% curl -v http://localhost:8000/api/omikuji
> GET /api/omikuji HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Mon, 09 Jul 2018 01:25:16 GMT
< Content-Length: 35
<
{"description":"中吉","value":3}
```

```
% curl -v http://localhost:8000/
> GET / HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 404 Not Found
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Mon, 09 Jul 2018 01:24:57 GMT
< Content-Length: 19
<
404 page not found
```
