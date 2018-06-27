package imageconverter

// Facade mainからの命令窓口
type Facade struct {
	Searcher  SearcherInterface
	Converter ConverterInterface
}

// Run Searcherを使ってファイル群を走査、Converterを使ってファイル群を変換処理にかける
func (f *Facade) Run(targetPath FilePath, in, out Format) {
	fileInfoList := f.Searcher.Run(FileInfo{Path: targetPath})
	for _, fileInfo := range fileInfoList {
		f.Converter.Run(fileInfo, in, out)
	}
}
