package models

import "fmt"

type TypeData struct {
	Type      string
	FileCount uint
	Bytes     int64
}

const (
	FORMAT_TYPE_DATA = "[%s]: %d files (%d bytes)"
)

// TODO: better byte printing, B, kB, MB, GB, TB ...
func (td TypeData) ToFormatted() string {
	return fmt.Sprintf(FORMAT_TYPE_DATA, td.Type, td.FileCount, td.Bytes)
}

// Returns true if x < y.
// Useful for sort.Slice or other similar functions.
//
// Specifically, sort by bytes, file count, type (alphabetical).
func LessTypeData(x, y TypeData) bool {
	if x.Bytes != y.Bytes {
		return x.Bytes < y.Bytes
	} else if x.FileCount != y.FileCount {
		return x.FileCount < y.FileCount
	}
	return x.Type < y.Type
}

// Returns true if x > y.
func MoreTypeData(x, y TypeData) bool {
	return LessTypeData(y, x)
}
