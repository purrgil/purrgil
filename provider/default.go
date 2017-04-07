package provider

DefaultProviders := map[string]GitProvider{
	"github": GitProvider{
		SshTemplate: "git@github.com:%s/%s.git",
		HttpsTemplate: "https://github.com/%s/%s.git",
		Name: "github",
	},
	"gitlab": GitProvider{
		SshTemplate: "git@gitlab.com:%s/%s.git",
		HttpsTemplate: "https://%s@gitlab.com/%s/%s.git",
		Name: "gitlab",
	},
	"bitbucket": GitProvider{
		SshTemplate: "git@bitbucket.org:%s/%s.git",
		HttpsTemplate: "https://%s@bitbucket.org/%s/%s.git",
		Name: "bitbucket",
	}
}
