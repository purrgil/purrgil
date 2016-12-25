package models

import (
	"github.com/guidiego/purrgil/utils"
)

var path = Path{}

type Purrgil struct {
	Name     string           `yaml:"name"`
	Packages []PurrgilPackage `yaml:"packages,omitempty"`
	Settings PurrgilSettings  `yaml:"settings,omitempty"`
}

type PurrgilPackage struct {
	Name     string `yaml:"name"`
	Identity string `yaml:"identity"`
	Provider string `yaml:"provider"`
	Service  bool   `yaml:"service"`
}

// TODO: Deploy settings and dev envs
type PurrgilSettings struct {
}

func (p *Purrgil) AddPackage(pkg PurrgilPackage) {
	p.Packages = append(p.Packages, pkg)
}

func (p *Purrgil) Load() {
	var configPath = path.PurrgilConfig()
	utils.ParseYamlFile(configPath, p)
}

func (p *Purrgil) Save() {
	var configPath = path.PurrgilConfig()
	utils.SaveYamlFile(configPath, p)
}
