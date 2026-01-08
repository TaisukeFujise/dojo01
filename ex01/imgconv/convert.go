package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

func OpenInput(path string) (io.ReadCloser, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func CreateOutput(path string, output Format) (io.WriteCloser, error) {
	w, err := os.Create(pathWithoutFormat(path) + string(output))
	if err != nil {
		return nil, err
	}
	return w, nil
}

func pathWithoutFormat(path string) string {
	return path[:len(path)-len(filepath.Ext(path))]
}

// Convert converts r to w by following specified format.
func Convert(r io.Reader, w io.Writer, output Format) (err error) {
	switch output {
	case ".jpeg":
		err = convertToJPEG(r, w)
	case ".png":
		err = convertToPNG(r, w)
	case ".gif":
		err = convertToGIF(r, w)
	default:
		err = fmt.Errorf("invalid format")
	}
	if err != nil {
		return err
	}
	return nil
}

func convertToJPEG(r io.Reader, w io.Writer) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, nil)
}

func convertToPNG(r io.Reader, w io.Writer) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return png.Encode(w, img)
}

func convertToGIF(r io.Reader, w io.Writer) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return gif.Encode(w, img, nil)
}
