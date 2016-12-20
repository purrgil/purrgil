package utils;

import (
    "io/ioutil"
  //  "github.com/ghodss/yaml"
)

func DockerCompose() {
    var path = GetDockerCompose()
    b, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }

    println(b)
}
