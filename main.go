package main

import (
	"os"

	"github.com/purrgil/purrgil/commands"
	"github.com/purrgil/purrgil/configs"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	avaibleProviders = []string{"github", "bitbucket", "dockerhub", "bitbucket_mcr"}
	app = kingpin.New("purrgil", "Bleh")

	install = app.Command("install", "Install Purrgil Project")
	up = app.Command("up", "Mount your application")
	down = app.Command("down", "Drop your application")
	upgrade = app.Command("upgrade", "Update your Purrgil version")
	version = app.Command("version", "Show purrgil version")

	push = app.Command("push", "commit/push a env update on git")
	pushMsg = push.Arg("commit message", "your commit message").String()

	packages   = app.Command("packages", "List all installed container packages")
	packageFilter = packages.Flag("filter", "filter packages by attributes").Short('f').StringMap()

	initCommand = app.Command("init", "Init purrgil.yml")
	initProjectName = initCommand.Arg("project name", "Name of the purrgil project").String()
	initProjectProvider = initCommand.Flag("provider", "select a provider to your project").Short('p').Enum(avaibleProviders...)
	initProjectRepo = initCommand.Flag("repository", "the repository to be downloaded in your provider").Short('r').String()
	initProjectHttps = initCommand.Flag("https", "you can use that if want https protocol to download").Bool()

	deploy  = app.Command("deploy", "Make project deploy")
	deployC = deploy.Flag("container", "Deploy a single container").String()

	add     = app.Command("add", "Add a dependency to project")
	addPkgIdentity = add.Arg("pkg", "Add a service").String()
	addServiceFlag  = add.Flag("not-a-service", "Add only a git repository").Bool()
	addProvider   = add.Flag("provider", "Install image directly from dockerhub").Enum(avaibleProviders...)
	addName = add.Flag("name", "Give a custom name to package").String()
	addHTTPS = add.Flag("https", "download package in https mode").Bool()
	addComposeConfig  = add.Flag("compose-helper", "Active an interface to inject basic compose infos").Bool()

	remove  = app.Command("remove", "Remove a dependency to project")
	removeP = remove.Arg("package", "Name of Purrgil Package to Remove").String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case install.FullCommand():
		commands.Install()

	case initCommand.FullCommand():
		commands.Init(*initProjectName, configs.InitConfig{
			Repo: *initProjectProvider,
			IsSSH: !*initProjectHttps,
			Provider: *initProjectRepo,
		})

	case deploy.FullCommand():
		commands.Deploy()

	case add.FullCommand():
		commands.Add(*addPkgIdentity, configs.AddConfig{
			IsService:  *addServiceFlag,
			Provider:  *addProvider,
			CustomName: *addName,
			ComposeConfig: *addComposeConfig,
			HttpsMode: *addHTTPS,
		})

	case push.FullCommand():
		commands.Push(*pushMsg)

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
			FilterSettings: *packageFilter,
		})
	}
}
