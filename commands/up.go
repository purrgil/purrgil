package commands

import (
	"github.com/purrgil/purrgil/file"
	"github.com/purrgil/purrgil/interactiveshell"
	"os/exec"
	"fmt"
	"os"
	"errors"
	"strings"
)

func Up() {
	ishell.PurrgilAlert("Reading informations....")

	path, _ := os.Getwd()
	dcpath := path + "/docker-compose.yml"
	purrgilpath := path + "/purrgil.yml"

	if !exists(purrgilpath) {
		err := "We think that this isn't a Purrgil Project :/"
		ishell.Err(err, errors.New(err))
	}

	if !exists(dcpath) {
		err := "Your docker-compose not exist"
		ishell.Err(err, errors.New(err))
	}

	purrgilconfig := file.NewPurrgil(path, "")
	dc := file.NewDockerCompose(path)

	if purrgilconfig.Name == "" {
		err := "Please enter with a name for your purrgil application!"
		ishell.Err(err, errors.New(err))
	}

	for _, value := range purrgilconfig.Packages {
		packageVerify(path, value)
	}

	ishell.PurrgilAlert("Starting your application, this may take some time :) ....")

	mappedVolumes := dc.GetNamedVolumes()

	for serviceName, serviceData := range dc.Services {
		for _, vol := range serviceData.Volumes {
			volName := strings.Split(vol, ":")
			volPath := strings.Replace(mappedVolumes[volName[0]], "./", "", 1)
			moveToDeps(volName[0], path + "/" + volPath, path + "/" + serviceName)
		}
	}

	composeCmd := exec.Command("docker-compose", "up", "--no-recreate", "-d")
	composeErr := composeCmd.Run()

	if composeErr != nil {
		ishell.Err("Problem with compose", composeErr)
	}

	for _, pkg := range purrgilconfig.Packages {
		for _, command := range pkg.DevCommands {
			composeExecCommand := exec.Command("docker-compose", "exec", pkg.Name, command)
			composeExecErr := composeExecCommand.Run()

			if composeExecErr != nil {
				ishell.Err("Problem with compose", composeExecErr)
			}
		}
	}

	ishell.PurrgilAlert("Your application are online!")

}

func moveToDeps(volName string, volPath string, serviceName string) {
	depsPath := serviceName + "/.deps/"

	fmt.Println("cp", "-rv", volPath, depsPath)
	copyFile := exec.Command("cp", "-rv", volPath, depsPath)
	err := copyFile.Run()

	if err != nil {
		ishell.Err("Problem in copy files", err)
	}
}

func packageVerify(path string, pkg file.PurrgilPackage) {
	if pkg.Provider == "github" && !exists(path + "/" + pkg.Name) {
		err := "We not found package: [" + pkg.Name + "], try a 'purrgil install'"
		ishell.Err(err, errors.New(err))
	}
}


func exists(path string) bool {
    _, err := os.Stat(path)
    if err == nil { return true }

    return false
}
