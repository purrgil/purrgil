package commands

import (
	"os/exec"
	"errors"
	"github.com/purrgil/purrgil/interactiveshell"
)

func Push(commitMessage string) {
	var err error

	ishell.PurrgilAlert("Starting git commit flow :) Wait a sec...")

	gitAdd := exec.Command("git", "add", "docker-compose.yml", "purrgil.yml", ".gitignore")
	gitCommit := exec.Command("git", "commit", "-m", commitMessage)
	gitPush := exec.Command("git", "push", "origin", "master")

	err = gitAdd.Run();

	if err != nil {
		addErr := "Error on add files"
		ishell.Err(addErr, errors.New(addErr))
	}

	err = gitCommit.Run();

	if err != nil {
		commitErr := "Error on commit command"
		ishell.Err(commitErr, errors.New(commitErr))
	}

	err = gitPush.Run();

	if err != nil {
		pushError := "Error on push command"
		ishell.Err(pushError, errors.New(pushError))
	}

	ishell.PurrgilAlert("Commit completed! Check your repo :D")
}
