package provider

import (
	"fmt"
	"os/exec"

	"github.com/op/go-logging"
	msg "github.com/purrgil/purrgil/log/message"
)

var log = logging.MustGetLogger("purrgil.provider.git")

// GitProvider is a struct that contains core infos about GitProvider
type GitProvider struct {
	SSHTemplate   string
	HTTPSTemplate string
	Name          string
}

// GitCloneConfig is a struct will params needed to clone a repo
type GitCloneConfig struct {
	User       string
	Repo       string
	FolderName string
	IsSSH      bool
}

// Clone method get a GitCloneConfig and exec clone func from Git cli
func (gp *GitProvider) Clone(config GitCloneConfig) {
	cloneURI := gp.getCloneURI(config)

	log.Infof(msg.START_CLONE, config.User, config.Repo)

	cmd := exec.Command("git", "clone", cloneURI, config.FolderName)
	output, err := cmd.Output()

	if err != nil {
		log.Error(msg.FAIL_CLONE)
		log.Fatalf(err.Error())
	}

	log.Infof(msg.END_CLONE, config.User, config.Repo, config.FolderName)
	log.Debug(output)
}

func (gp *GitProvider) getCloneURI(config GitCloneConfig) string {
	if config.IsSSH {
		return gp.getSSHURI(config.User, config.Repo)
	}

	return gp.getHTTPSURI(config.User, config.Repo)
}

func (gp *GitProvider) getSSHURI(user string, repo string) string {
	return fmt.Sprintf(gp.SSHTemplate, user, repo)
}

func (gp *GitProvider) getHTTPSURI(user string, repo string) string {
	return fmt.Sprintf(gp.HTTPSTemplate, user, repo)
}
