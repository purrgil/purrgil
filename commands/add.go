package commands

import (
	"github.com/purrgil/purrgil/configs"
	"github.com/purrgil/purrgil/file"
	"github.com/purrgil/purrgil/interactiveshell"
	"os"
)

func Add(pkgId string, opts configs.AddConfig) {
	ishell.PurrgilAlert("Starting Purrgil Add on Package: " + pkgId)

	path, _ := os.Getwd()

	purrgilconfig := file.NewPurrgil(path, "")
	dockercompose := file.NewDockerCompose(path)
	gitignore := file.NewGitIgnore(path)
	purrgilNewPackage := file.NewPurrgilPackage(pkgId, opts)

	PackageInstall(path, purrgilNewPackage)

	purrgilconfig.AddPackage(purrgilNewPackage)
	gitignore.AddIgnoredPath(purrgilNewPackage.Name)

	ishell.PurrgilAlert("Preparing to Collect Aditional Information...")

	if opts.ComposeConfig {
		callComposeConfigInterface(purrgilNewPackage, &dockercompose)
	}

	purrgilconfig.SaveFile()
	dockercompose.SaveFile()
	gitignore.SaveFile()

	ishell.PurrgilAlert(purrgilNewPackage.Name + " was successfuly added")
}

func callComposeConfigInterface(pkg file.PurrgilPackage, dc *file.DockerComposeFile) {
	if pkg.Service {
		serviceName, service := ishell.CollectDockerServiceInfo(pkg)
		dc.AddService(serviceName, service)
	} else {
		packages := ishell.CollectLinkPossibility(pkg)
		dc.LinkInService(pkg.Name, packages)
	}
}
