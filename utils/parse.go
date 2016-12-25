package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ParseYamlFile(filePath string, file Filer) {
	yamlFile, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, file)

	if err != nil {
		panic(err)
	}
}

func SaveYamlFile(filePath string, file Filer) {
	yamlByte, err := yaml.Marshal(file)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filePath, yamlByte, 0644)

	if err != nil {
		panic(err)
	}
}
