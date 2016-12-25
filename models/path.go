package models

import (
	"log"
	"os"
)

type Path struct{}

func (p *Path) Resolve(path string) string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return dir + path
}

func (p *Path) DockerCompose() string {
	path := p.Resolve("/docker-compose.yml")
	return path
}

func (p *Path) PurrgilConfig() string {
	path := p.Resolve("/purrgil.yml")
	return path
}

func (p *Path) ProjectRoot(project string) string {
	path := p.Resolve("/" + project)
	return path
}

func (p *Path) ProjectDockerfile(project string) string {
	path := p.ProjectRoot(project) + "/Dockerfile"
	return path
}
