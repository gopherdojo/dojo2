# 実行方法
```
go build -o imgconv main.go
./imgconv -i jpg -o png images
```
or 
```
go run main -i jpg -o png images
```

## options
-i 変換前画像形式

-o 変換後画像形式


## 説明
converter/Convertによって指定され他ディレクトリ以下を再帰的に指定された形式で画像変換する.

変換には"image/png","image/jpeg","image/gif"を使用する.

ユーザ定義型は以下の型を定義. 拡張子名操作の便利機能を持ちファイルパスを表現する.
```
type ConvertFile struct {
	absPath string
	isDir bool
}
```
