package file

import (
	"strings"
	"github.com/guidiego/purrgil/configs"
)

type PurrgilPackage struct {
	Name     string `yaml:"name"`
	Identity string `yaml:"identity"`
	Provider string `yaml:"provider"`
	Service  bool   `yaml:"service"`
}

func NewPurrgilPackage(id string, opts configs.AddConfig) PurrgilPackage {
	pkg := PurrgilPackage{}

	pkg.Identity = id
	pkg.Name = normalizeName(id, opts.CustomName)
	pkg.Provider = getProvider(opts.Dockerhub)
	pkg.Service = !opts.IsService

	return pkg
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
