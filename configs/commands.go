package configs

type AddConfig struct {
	IsService  bool
	Provider  string
	HttpsMode bool
	CustomName string
	ComposeConfig bool
}

type CommandPackageConfig struct {
	FilterSettings map[string]string
}

type InitConfig struct {
	IsSSH  bool
	Repo  string
	Provider string
}
