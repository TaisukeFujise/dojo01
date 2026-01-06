// Package imgconv provides utilities for converting image formats.
package imgconv

import (
	"fmt"
	"path/filepath"
	"strings"
)

// A Format value is a standardized and validated image format.
// A Format value always includes a leading dot.
type Format string

// ParseFormat parses a raw string and returns a Format and error.
// The returned Format is the file extension including a leading dot.
// The returned error is nil if the raw value is one of the supported formats.
func ParseFormat(raw string) (Format, error) {
	switch strings.ToLower(raw) {
	case "jpg", "jpeg":
		return Format(".jpg"), nil
	case "png":
		return Format(".png"), nil
	case "gif":
		return Format(".gif"), nil
	default:
		return Format(""), fmt.Errorf("%s is not a valid format", raw)
	}
}

// Validate reports whether the path is the one of the supported formats.
func (f Format) Validate(path string) bool {
	ext := filepath.Ext(path)
	if ext == "" {
		return false
	}
	_, err := ParseFormat(ext[1:])
	if err != nil {
		return false
	}
	return true
}

// Match reports whether the path matches f.
// It returns false if the path does not have an extension.
func (f Format) Match(path string) bool {
	ext := filepath.Ext(path)
	if ext == "" {
		return false
	}
	pathFormat, err := ParseFormat(ext[1:])
	if err != nil {
		return false
	}
	return pathFormat == f
}
