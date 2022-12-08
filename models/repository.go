package models

import (
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Repository struct {
	Name          string
	DefaultBranch string
	Files         []File // contains files in sub-directories too
	FileTypes     map[string]TypeData
}

func (r Repository) ListFileTypes() string {
	if len(r.FileTypes) == 0 {
		return NO_FILES_FOUND
	}

	var sb strings.Builder

	fileTypes := maps.Values(r.FileTypes)
	slices.SortFunc(fileTypes, MoreTypeData)

	for _, fileType := range fileTypes {
		sb.WriteString(fileType.ToFormatted() + "\n")
	}

	return sb.String()
}
