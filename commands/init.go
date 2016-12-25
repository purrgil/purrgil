package commands

import (
	"os"
	"fmt"

	"github.com/guidiego/purrgil/models"
	"github.com/fatih/color"
)

func Init(projectName string) {
	magenta := color.New(color.FgMagenta).SprintFunc()
	fmt.Printf("%s starting your app....\n", magenta("PURRGIL"))

	dockercompose := models.DockerComposeFile{}
	purrgil := models.Purrgil{}

	dockercompose.Version = "2"
	purrgil.Name = projectName

	os.Mkdir(projectName, 0777)
	os.Chdir(projectName)

	dockercompose.Save()
	purrgil.Save()

	fmt.Printf("%s app created! To enter in project give 'cd %s' \n", magenta("PURRGIL"), projectName)

}
