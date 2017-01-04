package commands

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"log"

	"github.com/google/go-github/github"
)

// Upgrade the current `version` of purrgil to the latest.
func Upgrade() error {
	log.Println("current release is v%s", PurrgilVersion)

	// fetch releases
	gh := github.NewClient(nil)
	releases, _, err := gh.Repositories.ListReleases("purrgil", "purrgil", nil)
	if err != nil {
		return err
	}

	// see if it's new
	latest := releases[0]
	log.Println("latest release is %s", *latest.TagName)

	if (*latest.TagName)[1:] == PurrgilVersion {
		log.Println("you're up to date :)")
		return nil
	}

	asset := findAsset(latest)
	if asset == nil {
		return errors.New("cannot find binary for your system")
	}

	// get the executable's path
	cmdPath, err := exec.LookPath("purrgil")
	if err != nil {
		return err
	}
	cmdDir := filepath.Dir(cmdPath)

	// create tmp file
	tmpPath := filepath.Join(cmdDir, "purrgil-upgrade")
	f, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	// download binary
	log.Println("downloading %s", *asset.BrowserDownloadURL)
	res, err := http.Get(*asset.BrowserDownloadURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// copy it down
	_, err = io.Copy(f, res.Body)
	if err != nil {
		return err
	}

	// replace it
	log.Println("replacing %s", cmdPath)
	err = os.Rename(tmpPath, cmdPath)
	if err != nil {
		return err
	}

	log.Println("visit https://github.com/purrgil/purrgil/releases for the changelog")
	return nil
}

// findAsset returns the binary for this platform.
func findAsset(release *github.RepositoryRelease) *github.ReleaseAsset {
	for _, asset := range release.Assets {
		if *asset.Name == fmt.Sprintf("purrgil_%s_%s", runtime.GOOS, runtime.GOARCH) {
			return &asset
		}
	}
	return nil
}
