package imgconv

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

// Convert converts a file specified by path to the specified format.
func Convert(path string, output Format) error {
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer reader.Close()
	writer, err := os.Create(pathWithoutFormat(path) + string(output))
	if err != nil {
		return err
	}
	defer writer.Close()
	switch output {
	case ".jpg":
		err = convertToJPEG(reader, writer)
	case ".png":
		err = convertToPNG(reader, writer)
	case ".gif":
		err = convertToGIF(reader, writer)
	}
	if err != nil {
		return err
	}
	return nil
}

func pathWithoutFormat(path string) string {
	return path[:len(path)-len(filepath.Ext(path))]
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
