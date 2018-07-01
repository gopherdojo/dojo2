# kadai2
## io.Readerとio.Writer
### 標準パッケージでどのように使われているか
メソッドの引数で用いられている。  
今回の課題で使用した`image/jpeg`パッケージの[`Decode`](https://github.com/golang/go/blob/f03ee913e210e1b09bd33ed35c03ec8e4fc270be/src/image/jpeg/writer.go#L575)でも使用されている。  

### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
`io.Reader`が引数として指定されている場合は、`Read(p []byte) (n int, err error)`というシグネチャ、`io.Writer`が指定されている場合は、`Write(p []byte) (n int, err error)`のシグネチャにmatchさえしていればどんなものでも引数として与えることができる。  
上記の`Decode()`の引数には`*os.File`型の変数を渡しているが、`*os.File`型は`Read(p []byte) (n int, err error)`にmatchするメソッドをもっている。  
同様に`Encode()`の引数にも`*os.File`型の変数を渡すことができ、`*os.File`型は`Write(p []byte) (n int, err error)`にmatchするメソッドをもっている。

今回の場合はファイルの読み書きをしたかったため、ファイルを引数として渡したが、標準入力/出力の読み書きを行い場合はそれらを引数とすることもできる。
つまり、Interfaceとして`io.Reader`、`io.Writer`があることで、抽象的に読み書きが可能になるという利点がある。

## kadai1との差分
* foo
* bar
