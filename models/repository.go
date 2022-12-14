package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

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

type PSQLRepository struct {
	Username      string    `db:"username" json:"username"`
	Repo          string    `db:"repo" json:"repo"`
	LastUpdated   time.Time `db:"last_updated" json:"last_updated"`
	DefaultBranch string    `db:"default_branch" json:"default_branch"`
}

func (pr PSQLRepository) Value() (driver.Value, error) {
	s := fmt.Sprintf("(%s,%s,%s,%s)",
		pr.Username,
		pr.Repo,
		pr.LastUpdated.Format(time.RFC3339),
		pr.DefaultBranch,
	)

	return []byte(s), nil
}
