# 課題1: 画像変換コマンドの実装

## build 

`go build -o convert-cli cli.go`

## 使い方

`./convert-cli dir_path`
dir_path 内のjpg画像がpngに変換され、dir_path内にpngファイルが追加で作成される

`./convert-cli -s png -d jpeg dir_path`
dir_path 内のsオプションで指定したフォーマットの画像がdオプションで指定されたフォーマットに変換され、dir_path内に変換後のファイルが追加で作成される

## 補足
対応している画像の形式は jpeg, jpg, png のみ。

