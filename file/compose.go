package file

import (
	"os"

	"github.com/purrgil/file-parser"
	"github.com/purrgil/file-parser/docker/compose"
)

// DockerCompose has a file pointer and a structured data
type DockerCompose struct {
	File fileparser.File
	Data compose.DockerCompose
}

// GetDockerCompose load a docker-compose.yml file and your data inside data field
func GetDockerCompose() (DockerCompose, error) {
	pwd, _ := os.Getwd()

	file, loadFileErr := fileparser.LoadFile(pwd, "docker-compose", "yml")

	if loadFileErr != nil {
		return DockerCompose{}, loadFileErr
	}

	data, parseErr := compose.ParseToDockerCompose(file.Content)

	if parseErr != nil {
		return DockerCompose{}, parseErr
	}

	return DockerCompose{file, data}, nil
}

// Save convert data to byte and trigger save file
func (dc *DockerCompose) Save() error {
	content, err := dc.Data.ParseToByte()

	if err != nil {
		return err
	}

	dc.File.Content = content
	return dc.File.SaveFile()
}
