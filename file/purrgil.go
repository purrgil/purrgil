package file

import (
	"github.com/purrgil/file-parser"
	"github.com/purrgil/file-parser/purrgil"
)

type PurrgilFile struct {
	File fileparser.File
	Data purrgil.PurrgilFile
}

func GetPurrgilFile() (PurrgilFile, error) {
	file := fileparser.LoadFile(os.GetWd(), "purrgil", "yml")

	if data, err := ParseToPurrgilFile(file.Content); err != nil {
		return nil, err
	}

	return PurrgilFile{ file, data }, nil
}

func (dc *PurrgilFile) Save() error {
	if content, err := dc.Data.ParseToByte(); err != nil {
		return err
	}

	dc.File.Content = content
	return dc.File.SaveFile()
}
