package models

import "fmt"

type TypeData struct {
	Type      string
	FileCount uint
	Bytes     int64
}

const (
	FORMAT_BYTES     = "%d%s"
	FORMAT_TYPE_DATA = "[%s]: %d %s (%s)"
)

var units []string = []string{"B", "kB", "MB", "GB", "TB"}

func (td TypeData) toFormattedBytes() string {
	index := 0
	bytes := td.Bytes
	for bytes >= 1000 {
		bytes /= 1000
		index++
	}

	return fmt.Sprintf(FORMAT_BYTES, bytes, units[index])
}

func (td TypeData) ToFormatted() string {
	files := "files"
	if td.FileCount == 1 {
		files = "file"
	}
	return fmt.Sprintf(FORMAT_TYPE_DATA, td.Type, td.FileCount, files, td.toFormattedBytes())
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
