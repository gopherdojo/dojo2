# 仕様
- コマンド実行時にディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGファイルに変換する
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる

- https://docs.google.com/presentation/d/1Ri0sN-jnQTDEhV0Oiet4T_voSg7Sr9qnRDSOIWQsdW8/edit#slide=id.g365eb9548c_1_1110

# 制約
- mainパッケージから画像変換ロジックを分離する
- 自作・標準・準標準パッケージのみを使用
- ユーザー定義型を利用する
- GoDocを生成する

# 使い方

## 実行する
- オプション無し

```bash
-> % go run command.go samples
Complete convert file samples/SampleJPGImage_100kbmb.jpg Output is out/SampleJPGImage_100kbmb.png
Complete convert file samples/SampleJPGImage_200kbmb.jpg Output is out/SampleJPGImage_200kbmb.png
Complete convert file samples/SampleJPGImage_50kbmb.jpg Output is out/SampleJPGImage_50kbmb.png
```

- 拡張子オプション有り

```bash
-> % go run command.go samples --in png --out jpg
Complete convert file samples/SampleJPGImage_100kbmb.jpg Output is out/SampleJPGImage_100kbmb.png
Complete convert file samples/SampleJPGImage_200kbmb.jpg Output is out/SampleJPGImage_200kbmb.png
Complete convert file samples/SampleJPGImage_50kbmb.jpg Output is out/SampleJPGImage_50kbmb.png
```

## godocを更新する

```bash
-> % make build-doc
```

# 実装中のTodoリスト
- [x] コマンドI/F
- [x] コマンドflagのparse
- [x] ディレクトリオープンして中のファイルを取得
- [x] 中のファイルをopenする
- [x] pngをjpegに変換する
- [x] ユーザー定義型を作る
- [x] convertパッケージを作る
- [x] GoDocを生成する
- [x] 拡張子の種類をバリデーションする
- [x] pngからjpg / jpgからpngの切り替え