
# io.Readerとio.Writerについて

## 標準パッケージでどのように使われているか

抽象化されたReadとWrite処理。
各パッケージのデータ型で実際のReadの処理とWriteの処理が実装されている。

また、パッケージによってはReader,WriterのInterfaceを拡張して使用している。

## `io.Reader` と `io.Writer` があることでどういう利点があるのか


