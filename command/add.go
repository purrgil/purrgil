package commands

import (
	"github.com/op/go-logging"
	"github.com/purrgil/file-parser/purrgil"
	"github.com/purrgil/purrgil/file"
)

var log = logging.MustGetLogger("purrgil.command.add")

type AddParams struct {
	From      string
	IsService bool
	Name      string
	Provider  string
	IsHTTPS   bool
	Helper    bool
}

func Add(params AddParams) {
	pf, err := file.GetPurrgilFile()

	if err != nil {
		log.Fatalf(err.Error())
	}

	pf.Data.AddPackage(params.Name, purrgil.NewPkg(purrgil.PurrgilPackageDownload{
		Provider:    params.Provider,
		From:        params.From,
		SSH:         !params.IsHTTPS,
		PostInstall: []string{},
	}))

	InstallPackage(params)
	pf.Save()
}
