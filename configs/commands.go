package configs

type AddConfig struct {
	IsService  bool
	Dockerhub  bool
	CustomName string
}

type CommandPackageConfig struct {
	IsGithub    bool
	IsDockerhub bool
	IsService   bool
	IsNormal    bool
}
