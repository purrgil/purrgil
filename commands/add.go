package commands

import (
	"github.com/guidiego/purrgil/configs"
	"github.com/guidiego/purrgil/file"
	"github.com/guidiego/purrgil/interactiveshell"
	"os"
)

func Add(pkgId string, opts configs.AddConfig) {
	ishell.PurrgilAlert("Starting Purrgil Add on Package: " + pkgId)

	path, _ := os.Getwd()

	purrgilconfig := file.NewPurrgil(path, "")
	dockercompose := file.NewDockerCompose(path)
	gitignore := file.NewGitIgnore(path)
	purrgilNewPackage := file.NewPurrgilPackage(pkgId, opts)

	purrgilconfig.AddPackage(purrgilNewPackage)
	gitignore.AddIgnoredPath(purrgilNewPackage.Name)

	ishell.PurrgilAlert("Preparing to Collect Aditional Information...")

	if purrgilNewPackage.Service {
		serviceName, service := ishell.CollectDockerServiceInfo(purrgilNewPackage)
		dockercompose.AddService(serviceName, service)
	} else {
		packages := ishell.CollectLinkPossibility(purrgilNewPackage)
		dockercompose.LinkInService(purrgilNewPackage.Name, packages)
	}

	purrgilconfig.SaveFile()
	dockercompose.SaveFile()
	gitignore.SaveFile()

	ishell.PurrgilAlert(purrgilNewPackage.Name + " was successfuly added")
}
