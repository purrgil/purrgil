package commands

import (
	"github.com/purrgil/purrgil/file"
	"github.com/purrgil/purrgil/interactiveshell"
	"github.com/purrgil/purrgil/configs"
	"os/exec"
	"fmt"
	"os"
	"errors"
	"strings"
)

func Up(opts configs.UpConfig) {
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

	if opts.NoCache {
		removeDeps := exec.Command("rm", "-rf", "./*/.deps")
		rmDepErr := removeDeps.Run()

		if rmDepErr != nil {
			ishell.Err("Problems in remove deps", rmDepErr)
		}
	}

	for serviceName, serviceData := range dc.Services {
		for _, vol := range serviceData.Volumes {
			volName := strings.Split(vol, ":")

			if checkDep(volName[0], serviceName) {
				moveToDeps(volName[0], serviceName)
			}
		}
	}

	ishell.PurrgilAlert("Starting your application, this may take some time :) ....")
	composeCmd := exec.Command("docker-compose", "up", "--no-recreate", "-d")
	composeErr := composeCmd.Run()

	if composeErr != nil {
		ishell.Err("Problem with compose", composeErr)
	}

	for _, pkg := range purrgilconfig.Packages {
		for _, command := range pkg.PostRunCommands {
			ishell.PurrgilAlert(fmt.Sprintf("%s || running '%s'", pkg.Name, command))

			command = "\"" + command + "\""
			commands := []string{ "exec", pkg.Name, "su", "root", "-c", command}
			value, err  := exec.Command("docker-compose", commands...).Output()

			if err != nil {
				ishell.Err(pkg.Name + " || " + string(value), err)
			} else {
				ishell.PurrgilAlert(pkg.Name + " || " + string(value))
			}
		}
	}

	ishell.PurrgilAlert("Your application are online!")

}

func moveToDeps(volPath string, serviceName string) {
	depsPath := fmt.Sprintf("./%s/.deps/", serviceName)

	ishell.PurrgilAlert(fmt.Sprintf("Add Dep: %s -> %s", volPath, depsPath))

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

func checkDep(name string, serviceName string) bool {
	servicePath := "./" + serviceName
	depsPath := servicePath + "/.deps/" + name
	return name != servicePath && exists(servicePath) && !exists(depsPath)
}


func exists(path string) bool {
    _, err := os.Stat(path)
    if err == nil { return true }

    return false
}
