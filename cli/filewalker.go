package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/seetohjinwei/repostats/cli/messages"
	"github.com/seetohjinwei/repostats/models"
)

const (
	maxRecursion = 100
)

var (
	// TODO: read .gitignore, if exists
	ignoreList = []string{".", "..", ".git", "node_modules"}
)

func shouldParse(dir string) bool {
	for _, ignore := range ignoreList {
		if dir == ignore {
			return false
		}
	}
	return true
}

func parseDirectory(dir, path string, recurseLevel int) (models.Directory, error) {
	result := models.NewDirectory(path, dir)

	// Just in case.
	if recurseLevel >= maxRecursion {
		return result, errors.New("maximum recursion reached")
	}

	files, err := os.ReadDir(path)
	if err != nil {
		msg := fmt.Sprintf(messages.INVALID_DIRECTORY, path)
		return result, errors.New(msg)
	}

	for _, f := range files {
		name := f.Name()
		if !shouldParse(name) {
			continue
		}
		if f.IsDir() {
			path := path + "/" + name
			subdir, err := parseDirectory(name, path, recurseLevel+1)
			if err != nil {
				return result, err
			}
			result.Dirs = append(result.Dirs, subdir)
		} else {
			info, err := f.Info()
			if err != nil {
				return result, err
			}
			file := models.NewFile(path, name, info.Size())
			result.Files = append(result.Files, file)
		}
	}

	return result, nil
}
