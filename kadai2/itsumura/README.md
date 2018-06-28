# テスト

実行コマンド
```
go test -run ""
```

テストカバレッジ表示
```
go test -cover ""
```

# io.Readerとio.Writer
## 標準パッケージでの使われ方
- 標準パッケージでの関数で、引数として使われていた
- 引数として渡された値は、Readerではbufio.NewReader()やbufio.NewScanner()で読み込まれていた

## メリット
- 成功したか否かを取得し、エラーハンドリングを用意にするため？
