package ishell

import (
	"github.com/purrgil/termput"
	"github.com/purrgil/purrgil/file"
)

func CollectDockerServiceInfo(pkg file.PurrgilPackage) (string, file.DockerComposeService) {
	var dcs = file.DockerComposeService{}
	var serviceName =  termput.Input(termput.InputConfig{
		DefaultValue: pkg.Name,
		Label: "What's the name of your service - default:" + pkg.Name,
	})

	dcs.Command = termput.Input(termput.InputConfig{
		Label: "Write container command",
	})

	dcs.Ports = termput.InputLoop(termput.InputConfig{
		Label: "Tell the ports that you want to export, ex. '8000:8000 (enter to skip)'",
	})

	dcs.Volumes = termput.InputLoop(termput.InputConfig{
		Label: "Tell the volumes that you want, ex. '.:/code' \nPS: By default the your code folder are a volume (./" + pkg.Name + ":/app)  (enter to skip)",
	})

	dcs.DependsOn = termput.InputLoop(termput.InputConfig{
		Label: "Your image depends of another service, ex. 'db' (enter to skip)",
	})

	dcs.Links = termput.InputLoop(termput.InputConfig{
		Label: "Your image links with of another?, ex. 'db' (enter to skip)",
	})

	if pkg.Provider == "github" {
		dcs.Build = "./" + pkg.Name + ""
		dcs.Volumes = append(dcs.Volumes, "./"+pkg.Name+":/app")
	} else {
		dcs.Image = pkg.Identity
	}

	return serviceName, dcs
}

func CollectLinkPossibility(pkg file.PurrgilPackage) []string {
	println("You add a Non Service Package, did this will be used as a Volume? If yes, did you want that we add this in dockercompose for you?")
	return termput.InputLoop(termput.InputConfig{
		Label: "Digit the name of service that you want to add a volume (enter to skip)",
	})
}
