# build
```
go build -o imageconv cmd/imageConverter/main.go
```

# Usage
```
./imageconv [-from=jpeg|jpg|png] [-to=jpeg|jpg|png] directory
```

|オプション|詳細        |必須 |備考|
|--------|------------|-----|--|
|from    |変換元の拡張子|  -  |jpeg, jpg, pngのどれか(デフォルトはjpeg)|
|to      |変換後の拡張子|  -  |jpeg, jpg, pngのどれか(デフォルトはpng)|


# 課題説明
`path/filepath` の `Walk` を使用し、画像変換用の関数を再帰的に実行。  
変換前と変換後の拡張子を変更しても画像変換用の関数をそのまま使用できるように画像変換を行うための  
ineterfaceを作成。
