package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	cmd "github.com/purrgil/purrgil/command"
	plog "github.com/purrgil/purrgil/log"
	msg "github.com/purrgil/purrgil/log/message"
)

var (
	app = kingpin.New("purrgil", msg.PURRGIL_DESCRIPTION)

	debugMode = app.Flag("debug", msg.DEBUG_FLAG_DESCRIPTION).Short('D').Bool()

	// purrgil install
	install = app.Command("install", msg.INSTALL_CMD_DESCRIPTION)

	// purrgil add <pkg> --not-a-service --name --provider --https --compose-helper
	add = app.Command("add", msg.ADD_CMD_DESCRIPTION)
	addParams = cmd.AddParams{
		From: app.Arg("pkg", msg.ADD_PKG_DESCRIPTION).String(),
		IsService: !(*app.Flag("not-a-service", msg.ADD_NS_FLAG_DESCRIPTION).Bool()),
		Name: *app.Flag("name", msg.ADD_NAME_FLAG_DESCRIPTION).Shot("N").String(),
		Provider: *app.Flag("provider", msg.ADD_PROVIDER_FLAG_DESCRIPTION).Short("P").String(),
		IsHTTPS: *app.Flag("https", msg.ADD_HTTPS_FLAG_DESCRIPTION).Bool(),
		Helper: *app.Flag("compose-helper", msg.ADD_HELPER_FLAG_DESCRIPTION).Bool(),
	}

)

func main() {
	p := kingpin.MustParse(app.Parse(os.Args[1:]))

	plog.ConfigureLoggin(*debugMode)

	switch p {
	case install.FullCommand():
		cmd.Install(installFlags)
	case add.FullCommand():
		cmd.Add(addFlags)
	}
}
