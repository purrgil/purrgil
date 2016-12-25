package file

type Purrgil struct {
	Yaml
	Name     string           `yaml:"name"`
	Packages []PurrgilPackage `yaml:"packages,omitempty"`
	Settings PurrgilSettings  `yaml:"settings,omitempty"`
}


// TODO: Deploy settings and dev envs
type PurrgilSettings struct {
}

func (p *Purrgil) AddPackage(pkg PurrgilPackage) {
	p.Packages = append(p.Packages, pkg)
}

func NewPurrgil(dir string, name string) Purrgil {
	purrgil := Purrgil{}

	purrgil.Name = name

	purrgil.InitFile(dir, "purrgil")
	purrgil.LoadFile()

	return purrgil
}
