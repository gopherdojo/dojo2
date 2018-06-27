// image converter.
package converter

import (
	"fmt"
	"os"
	"strings"
)

// convert image files from inputFormat to outputFormat in specific directory recursively.
// inputFormat and outputFormat is available only jpeg(jpg), gif, or png.
func RecursiveConvert(dir string, inputFormat string, outputFormat string, pather pather) ([]string, error) {
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
		if file.isDirectory() {
			c, e := RecursiveConvert(file.absolutePath(), inputFormat, outputFormat, pather)
			if e != nil {
				fmt.Fprintln(os.Stderr, "failed to convert " + file.absolutePath())
				continue
			}
			convertedFiles = append(convertedFiles, c...)
		} else if isSameExt(file.absolutePath(), inputFormat) {
			outputFile, err := file.convert(outputFormat)
			if err != nil {
				fmt.Fprintln(os.Stderr, "failed to convert " + file.absolutePath())
			}
			convertedFiles = append(convertedFiles, outputFile)
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


