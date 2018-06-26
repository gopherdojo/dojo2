package images

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

func ExampleConversion_ReplaceExt() {
	conversion := &Conversion{DestinationExt: "png"}
	destination := conversion.ReplaceExt("hello.jpg")
	fmt.Println(destination)
	// Output: hello.png
}

func ExampleConversion_Do() {
	// Create a JPEG image and plot a pixel
	jpegImage := image.NewRGBA(image.Rect(0, 0, 100, 200))
	jpegFile, err := ioutil.TempFile("", "jpeg")
	if err != nil {
		panic(err)
	}
	defer jpegFile.Close()
	defer os.Remove(jpegFile.Name())
	if err := jpeg.Encode(jpegFile, jpegImage, nil); err != nil {
		panic(err)
	}

	// Convert from JPEG to PNG
	conversion := &Conversion{
		Decoder: &JPEG{},
		Encoder: &PNG{},
	}
	source := jpegFile.Name()
	destination := conversion.ReplaceExt(source)
	if err := conversion.Do(source, destination); err != nil {
		panic(err)
	}

	// Read the PNG image
	pngFile, err := os.Open(destination)
	if err != nil {
		panic(err)
	}
	defer pngFile.Close()
	defer os.Remove(pngFile.Name())
	pngImage, err := png.Decode(pngFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("size=%+v", pngImage.Bounds().Size())
	// Output: size=(100,200)
}
