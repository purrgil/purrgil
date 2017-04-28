package provider

var DefaultProviders = map[string]GitProvider{
	"github": GitProvider{
		SSHTemplate:   "git@github.com:%s/%s.git",
		HTTPSTemplate: "https://github.com/%s/%s.git",
		Name:          "github",
	},
	"gitlab": GitProvider{
		SSHTemplate:   "git@gitlab.com:%s/%s.git",
		HTTPSTemplate: "https://%s@gitlab.com/%s/%s.git",
		Name:          "gitlab",
	},
	"bitbucket": GitProvider{
		SSHTemplate:   "git@bitbucket.org:%s/%s.git",
		HTTPSTemplate: "https://%s@bitbucket.org/%s/%s.git",
		Name:          "bitbucket",
	},
}
