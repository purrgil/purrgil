package commands

import (
	"fmt"
	"regexp"
	"errors"
	"os"
	"strconv"
	"github.com/purrgil/purrgil/configs"
	"github.com/purrgil/purrgil/file"
	"github.com/purrgil/purrgil/interactiveshell"
)

func PackageList(opts configs.CommandPackageConfig) {
	path, _ := os.Getwd()
	purrgilconfig := file.NewPurrgil(path, "")
	mypackages := purrgilconfig.Packages

	ishell.PurrgilAlert("Listing packages from " + purrgilconfig.Name + "...")

	for k, v := range opts.FilterSettings {
		if hasKey(file.PurrgilAvaibleFilters, k) {
			mypackages = file.PurrgilGetPackages(mypackages, func(pkg file.PurrgilPackage) bool {
				var val string

				switch k {
				case "name":
					val = pkg.Name
				case "identity":
					val = pkg.Identity
				case "service":
					val = strconv.FormatBool(pkg.Service)
				case "ssh":
					val = strconv.FormatBool(pkg.SSH)
				case "provider":
					val = pkg.Provider
				}

				m, err := regexp.MatchString(v, val)

				if err != nil {
					regError := "Error in regex"
					ishell.Err(regError, errors.New(regError))
				}

				return m
			})
		}
	}

	if len(mypackages) == 0 {
		ishell.PurrgilAlert("Not found packages from this filters :(")
	}

	for _, value := range mypackages {
		var template string

		if value.Provider != "dockerhub" {
			template = "Package: %s into [./%s] from %s  \n"
		} else {
			template = "Package: %s image named as %s on %s \n"
		}

		formatedString := fmt.Sprintf(template, value.Identity, value.Name, value.Provider)
		ishell.PurrgilAlert(formatedString)
	}
}

func hasKey(list []string, k string) bool {
	for _, v := range list {
		if v == k {
			return true
		}
	}

	return false
}
