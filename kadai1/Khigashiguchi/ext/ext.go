package ext

// Format format extentison add period
func Format(f string) string {
	return "." + f
}

// Validate formated extentision
func Validate(f string) bool {
	if f == Format("jpg") || f == Format("png") {
		return true
	}
	return false
}
