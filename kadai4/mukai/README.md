# 実行例
```
 curl -v http://localhost:8080
* Rebuilt URL to: http://localhost:8080/
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Wed, 11 Jul 2018 16:57:05 GMT
< Content-Length: 19
<
{"data":"大吉"}

curl -v http://localhost:8080
* Rebuilt URL to: http://localhost:8080/
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Wed, 11 Jul 2018 16:57:14 GMT
< Content-Length: 16
<
{"data":"凶"}

```

# 指定日のテスト
おみくじをひく日時を返却するメソッドを持ったインタフェースをHandlerにもたせる.

本番実装では、実行時の日地を返却するメソッドを持った構造体を渡し、テスト時には任意の日時を返却する構造体を渡すことで、実装を変えずに任意の日付のおみくじ結果を取得できる.

