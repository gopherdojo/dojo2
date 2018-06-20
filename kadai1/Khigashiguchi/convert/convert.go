package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo2/kadai1/Khigashiguchi/ext"
)

// Converter Converter image file
type Converter struct {
	Dir    string
	InExt  string
	OutExt string
}

// Exec exec to convert image file
func (c *Converter) Exec() {
	err := filepath.Walk(c.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == c.InExt {
			sf, err := os.Open(path)
			if err != nil {
				return err
			}
			defer sf.Close()

			var img image.Image
			switch c.InExt {
			case ext.Format("jpg"):
				img, err = jpeg.Decode(sf)
			case ext.Format("png"):
				img, err = png.Decode(sf)
			}
			if err != nil {
				return err
			}
			fp := "out/" + strings.Replace(filepath.Base(path), c.InExt, c.OutExt, -1)
			ds, err := os.Create(fp)
			if err != nil {
				return err
			}
			defer ds.Close()
			switch c.OutExt {
			case ext.Format("jpg"):
				err = jpeg.Encode(ds, img, nil)
			case ext.Format("png"):
				err = png.Encode(ds, img)
			}
			fmt.Printf("Complete convert file %s Output is %s\n", path, fp)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", c.Dir, err)
	}
}
