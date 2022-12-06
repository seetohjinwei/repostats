package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Directory struct {
	Path  string
	Name  string
	Dirs  []Directory
	Files []File
}

const (
	FORMAT_NAME = "Directory [%s]:\n"
	FORMAT_SUB  = "[%d] - %s\n"
	FORMAT_FILE = "%s\n"
)

func (d Directory) ListOptions() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(FORMAT_NAME, d.Name))

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
