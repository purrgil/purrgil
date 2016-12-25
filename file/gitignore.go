package file

type GitIgnore struct {
  File
  ignored []string
}

func (g *GitIgnore) LoadFile() {
    g.LoadFile()
    // Transform g.content in a []string into g.ignored
}

func (g *GitIgnore) UpdateContent() []byte {
  // Transform g.ignores in a []byte
  return []byte{}
}

func NewGitIgnore(dir string) GitIgnore {
  gitignore := GitIgnore{}

  gitignore.InitFile(dir, ".gitignore", "")
  gitignore.LoadFile()

  return gitignore
}
