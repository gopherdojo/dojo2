### 動作方法
- ```git clone https://github.com/chillout2san/dojo2.git```
- ```cd ./dojo2/kadai1/chillout2san```
- ```go build main.go```
- ```./main -before {変換前拡張子} -after {変換後の拡張子} -path {ディレクトリ}```
- ```./main -before jpg -after png -path ./```

### 備考
- beforeとafterは省略可能。省略した場合、beforeはjpg、afterはpngになる。
- pathは省略不可。省略した場合、エラーを投げる。
- 現在のファイルの拡張子とbeforeの拡張子が異なる場合、エラーを投げる。