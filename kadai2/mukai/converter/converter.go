// image converter.
package converter

import (
	"fmt"
	"os"
	"strings"
)

// convert image files from inputFormat to outputFormat in specific directory recursively.
// inputFormat and outputFormat is available only jpeg(jpg), gif, or png.
func RecursiveConvert(dir string, inputFormat string, outputFormat string, pather Pather) ([]string, error) {
	if !isAvailableFormat(inputFormat) || !isAvailableFormat(outputFormat) {
		return nil, fmt.Errorf(
			"available formats are jpg(jpeg), png, gif ONLY. input parameter is %s, outpur parameter is %s",
			inputFormat, outputFormat)
	}
	infos, err := pather.files(dir)
	if err != nil {
		return nil, err
	}
	var convertedFiles []string
	for _, file := range infos {
		if file.IsDir() {
			c, e := RecursiveConvert(file.AbsPath(), inputFormat, outputFormat, pather)
			if e != nil {
				fmt.Fprintln(os.Stderr, "failed to convert " + file.AbsPath())
				continue
			}
			convertedFiles = append(convertedFiles, c...)
		} else if IsSameExt(file.AbsPath(), inputFormat) {
			err := file.Convert(outputFormat)
			if err != nil {
				fmt.Fprintln(os.Stderr, "failed to convert " + file.AbsPath())
			}
			convertedFiles = append(convertedFiles, file.AbsPath())
		}
	}
	return convertedFiles, nil
}

func isAvailableFormat(format string) bool {
	lowerFormat := strings.ToLower(format)
	switch lowerFormat {
	case "jpg", "jpeg", "gif", "png":
		return true
	default:
		return false
	}
}


