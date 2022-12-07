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
