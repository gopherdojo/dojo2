# 課題2-1【TRY】io.Readerとio.Writer
## 概要
- io.Readerとio.Writerについて調べてみよう
    - 標準パッケージでどのように使われているか
    - io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
## 内容
### 標準パッケージでどのように使われているか
- tcp/udp層の読み書き・ファイルの読み書き・標準入出力など、「読み・書き」に関しては使われている
- 具体例
    - bufio.NewReader
        - https://golang.org/src/bufio/bufio.go?s=1647:1683#L51
        - 引数にio.Readerを渡している
    - http.Request.Write
        - https://golang.org/src/net/http/request.go?s=16549:16591#L467
        - 引数にio.Writeを渡している

### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
- ログ出力をテストしたいといったときにio.Writerに対して実装すれば、出力だけ他の実装に差し替えることでテスト可能になる
- 最初はファイル出力・途中からflentdに直接投げようといった差し替えが用意になる

## 雑記メモ
### そもそもio.Reader / io.Writerは
- io.Reader
    - Reader is the interface that wraps the basic Read method.
        - https://golang.org/pkg/io/#Reader
    - 講義内資料
        - https://docs.google.com/presentation/d/1Ri0sN-jnQTDEhV0Oiet4T_voSg7Sr9qnRDSOIWQsdW8/edit#slide=id.g384cb05baf_0_236s