package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Directory struct {
	Path               string
	Name               string
	Dirs               []Directory
	Files              []File
	FileTypes          map[string]TypeData
	RecursiveFileTypes map[string]TypeData
}

func NewDirectory(path, name string) Directory {
	return Directory{
		Path:      path,
		Name:      name,
		Dirs:      []Directory{},
		Files:     []File{},
		FileTypes: nil,
	}
}

const (
	LISTING_TYPES = "--- Types ---\n"

	FORMAT_NAME = "--- Directory [%s] ---\n"
	FORMAT_SUB  = "[%d] - %s\n"
	FORMAT_FILE = "%s\n"
)

func (d Directory) ListEverything() string {
	return d.ListTitle() + d.ListFileTypes() + "\n" + d.ListOptions()
}

func (d Directory) ListTitle() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(FORMAT_NAME, d.Name))

	return sb.String()
}

// TODO: Message for no options
func (d Directory) ListOptions() string {
	var sb strings.Builder

	for i, sub := range d.Dirs {
		sb.WriteString(fmt.Sprintf(FORMAT_SUB, i, sub.Name))
	}

	return sb.String()
}

func (d Directory) SubDir(index int) (*Directory, error) {
	if index < 0 || index >= len(d.Dirs) {
		return nil, errors.New("invalid index")
	}

	return &d.Dirs[index], nil
}

func (d Directory) SubDirString(index string) (*Directory, error) {
	ind, err := strconv.Atoi(index)
	if err != nil {
		return nil, err
	}
	return d.SubDir(ind)
}

func (d *Directory) loadFileTypes() bool {
	if d.FileTypes == nil {
		d.FileTypes = map[string]TypeData{}

		for _, file := range d.Files {
			_, ok := d.FileTypes[file.TypeData.Type]
			if !ok {
				d.FileTypes[file.TypeData.Type] = TypeData{
					Type:      file.TypeData.Type,
					FileCount: 0,
					Bytes:     0,
				}
			}

			entry := d.FileTypes[file.TypeData.Type]
			entry.FileCount += file.TypeData.FileCount
			entry.Bytes += file.TypeData.Bytes
			d.FileTypes[file.TypeData.Type] = entry
		}

		return true
	}

	return false
}

// TODO: Message for no files
// TODO: Sort by size, count, name (alphabetical)
func (d Directory) ListFileTypes() string {
	d.loadFileTypes()

	var sb strings.Builder

	sb.WriteString(LISTING_TYPES)

	for _, typeData := range d.FileTypes {
		sb.WriteString(typeData.ToFormatted() + "\n")
	}

	return sb.String()
}
