課題2

## io.Readerとio.Writer
---
io.Readerとio.Writerについて調べてみよう
標準パッケージでどのように使われているか
io.Readerとio.Writerがあることで
どういう利点があるのか具体例を挙げて考えてみる


## テストを書いてみよう
---
- 1回目の宿題のテストを作ってみて下さい
- テストのしやすさを考えてリファクタリングしてみる
- テストのカバレッジを取ってみる
- テーブル駆動テストを行う
- テストヘルパーを作ってみる


### テストヘルパー
https://github.com/gopherdojo/dojo2/commit/0bf8f7752080729a7b43ebd3ede1e177286c9695  
-> interfaceをDecode, Encode単位に分割
### テーブル駆動テスト
https://github.com/gopherdojo/dojo2/commit/b73b1d42aeff663637bed153a0b44178c9523f68

### リファクタリング
main処理からflagのパース,画像変換コマンド切り出し  
https://github.com/gopherdojo/dojo2/commit/70d7d8b72f0d5f69b1667fc5f99a1f2f0d5d561e  

参考:  
https://github.com/hashicorp/atlas-upload-cli  
https://deeeet.com/writing/2014/12/18/golang-cli-test/  
