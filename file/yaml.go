package file

import (
	"gopkg.in/yaml.v2"
)

type Yaml struct {
	File
}

func (y *Yaml) SaveFile(f Filer) {
	yamlByte, err := yaml.Marshal(f)

	if err != nil {
		panic(err)
	}

	y.File.content = yamlByte
	y.File.SaveFile()
}

func (y *Yaml) LoadFile(f Filer) {
	y.File.LoadFile()
	err := yaml.Unmarshal(y.File.content, f)

	if err != nil {
		panic(err)
	}
}

func (y *Yaml) InitFile(dir string, title string) {
	y.File.InitFile(dir, title, "yml")
}
