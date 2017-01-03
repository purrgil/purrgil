package commands

import (
	"github.com/guidiego/purrgil/file"
	"github.com/guidiego/purrgil/interactiveshell"
	"os"
)

func Remove(pkgName string) {
	ishell.PurrgilAlert("Removing " + pkgName + "...")

	path, _ := os.Getwd()
	purrgilconfig := file.NewPurrgil(path, "")
	dockercompose := file.NewDockerCompose(path)
	gitignore := file.NewGitIgnore(path)

	err := os.RemoveAll(path + "/" + pkgName)

	if err != nil {
		panic(err)
	}

	purrgilconfig.RemoveFromPackages(pkgName)
	gitignore.RemoveFromIgnored(pkgName)
	dockercompose.RemoveFromServices(pkgName)

	purrgilconfig.SaveFile()
	dockercompose.SaveFile()
	gitignore.SaveFile()

	ishell.PurrgilAlert(pkgName + " was successfuly removed!")
}
