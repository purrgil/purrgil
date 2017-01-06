package commands

import (
	"fmt"
	"github.com/purrgil/purrgil/configs"
	"github.com/purrgil/purrgil/file"
	"github.com/purrgil/purrgil/interactiveshell"
	"os"
)

func PackageList(opts configs.CommandPackageConfig) {
	path, _ := os.Getwd()
	purrgilconfig := file.NewPurrgil(path, "")
	mypackages := purrgilconfig.Packages

	ishell.PurrgilAlert("Listing packages from " + purrgilconfig.Name + "...")

	// if opts.IsGithub && !opts.IsDockerhub {
	// 	mypackages = file.PurrgilGetPackages(mypackages, filterGithub)
	// } else if !opts.IsGithub && opts.IsDockerhub {
	// 	mypackages = file.PurrgilGetPackages(mypackages, filterDockerhub)
	// }

	// if opts.IsService && !opts.IsNormal {
	// 	mypackages = file.PurrgilGetPackages(mypackages, filterService)
	// } else if !opts.IsService && opts.IsNormal {
	// 	mypackages = file.PurrgilGetPackages(mypackages, filterNonservice)
	// }

	if len(mypackages) == 0 {
		ishell.PurrgilAlert("Not found packages from this filters :(")
	}

	for _, value := range mypackages {
		fmt.Printf("  - %s <- %s \n", value.Name, value.Identity)
	}
}

func filterGithub(pkg file.PurrgilPackage) bool {
	if pkg.Provider == "github" {
		return true
	}

	return false
}

func filterService(pkg file.PurrgilPackage) bool {
	if pkg.Service {
		return true
	}

	return false
}

func filterDockerhub(pkg file.PurrgilPackage) bool {
	return !filterGithub(pkg)
}

func filterNonservice(pkg file.PurrgilPackage) bool {
	return !filterService(pkg)
}
