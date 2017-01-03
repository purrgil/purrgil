package commands

import (
	"github.com/purrgil/purrgil/file"
	"github.com/purrgil/purrgil/interactiveshell"
	"os/exec"
	"strings"
	"os"
)

func Down() {
	ishell.PurrgilAlert("Starting to drop containers....")

	path, _ := os.Getwd()
	purrgilconfig := file.NewPurrgil(path, "")

	filter := "name=" + purrgilconfig.Name + "_"
	stdout, _ := exec.Command("docker", "ps", "--filter", filter, "-q").Output()
	containers := strings.Split(string(stdout), "\n")
	command := append([]string{"kill"}, containers...)

	cmd := exec.Command("docker", command...)
	err := cmd.Run()

	if err != nil {
		ishell.Err("Docker fail in kill containers", err)
	}

	ishell.PurrgilAlert("Your application goes down! To start again: 'purrgil up'")

}
