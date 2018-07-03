# kadai3-2

> 分割ダウンロードを行う
> 
> - Rangeアクセスを用いる
> - いくつかのゴルーチンでダウンロードしてマージする
> - エラー処理を工夫する: golang.org/x/sync/errgourpパッケージなどを使ってみる
> - キャンセルが発生した場合の実装を行う

実行例:

```
% go run main.go https://upload.wikimedia.org/wikipedia/en/a/a9/Example.jpg
2018/07/04 17:14:28 Downloading https://upload.wikimedia.org/wikipedia/en/a/a9/Example.jpg to Example.jpg
2018/07/04 17:14:29 Total 27661 bytes
2018/07/04 17:14:29 Get 20748-27660 bytes of content
2018/07/04 17:14:29 Get 6916-13831 bytes of content
2018/07/04 17:14:29 Get 0-6915 bytes of content
2018/07/04 17:14:29 Get 13832-20747 bytes of content
2018/07/04 17:14:30 Wrote 6916-13831 bytes of content
2018/07/04 17:14:30 Wrote 20748-27660 bytes of content
2018/07/04 17:14:30 Wrote 0-6915 bytes of content
2018/07/04 17:14:30 Wrote 13832-20747 bytes of content
2018/07/04 17:14:30 Wrote 27661 bytes
```

大きいファイルの場合:

```
% go run main.go https://storage.googleapis.com/kubernetes-release/release/v1.11.0/bin/darwin/amd64/kubectl
2018/07/04 19:11:22 Downloading https://storage.googleapis.com/kubernetes-release/release/v1.11.0/bin/darwin/amd64/kubectl to kubectl
2018/07/04 19:11:23 Total 54949920 bytes
2018/07/04 19:11:23 Get 41212440-54949919 bytes of content
2018/07/04 19:11:23 Get 0-13737479 bytes of content
2018/07/04 19:11:23 Get 13737480-27474959 bytes of content
2018/07/04 19:11:23 Get 27474960-41212439 bytes of content
2018/07/04 19:11:27 Wrote 27474960-41212439 bytes of content
2018/07/04 19:11:28 Wrote 13737480-27474959 bytes of content
2018/07/04 19:11:30 Wrote 0-13737479 bytes of content
2018/07/04 19:11:30 Wrote 41212440-54949919 bytes of content
2018/07/04 19:11:30 Wrote 54949920 bytes
```

途中でWifi接続を切断した場合:

```
% go run main.go https://storage.googleapis.com/kubernetes-release/release/v1.11.0/bin/darwin/amd64/kubectl
2018/07/04 19:13:22 Downloading https://storage.googleapis.com/kubernetes-release/release/v1.11.0/bin/darwin/amd64/kubectl to kubectl
2018/07/04 19:13:23 Total 54949920 bytes
2018/07/04 19:13:23 Get 41212440-54949919 bytes of content
2018/07/04 19:13:23 Get 13737480-27474959 bytes of content
2018/07/04 19:13:23 Get 0-13737479 bytes of content
2018/07/04 19:13:23 Get 27474960-41212439 bytes of content
2018/07/04 19:13:30 Wrote 27474960-41212439 bytes of content
2018/07/04 19:13:30 Wrote 13737480-27474959 bytes of content
2018/07/04 19:13:49 Could not download https://storage.googleapis.com/kubernetes-release/release/v1.11.0/bin/darwin/amd64/kubectl: Could not write partial content: read tcp 172.16.3.103:59942->172.217.161.208:443: read: network is down
exit status 1
```
