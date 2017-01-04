package commands

import (
	"os"

	"github.com/purrgil/purrgil/file"
	"github.com/purrgil/purrgil/interactiveshell"
)

func Init(projectName string, projectGithub string) {
	wd, _ := os.Getwd()
	wdpath := wd + "/" + projectName

	ishell.PurrgilAlert("Init a new Purrgil App into: " + projectName)

	if (projectGithub == "") {
		os.Mkdir(projectName, 0777)
	} else {
		GitClone(projectGithub, projectName)
	}

	dockercompose := file.NewDockerCompose(wdpath)
	purrgil := file.NewPurrgil(wdpath, projectName)
	gitignore := file.NewGitIgnore(wdpath)

	dockercompose.SaveFile()
	purrgil.SaveFile()
	gitignore.SaveFile()

	ishell.PurrgilAlert("Project created! To enter in project give 'cd " + projectName + "'")
}
