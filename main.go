package main

import (
	"os"

	"github.com/guidiego/purrgil/commands"
	"github.com/guidiego/purrgil/configs"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("purrgil", "Bleh")

	install  = app.Command("install", "Install Purrgil Project")
	services = app.Command("services", "List all installed packages")

	initM   = app.Command("init", "Init purrgil.yml")
	pName   = initM.Arg("project name", "Name of the purrgil project").String()

	deploy  = app.Command("deploy", "Make project deploy")
	deployC = deploy.Flag("container", "Deploy a single container").String()

	add     = app.Command("add", "Add a dependency to project")
	addS    = add.Arg("pkg", "Add a service").String()
	addNs   = add.Flag("not-a-service", "Add only a git repository").Bool()
	addDk   = add.Flag("dockerhub", "Install image directly from dockerhub").Bool()
	addName = add.Flag("name", "Give a custom name to package").String()

	remove  = app.Command("rm", "Remove a dependency to project")
	removeV = remove.Flag("virtual", "Remove only from root not from dependencies").Bool()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case install.FullCommand():
		commands.Install()

	case initM.FullCommand():
		commands.Init(*pName)

	case deploy.FullCommand():
		commands.Deploy()

	case add.FullCommand():
		commands.Add(*addS, configs.AddConfig{
			IsService:  *addNs,
			Dockerhub:  *addDk,
			CustomName: *addName,
		})

	case remove.FullCommand():
		commands.Remove()

	case services.FullCommand():
		commands.Services()
	}
}
