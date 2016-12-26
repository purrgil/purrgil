package file

import (
	"io/ioutil"
)

type Filer interface {
	LoadFile()
	SaveFile()
}

type File struct {
	content   []byte
	title     string
	dir       string
	extension string
}

func (f *File) GetFilePath() string {
	basePath := f.dir + "/" + f.title

	if f.extension != "" {
		return basePath + "." + f.extension
	}

	return basePath
}

func (f *File) SaveFile() {
	path := f.GetFilePath()
	err := ioutil.WriteFile(path, f.content, 0777)

	if err != nil {
		panic(err)
	}
}

func (f *File) LoadFile() {
	path := f.GetFilePath()
	content, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	f.content = content
}

func (f *File) InitFile(dir string, title string, extension string) {
	f.dir = dir
	f.title = title
	f.extension = extension
	f.content = []byte{}
}
