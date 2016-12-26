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

func (p *Purrgil) RemoveFromPackages(pkgName string) {
	filteredPkgs := []PurrgilPackage{}

	for _, val := range p.Packages {
		if val.Name != pkgName {
			filteredPkgs = append(filteredPkgs, val)
		}
	}

	p.Packages = filteredPkgs
}

func NewPurrgil(dir string, name string) Purrgil {
	purrgil := Purrgil{}

	purrgil.Name = name

	purrgil.yaml.InitFile(dir, "purrgil")
	purrgil.yaml.LoadFile(&purrgil)

	return purrgil
}

func PurrgilGetPackages(pkgs []PurrgilPackage, filter func(PurrgilPackage) bool) []PurrgilPackage {
	filteredList := []PurrgilPackage{}

	for _, value := range pkgs {
		if filter(value) {
			filteredList = append(filteredList, value)
		}
	}

	return filteredList
}
