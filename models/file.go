package models

import (
	"path/filepath"
)

type File struct {
	Path      string
	Name      string
	Extension string
	TypeData  TypeData
}

func NewFile(path, name string, bytes int64) File {
	extension := filepath.Ext(name)
	if len(extension) > 0 {
		// trim the "."
		extension = extension[1:]
	}
	typeData := TypeData{
		Type:      GetFileType(extension),
		FileCount: 1,
		Bytes:     bytes,
	}

	return File{
		path,
		name,
		extension,
		typeData,
	}
}
