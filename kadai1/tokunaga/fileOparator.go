package tokunaga

import (
	"path"
	)

func FullBasename(fullFIlePath string) string {
	return fullFIlePath[0 : len(fullFIlePath)-len(path.Ext(fullFIlePath))]
}