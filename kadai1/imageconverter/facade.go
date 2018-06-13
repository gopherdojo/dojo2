package imageconverter

// Facade mainからの命令窓口
type Facade struct{}

// Run Searcherを使ってファイル群を走査、Converterを使ってファイル群を変換処理にかける
func (f *Facade) Run(
	targetPath FilePath,
	inputFormat Format,
	outputFormat Format) {

	var searcher Searcher
	var converter Converter

	fileInfoList := searcher.Run(FileInfo{Path: targetPath})
	for _, fileInfo := range fileInfoList {
		converter.Run(fileInfo, inputFormat, outputFormat)
	}
}
