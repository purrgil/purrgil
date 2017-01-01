package commands

import (
	"os"
	"os/exec"

	"github.com/guidiego/purrgil/file"
	"github.com/guidiego/purrgil/interactiveshell"
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

	if pkg.Provider == "github" {
		GitClone(pkg.Identity)
	} else {
		DockerPull(pkg.Identity)
	}

	ishell.PurrgilAlert(pkg.Name + " was successfuly installed")
}

func GitClone(repo string) {
	ssh := "git@github.com:" + repo + ".git"
	cmd := exec.Command("git", "clone", ssh)
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
