package file

type DockerComposeFile struct {
	Yaml
	Version  string                          `yaml:"version"`
	Services map[string]DockerComposeService `yaml:"services"`
}

type DockerComposeService struct {
	Build     string   `yaml:"build"`
	Image     string   `yaml:"image"`
	Command   string   `yaml:"command"`
	Ports     []string `yaml:"ports"`
	Volumes   []string `yaml:"volumes"`
	DependsOn []string `yaml:"depends_on"`
}

func (d *DockerComposeFile) AddService(key string, dcs DockerComposeService) {
	if len(d.Services) == 0 {
		d.Services = map[string]DockerComposeService{}
	}

	d.Services[key] = dcs
}

func (d *DockerComposeFile) LinkInService(services []string) {

}

func NewDockerCompose(dir string) DockerComposeFile {
	dockercompose := DockerComposeFile{}

	dockercompose.Version = "2"

	dockercompose.InitFile(dir, "docker-compose")
	dockercompose.LoadFile()

	return dockercompose
}
