package commands

import (
	"github.com/guidiego/purrgil/models"
	"strings"
)

func Add(pkdId string, opts models.AddConfig) {
	println("Add Command")

	purrgilconfig := models.Purrgil{}
	dockercompose := models.DockerComposeFile{}
	purrgilNewPackage := models.PurrgilPackage{}

	purrgilconfig.Load()
	dockercompose.Load()
	mountPurrgilPackage(&purrgilNewPackage, pkdId, opts)

	purrgilconfig.AddPackage(purrgilNewPackage)

	if purrgilNewPackage.Service {
		dockercompose.AddService(purrgilNewPackage)
	} else {
		dockercompose.TryLinkUtil(purrgilNewPackage)
	}

	purrgilconfig.Save()
	dockercompose.Save()
}

func mountPurrgilPackage(pkg *models.PurrgilPackage, id string, opts models.AddConfig) {
	pkg.Identity = id
	pkg.Name = normalizeName(id, opts.CustomName)
	pkg.Provider = getProvider(opts.Dockerhub)
	pkg.Service = !opts.IsService
}

func normalizeName(id string, custom string) string {
	if custom != "" {
		return custom
	}

	if strings.Contains(id, "/") {
		splited := strings.Split(id, "/")
		lastValue := len(splited) - 1

		return splited[lastValue]
	}

	return id
}

func getProvider(isDockerhub bool) string {
	if isDockerhub {
		return "dockerhub"
	}

	return "github"
}
