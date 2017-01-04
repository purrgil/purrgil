package main

import (
	"os"

	"github.com/purrgil/purrgil/commands"
	"github.com/purrgil/purrgil/configs"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("purrgil", "Bleh")

	install = app.Command("install", "Install Purrgil Project")
	up = app.Command("up", "Mount your application")
	down = app.Command("down", "Drop your application")
	upgrade = app.Command("upgrade", "Update your Purrgil version")
	version = app.Command("version", "Show purrgil version")

	packages   = app.Command("packages", "List all installed container packages")
	pkgGit     = packages.Flag("github", "Filter only GITHUB provider packages").Bool()
	pkgDock    = packages.Flag("dockerhub", "Filter only DOCKERHUB provider packages").Bool()
	pkgService = packages.Flag("services", "Filter only SERVICES packages").Bool()
	pkgNormal  = packages.Flag("non-service", "Filter only NON SERVICE packages").Bool()

	initM = app.Command("init", "Init purrgil.yml")
	pName = initM.Arg("project name", "Name of the purrgil project").String()
	pRepo = initM.Flag("github", "github repository for your env").String()

	deploy  = app.Command("deploy", "Make project deploy")
	deployC = deploy.Flag("container", "Deploy a single container").String()

	add     = app.Command("add", "Add a dependency to project")
	addS    = add.Arg("pkg", "Add a service").String()
	addNs   = add.Flag("not-a-service", "Add only a git repository").Bool()
	addDk   = add.Flag("dockerhub", "Install image directly from dockerhub").Bool()
	addName = add.Flag("name", "Give a custom name to package").String()
	addDcConfig  = add.Flag("compose-helper", "Active an interface to inject basic compose infos").Bool()

	remove  = app.Command("rm", "Remove a dependency to project")
	removeP = remove.Arg("package", "Name of Purrgil Package to Remove").String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case install.FullCommand():
		commands.Install()

	case initM.FullCommand():
		commands.Init(*pName, *pRepo)

	case deploy.FullCommand():
		commands.Deploy()

	case add.FullCommand():
		commands.Add(*addS, configs.AddConfig{
			IsService:  *addNs,
			Dockerhub:  *addDk,
			CustomName: *addName,
			ComposeConfig: *addDcConfig,
		})

	case version.FullCommand():
		commands.Version()

	case remove.FullCommand():
		commands.Remove(*removeP)

	case up.FullCommand():
		commands.Up()

	case down.FullCommand():
		commands.Down()

	case upgrade.FullCommand():
		commands.Upgrade()

	case packages.FullCommand():
		commands.PackageList(configs.CommandPackageConfig{
			IsGithub:    *pkgGit,
			IsDockerhub: *pkgDock,
			IsService:   *pkgService,
			IsNormal:    *pkgNormal,
		})
	}
}
