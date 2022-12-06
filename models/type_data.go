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

func (td TypeData) ToFormatted() string {
	return fmt.Sprintf(FORMAT_TYPE_DATA, td.Type, td.FileCount, td.Bytes)
}
