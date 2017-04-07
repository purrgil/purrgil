package file

import (
	"os"
	"github.com/purrgil/file-parser"
	"github.com/purrgil/file-parser/docker/compose"
)

type DockerCompose struct {
	File fileparser.File
	Data compose.DockerCompose
}

func GetDockerCompose() (DockerCompose, error) {
	file := fileparser.LoadFile(os.GetWd(), "docker-compose", "yml")

	if data, err := ParseToDockerCompose(file.Content); err != nil {
		return nil, err
	}

	return DockerCompose{ file, data }, nil
}

func (dc *DockerCompose) Save() error {
	if content, err := dc.Data.ParseToByte(); err != nil {
		return err
	}

	dc.File.Content = content
	return dc.File.SaveFile()
}
