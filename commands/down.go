package commands

import (
	"github.com/purrgil/purrgil/interactiveshell"
	"os/exec"
)

func Down() {
	ishell.PurrgilAlert("Starting to drop containers....")

	composeDown := exec.Command("docker-compose", "down")
	composeErr := composeDown.Run()

	if composeErr != nil {
		ishell.Err("Docker fail in kill containers", composeErr)
	}

	rmDeps := exec.Command("rm", "-rf", "./*/.deps")
	rmErr := rmDeps.Run()

	if rmErr != nil {
		ishell.Err("Docker fail in kill containers", rmErr)
	}

	ishell.PurrgilAlert("Your application goes down! To start again: 'purrgil up'")

}
