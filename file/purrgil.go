package file

type Purrgil struct {
	yaml     Yaml
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

func (p *Purrgil) SaveFile() {
	p.yaml.SaveFile(p)
}

func (p *Purrgil) LoadFile() {
	p.yaml.LoadFile(p)
}

func NewPurrgil(dir string, name string) Purrgil {
	purrgil := Purrgil{}

	purrgil.Name = name

	purrgil.yaml.InitFile(dir, "purrgil")
	purrgil.yaml.LoadFile(&purrgil)

	return purrgil
}
