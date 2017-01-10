package commands

import (
	"os"
	"os/exec"
	"strings"
	"fmt"
	"errors"

	"github.com/purrgil/purrgil/file"
	"github.com/purrgil/purrgil/interactiveshell"
)

type ProviderMap struct {
	ssh string
	https string
}

var (
	githubProvider = ProviderMap{
		ssh: "git@github.com:%s.git",
		https: "https://github.com/%s.git",
	}
	bitbucketProvider = ProviderMap{
		ssh: "git@bitbucket.org:%s.git",
		https: "https://%s@bitbucket.org/%s.git",
	}
	bitbucketMcrProvider = ProviderMap{
		ssh: "ssh://hg@bitbucket.org/%s",
		https: "https://%s@bitbucket.org/%s",
	}
)

func Install() {
	ishell.PurrgilAlert("Starting Purrgil Install...")

	wdpath, _ := os.Getwd()
	purrgil := file.NewPurrgil(wdpath, "")

	for _, pkg := range purrgil.Packages {
		PackageInstall(pkg)
	}

}

func PackageInstall(pkg file.PurrgilPackage) {
	ishell.PurrgilAlert("Installing " + pkg.Name + " package...")

	RunDownloadFromProvider(pkg)

	ishell.PurrgilAlert(pkg.Name + " was successfuly installed")
}

func RunDownloadFromProvider(pkg file.PurrgilPackage) {
	switch pkg.Provider {
	case "github":
		GithubClone(pkg, githubProvider)
	case "bitbucket":
		BitbucketClone("git", pkg, bitbucketProvider)
	case "bitbucket_mcr":
		BitbucketClone("hg", pkg, bitbucketMcrProvider)
	case "dockerhub":
		DockerPull(pkg.Identity)
	default:
		notFoundProvider := "We not found this repo! Try: github, bitbucket, dockerhub or bitbucket_mcr (support mercurial)"
		ishell.Err(notFoundProvider, errors.New(notFoundProvider))
	}
}

func GithubClone(pkg file.PurrgilPackage, kind ProviderMap) {
	var template string

	if !pkg.SSH {
		template = fmt.Sprintf(kind.https, pkg.Identity)
	} else {
		template = fmt.Sprintf(kind.ssh, pkg.Identity)
	}

	println(template)
	RepoClone("git", template, pkg.Name)
}

func BitbucketClone(providerCommand string, pkg file.PurrgilPackage, kind ProviderMap) {
	var template string

	if !pkg.SSH {
		user, repoId := getBitBucketParms(pkg.Identity)
		template = fmt.Sprintf(kind.https, user, repoId)
	} else {
		template = fmt.Sprintf(kind.ssh, pkg.Identity)
	}


	RepoClone(providerCommand, template, pkg.Name)
}

func getBitBucketParms(id string) (string, string) {
	splited := strings.Split(id, ":")
	return splited[0], splited[1]
}

func RepoClone(providerCmd string, repo string, folderName string) {
	cmd := exec.Command(providerCmd, "clone", repo , folderName)
	err := cmd.Run()

	if err != nil {
		ishell.Err("Somethings goes, confeer the identity of your package", err)
	}
}

func DockerPull(repo string) {
	cmd := exec.Command("docker", "pull", repo)
	err := cmd.Run()

	if err != nil {
		ishell.Err("Somethings goes, confeer the identity of your package", err)
	}
}
