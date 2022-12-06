package models

import (
	"strings"
)

// TODO: put this in JSON / TOML file
var types map[string]string = map[string]string{
	"audio":            "aif cda mid midi mp3 mpa ogg wav wma wpl",
	"data":             "csv dat db dbf log mdb sav tar xml",
	"font":             "fnt fon",
	"image":            "gif ico jpeg jpg png svg",
	"astro":            "astro",
	"c":                "c",
	"c++":              "cpp cc",
	"css":              "css sass scss",
	"c/c++ header":     "h",
	"go":               "go",
	"html":             "html htm",
	"java":             "java",
	"javascript react": "jsx",
	"javascript":       "js",
	"python":           "py",
	"sql":              "sql",
	"ruby":             "rb",
	"lua":              "lua",
	"typescript react": "tsx",
	"typescript":       "ts",
}

var fileTypeMapper map[string]string = nil

const unknown string = "unknown"

func loadFileTypeMapper() bool {
	// instantiate singleton
	if fileTypeMapper == nil {
		fileTypeMapper = map[string]string{}

		for fileType, s := range types {
			for _, ext := range strings.Split(s, " ") {
				fileTypeMapper[fileType] = ext
			}
		}

		return true
	}

	return false
}

func GetFileType(extension string) string {
	loadFileTypeMapper()

	val, ok := fileTypeMapper[extension]
	if !ok {
		return unknown
	}

	return val
}
