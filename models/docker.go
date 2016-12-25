package models

import (
	"github.com/guidiego/gommunicate-shell"
	"github.com/guidiego/purrgil/utils"
)

type DockerComposeFile struct {
	Version  string                           `yaml:"version,omitempty"`
	Services map[string]DockerComposeServices `yaml:"services,omitempty"`
}

type DockerComposeServices struct {
	Build     string   `yaml:"build,omitempty"`
	Image     string   `yaml:"image,omitempty"`
	Command   string   `yaml:"command,omitempty"`
	Ports     []string `yaml:"ports,omitempty"`
	Volumes   []string `yaml:"volumes,omitempty"`
	DependsOn []string `yaml:"depends_on,omitempty"`
}

func (d *DockerComposeFile) AddService(pkg PurrgilPackage) {
	var dcs = DockerComposeServices{}
	var serviceName = gshell.Ask("What's the name of your service", pkg.Name)

	dcs.Command = gshell.Ask("Write container command", "")
	dcs.Ports = gshell.AskTilBlankEnter("Tell the ports that you want to export, ex. '8000:8000'", "(enter to skip)")
	dcs.Volumes = gshell.AskTilBlankEnter("Tell the volumes that you want, ex. '.:/code'", "(enter to skip)")
	dcs.DependsOn = gshell.AskTilBlankEnter("Your image depends of another service, ex. 'db'", "(enter to skip)")

	if pkg.Provider == "github" {
		dcs.Build = "./" + pkg.Name + "/Dockerfile"
	} else {
		dcs.Image = pkg.Identity
	}

	if len(d.Services) == 0 {
		d.Services = map[string]DockerComposeServices{}
	}

	d.Services[serviceName] = dcs
}

func (d *DockerComposeFile) TryLinkUtil(pkg PurrgilPackage) {
	println("Adicionou um util")
}

func (d *DockerComposeFile) Load() {
	var configPath = path.DockerCompose()
	utils.ParseYamlFile(configPath, d)
}

func (d *DockerComposeFile) Save() {
	var configPath = path.DockerCompose()
	utils.SaveYamlFile(configPath, d)
}
