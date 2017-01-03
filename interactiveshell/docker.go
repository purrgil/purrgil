package ishell

import (
	"github.com/purrgil/termput"
	"github.com/purrgil/purrgil/file"
)

func CollectDockerServiceInfo(pkg file.PurrgilPackage) (string, file.DockerComposeService) {
	var dcs = file.DockerComposeService{}
	var serviceName = termput.Input("What's the name of your service", pkg.Name)

	dcs.Command = termput.Input("Write container command", "")
	dcs.Ports = termput.LoopInput("Tell the ports that you want to export, ex. '8000:8000'", "(enter to skip)")
	dcs.Volumes = termput.LoopInput("Tell the volumes that you want, ex. '.:/code' \nPS: By default the your code folder are a volume (./" + pkg.Name + ":/app)" , "(enter to skip)")
	dcs.DependsOn = termput.LoopInput("Your image depends of another service, ex. 'db'", "(enter to skip)")

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
	return termput.LoopInput("Digit the name of service that you want to add a volume", "(enter to skip)")
}
