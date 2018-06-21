# imgconverter
.png/.jpg/.gifの相互変換を行うコマンドツールです。

## 使い方

    imgconverter [-src <source-format>] [-dst <destination-format>] <path-to-dir>

<path-to-dir>配下にある指定フォーマットの画像をすべて指定したフォーマットに変換します。
<path-to-dir>配下にディレクトリがある場合は、そのディレクトリを再帰的に検索して、画像を変換します。

フラグ:

    -src 画像の変換元のフォーマットを指定します。例:jpg デフォルトはjpg
    -dst 画像の変換後のフォーマットを指定します。例:png デフォルトはpng

## 使用例

~~~
% imgconverter -src jpg -dst png /testImg/
~~~

## 工夫したところ

* 前回と同様にテストを行いやすくするためCLIインターフェースを定義して、そこから処理を呼び出すようにした。
* Godocで表示した時に見やすいコメントのインデントの仕方（Godocちょうべんり）

## 疑問・質問

* main.goとimgconverter.goでほぼほぼ処理をimgconverter.goに書いてしまった。main.goとライブラリで処理を分ける基準があれば教えてほしいです。
* メインの処理をimgconverter.goで書いてmain.goで呼び出す時に、GOPATHのsrc以下のディレクトリ構造によってimportの書き方が変わってしまう。
  この際のなにか行う工夫などがあれば教えてほしいです。（自宅PCと仕事PCでGOPATH/src以下の構造が異なっており困りましたｗ）