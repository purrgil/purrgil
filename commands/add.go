package commands

import (
	"github.com/guidiego/purrgil/file"
	"github.com/guidiego/purrgil/configs"
	"github.com/guidiego/purrgil/interactiveshell"
	"os"
)

func Add(pkdId string, opts configs.AddConfig) {
	println("Add Command")

	path, _ := os.Getwd()

	purrgilconfig := file.NewPurrgil(path, "")
	dockercompose := file.NewDockerCompose(path)
	purrgilNewPackage := file.NewPurrgilPackage(pkdId, opts)

	purrgilconfig.AddPackage(purrgilNewPackage)

	if purrgilNewPackage.Service {
		serviceName, service := ishell.CollectDockerServiceInfo(purrgilNewPackage)
		dockercompose.AddService(serviceName, service)
	} else {
		packages := ishell.CollectLinkPossibility(purrgilNewPackage)
		dockercompose.LinkInService(packages)
	}

	purrgilconfig.File.SaveFile()
	dockercompose.File.SaveFile()
}
