package models;

import (
    "fmt"
    "github.com/guidiego/purrgil/utils"
    "github.com/guidiego/gommunicate-shell"
)

type DockerComposeFile struct {
    Version string `yaml:"version"`
    Services map[string]DockerComposeServices `yaml:"services"`
}

type DockerComposeServices struct {
    Build      string    `yaml:"build"`
    Image      string    `yaml:"image"`
    Command    string    `yaml:"command"`
    Ports      []string  `yaml:"ports"`
    Volumes    []string  `yaml:"volumes"`
    DependsOn  []string  `yaml:"depends_on"`
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

  fmt.Println(serviceName)
  fmt.Printf("%+v\n", dcs)

  d.Services[serviceName] = dcs
  fmt.Println(d.Services)
}

func (d *DockerComposeFile) TryLinkUtil(pkg PurrgilPackage) {
  println("Adicionou um util")
}

func (d *DockerComposeFile) Load() {
    var configPath = path.DockerCompose();
    utils.ParseYamlFile(configPath, d)
}

func (d *DockerComposeFile) Save() {
    var configPath = path.DockerCompose();
    utils.SaveYamlFile(configPath, d)
}
