package utils;

import (
    "log"
    "os"
)

func resolve(path string) string {
    dir, err := os.Getwd()

    if err != nil {
        log.Fatal(err)
    }

    return dir + path
}

func GetPurrgilConfig() string {
    var path = resolve("/purrgil.yml")
    return path
}

func GetDockerCompose() string {
    var path = resolve("/docker-compose.yml")
    return path
}

func GetProject(project string) string {
    var path = resolve("/" + project)
    return path
}

func GetProjectDockerfile(project string) string {
    return GetProject(project) + "/Dockerfile"
}
