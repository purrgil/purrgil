package models;

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
