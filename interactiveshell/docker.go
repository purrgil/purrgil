package ishell

import (
  "github.com/guidiego/purrgil/file"
	"github.com/guidiego/gommunicate-shell"
)

func CollectDockerServiceInfo(pkg file.PurrgilPackage) (string, file.DockerComposeService) {
	var dcs = file.DockerComposeService{}
	var serviceName = gshell.Ask("What's the name of your service", pkg.Name)

	dcs.Command = gshell.Ask("Write container command", "")
	dcs.Ports = gshell.AskTilBlankEnter("Tell the ports that you want to export, ex. '8000:8000'", "(enter to skip)")
	dcs.Volumes = gshell.AskTilBlankEnter("Tell the volumes that you want, ex. '.:/code'", "(enter to skip)")
	dcs.DependsOn = gshell.AskTilBlankEnter("Your image depends of another service, ex. 'db'", "(enter to skip)")

	if pkg.Provider == "github" {
		dcs.Build = "./" + pkg.Name + "/Dockerfile"
	} else {
		dcs.Image = pkg.Identity
	}

  dcs.Volumes = append(dcs.Volumes, ".:/" + pkg.Name)

  return serviceName, dcs
}

func CollectLinkPossibility(pkg file.PurrgilPackage) []string {
  return []string{}
}
