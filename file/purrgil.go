package file

import (
	"os"

	"github.com/purrgil/file-parser"
	"github.com/purrgil/file-parser/purrgil"
)

// PurrgilFile has a file pointer and a structured data
type PurrgilFile struct {
	File fileparser.File
	Data purrgil.Purrgil
}

// GetPurrgilFile load a purrgil.yml file and your data inside data field
func GetPurrgilFile() (PurrgilFile, error) {
	pwd, _ := os.Getwd()

	file, loadFileErr := fileparser.LoadFile(pwd, "purrgil", "yml")
	if loadFileErr != nil {
		return PurrgilFile{}, loadFileErr
	}

	data, parseErr := purrgil.ParseToPurrgil(file.Content)
	if parseErr != nil {
		return PurrgilFile{}, parseErr
	}

	return PurrgilFile{file, data}, nil
}

// Save convert data to byte and trigger save file
func (dc *PurrgilFile) Save() error {
	content, err := dc.Data.ParseToByte()

	if err != nil {
		return err
	}

	dc.File.Content = content
	return dc.File.SaveFile()
}
