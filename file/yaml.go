package file

import (
	"gopkg.in/yaml.v2"
)

type Yaml struct {
	File
}


func (y *Yaml) UpdateContent() []byte {
	yamlByte, err := yaml.Marshal(y)

	if err != nil {
		panic(err)
	}

	return yamlByte
}

func (y *Yaml) LoadFile() {
	y.File.LoadFile()
	err := yaml.Unmarshal(y.content, y)

	if err != nil {
		panic(err)
	}
}

func (y *Yaml) InitFile(dir string, title string) {
	y.File.InitFile(dir, title, "yml")
}
