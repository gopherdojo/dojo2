package imageconverter

// CommandValidator mainで取得したoptionの検査器
type CommandValidator struct{}

// ExtValidate 対象フォーマットならばtrue
func (CommandValidator) ExtValidate(f Format) bool {
	if f == Format("jpg") || f == Format("png") {
		return true
	}
	return false
}
