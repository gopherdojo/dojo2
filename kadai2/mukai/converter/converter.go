// image converter.
package converter

import (
	"fmt"
	"os"
	"strings"
)

// convert image files from inputFormat to outputFormat in specific directory recursively.
// inputFormat and outputFormat is available only jpeg(jpg), gif, or png.
func RecursiveConvert(dir string, inputFormat string, outputFormat string, pather Pather) error {
	if !isAvailableFormat(inputFormat) || !isAvailableFormat(outputFormat) {
		return fmt.Errorf(
			"available formats are jpg(jpeg), png, gif ONLY. input parameter is %s, outpur parameter is %s",
			inputFormat, outputFormat)
	}
	infos, err := pather.files(dir)
	if err != nil {
		return err
	}
	for _, file := range infos {
		//inputConvertFile := convertFile{absPath: filepath.Join(dir, file.Name()), isDir: file.IsDir()}
		if file.isDir {
			RecursiveConvert(file.absPath, inputFormat, outputFormat, pather)
		} else if file.isSameExt(inputFormat) {
			err := file.Convert(outputFormat)
			if err != nil {
				fmt.Fprintln(os.Stderr, "failed to convert " + file.absPath)
			}
		}
	}
	return nil
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


