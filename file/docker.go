package file

type DockerComposeFile struct {
	yaml     Yaml
	Version  string                          `yaml:"version,omitempty"`
	Services map[string]DockerComposeService `yaml:"services,omitempty"`
}

type DockerComposeService struct {
	Build     string   `yaml:"build,omitempty"`
	Image     string   `yaml:"image,omitempty"`
	Command   string   `yaml:"command,omitempty"`
	Ports     []string `yaml:"ports,omitempty"`
	Volumes   []string `yaml:"volumes,omitempty"`
	DependsOn []string `yaml:"depends_on,omitempty"`
	Links []string `yaml:"links,omitempty"`
}

func (d *DockerComposeFile) AddService(key string, dcs DockerComposeService) {
	if len(d.Services) == 0 {
		d.Services = map[string]DockerComposeService{}
	}

	d.Services[key] = dcs
}

func (d *DockerComposeFile) LinkInService(volumeName string, services []string) {
	for _, val := range services {
		service := d.Services[val]
		service.Volumes = append(service.Volumes, ".:/"+volumeName)

		d.Services[val] = service
	}
}

func (d *DockerComposeFile) SaveFile() {
	d.yaml.SaveFile(d)
}

func (d *DockerComposeFile) LoadFile() {
	d.yaml.LoadFile(d)
}

func (d *DockerComposeFile) RemoveFromServices(pkgName string) {
	delete(d.Services, pkgName)
}

func NewDockerCompose(dir string) DockerComposeFile {
	dockercompose := DockerComposeFile{}

	dockercompose.Version = "2"

	dockercompose.yaml.InitFile(dir, "docker-compose")
	dockercompose.yaml.LoadFile(&dockercompose)

	return dockercompose
}
