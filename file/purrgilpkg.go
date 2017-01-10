package file

import (
	"github.com/purrgil/purrgil/configs"
	"strings"
)

type PurrgilPackage struct {
	Name     string `yaml:"name"`
	Identity string `yaml:"identity"`
	Provider string `yaml:"provider"`
	Service  bool   `yaml:"service"`
	SSH bool `yaml:"ssh"`
	DevCommands []string `yaml:"dev_commands,omitempty"`
}

var PurrgilAvaibleFilters = []string {
	"name",
	"identity",
	"provider",
	"service",
	"ssh",
}

func NewPurrgilPackage(id string, opts configs.AddConfig) PurrgilPackage {
	pkg := PurrgilPackage{}

	pkg.Identity = id
	pkg.Name = normalizeName(id, opts.CustomName)
	pkg.Provider = getProvider(opts.Provider)
	pkg.Service = !opts.IsService
	pkg.SSH = !opts.HttpsMode

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

func getProvider(provider string) string {
	if provider == "" {
		return "github"
	}

	return provider
}
