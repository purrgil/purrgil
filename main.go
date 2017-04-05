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

	install = app.Command("install", msg.INSTALL_CMD_DESCRIPTION)

)

func main() {
	p := kingpin.MustParse(app.Parse(os.Args[1:]))

	plog.ConfigureLoggin(*debugMode)

	switch p {
	case install.FullCommand():
		cmd.Install()
	}
}
