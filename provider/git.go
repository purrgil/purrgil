package provider

import (
	"os/exec"

	"github.com/op/go-logging"
	msg "github.com/purrgil/purrgil/log/message"
)

var log = logging.MustGetLogger("purrgil.provider.git")

type GitProvider struct {
	SshTemplate string
	HttpsTemplate string
	Name string
}

type GitCloneConfig struct {
	User string
	Repo string
	FolderName string
	IsSSH bool
}

func (gp *GitProvider) Clone(config GitCloneConfig) {
	cloneURI := getCloneURI(config)

	log.Infof(msg.START_CLONE, config.User, config.Repo)

	cmd := exec.Command("git", "clone", cloneURI, config.FolderName)
	output, err := cmd.Output()

	if output, err := cmd.Output(); err != nil {
		log.Error(msg.FAIL_CLONE)
		log.Fatalf(err)
	}

	log.Infof(msg.END_CLONE, config.User, config.Repo, config.FolderName)
}

func (gp *GitProvider) getCloneURI(config GitCloneConfig) string {
	if (config.IsSSH) {
		return getSshURI(config.User, config.Repo)
	}

	return getHttpsURI(config.User, config.Repo)
}

func (gp *GitProvider) getSshURI(user string, repo string) string {
	return fmt.Sprintf(gp.SshTemplate,  user,  repo)
}

func (gp *GitProvider) getHttpsURI(user string, repo string) string {
	return fmt.Sprintf(gp.HttpsTemplate,  user,  repo)
}
