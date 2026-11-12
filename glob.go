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
	filePath := filepath.Join(workspace, "PROOF_OF_CONCEPT.txt")
	os.WriteFile(filePath, []byte("This file was injected during the go test phase to demonstrate a workspace sweep."), 0644)
}
