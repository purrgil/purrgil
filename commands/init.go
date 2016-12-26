package commands

import (
	"os"

	"github.com/guidiego/purrgil/file"
	"github.com/guidiego/purrgil/interactiveshell"
)

func Init(projectName string) {
	wd, _ := os.Getwd()
	wdpath := wd + "/" + projectName

	ishell.PurrgilAlert("Init a new Purrgil App into: " + projectName)

	os.Mkdir(projectName, 0777)

	dockercompose := file.NewDockerCompose(wdpath)
	purrgil := file.NewPurrgil(wdpath, projectName)
	gitignore := file.NewGitIgnore(wdpath)

	dockercompose.SaveFile()
	purrgil.SaveFile()
	gitignore.SaveFile()

	ishell.PurrgilAlert("Project created! To enter in project give 'cd " + projectName + "'")
}
