package commands

import (
	"github.com/purrgil/purrgil/file"
	"github.com/purrgil/purrgil/interactiveshell"
	"os/exec"

	"os"
	"errors"
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

	if purrgilconfig.Name == "" {
		err := "Please enter with a name for your purrgil application!"
		ishell.Err(err, errors.New(err))
	}

	for _, value := range purrgilconfig.Packages {
		packageVerify(path, value)
	}

	ishell.PurrgilAlert("Starting your application, this may take some time :) ....")

	cmd := exec.Command("docker-compose", "-p", purrgilconfig.Name , "up", "--no-recreate", "-d")
	err := cmd.Run()

	if err != nil {
		ishell.Err("Docker compose fail in Up your application", err)
	}

	ishell.PurrgilAlert("Your application are running!")

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
