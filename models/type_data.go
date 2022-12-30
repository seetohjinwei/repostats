package models

import (
	"database/sql/driver"
	"fmt"
)

type TypeData struct {
	Type      string `db:"language" json:"language"`
	FileCount int64  `db:"file_count" json:"file_count"`
	Bytes     int64  `db:"bytes" json:"bytes"`
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
	if x.FileCount != y.FileCount {
		return x.FileCount < y.FileCount
	} else if x.Bytes != y.Bytes {
		return x.Bytes < y.Bytes
	}
	return x.Type < y.Type
}

// Returns true if x > y.
func MoreTypeData(x, y TypeData) bool {
	return LessTypeData(y, x)
}

func (td TypeData) Value() (driver.Value, error) {
	s := fmt.Sprintf("(%s,%d,%d)",
		td.Type,
		td.FileCount,
		td.Bytes,
	)

	return []byte(s), nil
}
