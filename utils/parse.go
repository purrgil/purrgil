package utils;

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "github.com/guidiego/purrgil/models"
)

func parseYamlFile(filePath string, dockercompose *models.DockerComposeFile) {
    yamlFile, err := ioutil.ReadFile(filePath)

    if err != nil {
        panic(err)
    }

    err = yaml.Unmarshal(yamlFile, dockercompose)

    if err != nil {
        panic(err)
    }
}

func GetDockerCompose() models.DockerComposeFile {
    var dockercompose models.DockerComposeFile

    dockerComposePath := Path.DockerCompose()
    parseYamlFile(dockerComposePath, &dockercompose)

    return dockercompose;
}
