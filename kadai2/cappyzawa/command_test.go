package conv_test

import (
	"image"
	"io"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo2/kadai2/cappyzawa"
	"errors"
)

func TestCommand_Run(t *testing.T) {
	t.Run("jpegToPng", func(t *testing.T) {
		expect := []string{
			"testdata/jpeg/gopher.png",
			"testdata/jpeg/kadai1.png",
			"testdata/jpeg/test/gopher.png",
			"testdata/jpeg/test/kadai1.png",
		}
		decoder := &MockDecoder{}
		encoder := &MockEncoder{}
		command := conv.NewCommand(decoder, encoder)
		actual, err := command.Run("testdata/jpeg", "jpeg", "png")
		if err != nil {
			t.Error(err.Error())
		}
		if len(actual[0]) == 0 {
			t.Error("actual's size is 0")
		}
		if !reflect.DeepEqual(expect, actual) {
			t.Error("expect is different from actual")
		}
	})
	t.Run("pngToJpeg", func(t *testing.T) {
		expect := []string{
			"testdata/png/gopher.jpeg",
			"testdata/png/kadai1.jpeg",
			"testdata/png/test/gopher.jpeg",
			"testdata/png/test/kadai1.jpeg",
		}
		decoder := &MockDecoder{}
		encoder := &MockEncoder{}
		command := conv.NewCommand(decoder, encoder)
		actual, err := command.Run("testdata/png", "png", "jpeg")
		if err != nil {
			t.Error(err.Error())
		}
		if len(actual[0]) == 0 {
			t.Error("actual's size is 0")
		}
		if !reflect.DeepEqual(expect, actual) {
			t.Error("expect is different from actual")
		}
	})
	t.Run("directory not found", func(t *testing.T) {
		decoder := &MockDecoder{}
		encoder := &MockEncoder{}
		command := conv.NewCommand(decoder, encoder)
		actual, err := command.Run("invalidDirectory", "png", "jpeg")
		if err == nil {
			t.Error("err is supposed to have occurred")
		}
		if actual != nil {
			t.Error("actual should be nil")
		}
	})
	t.Run("target file does not exist", func(t *testing.T) {
		decoder := &MockDecoder{}
		encoder := &MockEncoder{}
		command := conv.NewCommand(decoder, encoder)
		actual, err := command.Run("testdata/empty", "png", "jpeg")
		if err != nil {
			t.Error(err.Error())
		}
		if len(actual) != 0 {
			t.Error("actrual's size should be 0")
		}
	})
	t.Run("failed to decode", func(t *testing.T) {
		decoder := &MockErrDecoder{}
		encoder := &MockEncoder{}
		command := conv.NewCommand(decoder, encoder)
		actual, err := command.Run("testdata/png", "png", "jpeg")
		errMassage := "failed to decode"
		if err.Error() != errMassage {
			t.Errorf("err.Error should be %s", errMassage)
		}
		if actual != nil {
			t.Error("actual should be nil")
		}
	})
	t.Run("failed to encode", func(t *testing.T) {
		decoder := &MockDecoder{}
		encoder := &MockErrEncoder{}
		command := conv.NewCommand(decoder, encoder)
		actual, err := command.Run("testdata/png", "png", "jpeg")
		errMassage := "failed to encode"
		if err.Error() != errMassage {
			t.Errorf("err.Error should be %s", errMassage)
		}
		if actual != nil {
			t.Error("actual should be nil")
		}
	})
}

type MockDecoder struct {
}

type MockEncoder struct {
}

type MockErrDecoder struct {
}

type MockErrEncoder struct {
}

func (d *MockDecoder) Decode(r io.Reader) (image.Image, string, error) {
	return *new(image.Image), "", nil
}

func (e *MockEncoder) Encode(w io.Writer, m image.Image) error {
	return nil
}

func (d *MockErrDecoder) Decode(r io.Reader) (image.Image, string, error) {
	return nil, "", errors.New("failed to decode")
}

func (e *MockErrEncoder) Encode(w io.Writer, m image.Image) error {
	return errors.New("failed to encode")
}
