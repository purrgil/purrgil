package file

import (
	"strings"
)

type GitIgnore struct {
	File
	ignoredList []string
}

func (g *GitIgnore) LoadFile() {
	g.File.LoadFile()

	g.ignoredList = strings.Split(string(g.File.content), "\n")
}

func (g *GitIgnore) SaveFile() {
	compiledString := strings.Join(g.ignoredList, "\n")

	g.File.content = []byte(compiledString)
	g.File.SaveFile()
}

func (g *GitIgnore) AddIgnoredPath(path string) {
	g.ignoredList = append(g.ignoredList, path)
}

func (g *GitIgnore) RemoveFromIgnored(pkgName string) {
	filteredIgnored := []string{}

	for _, val := range g.ignoredList {
		if val != pkgName {
			filteredIgnored = append(filteredIgnored, val)
		}
	}

	g.ignoredList = filteredIgnored
}

func NewGitIgnore(dir string) GitIgnore {
	gitignore := GitIgnore{}

	gitignore.File.InitFile(dir, ".gitignore", "")
	gitignore.LoadFile()

	return gitignore
}
