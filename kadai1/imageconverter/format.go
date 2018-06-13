package imageconverter

// Format ファイルのフォーマットを表す型
type Format string

// Ext 拡張子を取得する関数
func (f *Format) Ext() string {
	return `.` + string(*f)
}

// NormalizedFormat 本ツール内で扱いやすいフォーマットに変換する。(jpeg => jpg をやりたいがための実装)
func (f *Format) NormalizedFormat() Format {
	switch string(*f) {
	case "jpeg":
		return Format("jpg")
	default:
		return *f
	}
}
