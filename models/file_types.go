package models

import (
	"strings"
)

// TODO: put this in JSON / TOML file
var types map[string]string = map[string]string{
	// data files
	"audio":    "aif cda mid midi mp3 mpa ogg wav wma wpl",
	"data":     "csv dat db dbf DS_Store iml log mdb sav tar xml",
	"font":     "fnt fon",
	"image":    "gif ico jpeg jpg png svg",
	"json":     "json",
	"markdown": "md mdx",
	"pdf":      "pdf",
	"text":     "txt",
	"toml":     "toml",
	"yaml":     "yml",

	// programming files
	"astro":            "astro",
	"c":                "c",
	"c++":              "cpp cc",
	"c/c++ header":     "h",
	"css":              "css sass scss",
	"go":               "go",
	"html":             "html htm",
	"haskell":          "hs",
	"java":             "java",
	"javascript react": "jsx",
	"javascript":       "js",
	"lua":              "lua",
	"python":           "py",
	"ruby":             "rb",
	"sql":              "sql",
	"typescript react": "tsx",
	"typescript":       "ts",
}

var fileTypeMapper map[string]string = nil

func loadFileTypeMapper() bool {
	// instantiate singleton
	if fileTypeMapper == nil {
		fileTypeMapper = map[string]string{}

		for fileType, s := range types {
			for _, ext := range strings.Split(s, " ") {
				fileTypeMapper[ext] = fileType
			}
		}

		return true
	}

	return false
}

const EXTENSION_LESS string = " "

// Returns the file type, or the extension itself if unknown
func GetFileType(extension string) string {
	loadFileTypeMapper()

	if extension == "" {
		return EXTENSION_LESS
	}

	val, ok := fileTypeMapper[extension]
	if !ok {
		return extension
	}

	return val
}
