package commands

import (
	"github.com/purrgil/file-parser/purrgil"
)

// var log = logging.MustGetLogger("purrgil.command.install")

func Install() {

}

func PkgDownloadToAddParams(name string, isService bool, pkg purrgil.PurrgilPackageDownload) AddParams {
	return AddParams{
		From:      pkg.From,
		IsService: isService,
		Name:      name,
		Provider:  pkg.Provider,
		IsHTTPS:   !pkg.SSH,
		Helper:    false,
	}
}

func InstallPackage(params AddParams) {

}
