package commands;

import (
    "strings"
    "github.com/guidiego/purrgil/models"
)

func Add(pkdId string, opts models.AddConfig) {
    println("Add Command")

    purrgilconfig := models.Purrgil{}
    purrgilNewPackage := models.PurrgilPackage{}

    purrgilconfig.Load()
    // dockercompose := utils.GetDockerCompose()


    mountPurrgilPackage(&purrgilNewPackage, pkdId, opts)
    saveInFiles(purrgilNewPackage, &purrgilconfig)
}

func mountPurrgilPackage(pkg *models.PurrgilPackage, id string, opts models.AddConfig) {
    pkg.Identity = id
    pkg.Name = normalizeName(id, opts.CustomName)
    pkg.Provider = getProvider(opts.Dockerhub)
    pkg.Service = !opts.IsService
}

func saveInFiles(pkg models.PurrgilPackage, purrgilconfig *models.Purrgil) {
    purrgilconfig.AddPackage(pkg)
    purrgilconfig.Save()
}

func normalizeName(id string, custom string) string {
    if custom != "" {
        return custom
    }

    if (strings.Contains(id, "/")) {
        splited := strings.Split(id, "/")
        lastValue := len(splited) - 1

        return splited[lastValue]
    }

    return id
}

func getProvider(isDockerhub bool) string {
    if isDockerhub {
        return "dockerhub"
    }

    return "github"
}
