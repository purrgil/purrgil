package file

import (
  "io/ioutil"
)

type Filer interface {
  LoadFile()
  SaveFile()
  InitFile(string, string, string)
  GetFilePath() string
  UpdateContent() []byte
}

type File struct {
  content   []byte
  title     string
  dir       string
  extension string
}

func (f *File) UpdateContent() []byte {
  return f.content
}

func (f *File) GetFilePath() string {
  return f.dir + "/" + f.title + "." + f.extension
}

func (f *File) SaveFile() {
  f.content = f.UpdateContent()

  path := f.GetFilePath()
  err  := ioutil.WriteFile(path, f.content, 0777)

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
