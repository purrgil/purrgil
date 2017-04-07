package commands

import (
	"github.com/op/go-logging"
	"github.com/purrgil/file"
)

var log = logging.MustGetLogger("purrgil.command.add")

type AddParams struct {
	From string
	IsService bool
	Name string
	Provider string
	IsHTTPS bool
	Helper bool
}

func Add(params AddParams) {
	pf := file.GetPurrgilFile()

	pf.AddPackage(params.Name, purrgil.PurrgilPackage{
		Local,
		purrgil.PurrgilPackageDownload{
			Provider: params.Provider,
			From: params.From,
			SSH: !params.IsHTTPS,
			PostInstall: []string{},
		}
		PurrgilPackageDeploy{},
		[]string{},
		map[string]string{},
		map[string]string{},
	})

	InstallPackage(AddParams)
	pf.Save()
}
