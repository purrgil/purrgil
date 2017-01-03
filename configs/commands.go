package configs

type AddConfig struct {
	IsService  bool
	Dockerhub  bool
	CustomName string
	ComposeConfig bool
}

type CommandPackageConfig struct {
	IsGithub    bool
	IsDockerhub bool
	IsService   bool
	IsNormal    bool
}
