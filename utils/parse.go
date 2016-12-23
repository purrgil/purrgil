package utils;

import (
    "io/ioutil"
    "fmt"
    "gopkg.in/yaml.v2"
)

type DockerComposeStruct struct {
    Version string `yaml:"version"`
    Services interface{} `yaml:"services"`
}

// type PurrgilFile struct {
//     version  string  `json:version`
//     services interface{} `json:services`
// }

func GetDockerCompose() {
    var path = GetDockerComposePath()
    yamlFile, err := ioutil.ReadFile(path)

    if err != nil {
        panic(err)
    }

    var config DockerComposeStruct

    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        panic(err)
    }

    return config
}
//
// func PurrgilConfig() {
//     var path = GetDockerCompose()
//     b, err := ioutil.ReadFile(path)
//     if err != nil {
//         panic(err)
//     }
//
//     parsed := yaml.Marshal(string(b))
//
//     println(parsed.version)
// }
