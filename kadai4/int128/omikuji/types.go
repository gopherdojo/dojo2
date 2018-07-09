package omikuji

// Omikuji represents a result of omikuji.
type Omikuji int

const (
	_ Omikuji = iota
	// Kyo means 凶.
	Kyo
	// ShoKichi means 小吉.
	ShoKichi
	// ChuKichi means 中吉.
	ChuKichi
	// DaiKichi means 大吉.
	DaiKichi
)

func (o Omikuji) String() string {
	switch o {
	case Kyo:
		return "凶"
	case ShoKichi:
		return "小吉"
	case ChuKichi:
		return "中吉"
	case DaiKichi:
		return "大吉"
	default:
		return ""
	}
}
