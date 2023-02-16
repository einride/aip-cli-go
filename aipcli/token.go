package aipcli

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func gcloudAuthPrintIdentityToken() (string, bool) {
	if _, err := exec.LookPath("gcloud"); err != nil {
		return "", false
	}
	var stdout strings.Builder
	cmd := exec.Command("gcloud", "auth", "print-identity-token")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", false
	}
	return strings.TrimSpace(stdout.String()), true
}

type IdentityTokenFile struct {
	IdentityToken string
}

func identityTokenFromConfigFile(tokenFile string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configFile := path.Join(homeDir, ".config", "saga", tokenFile)
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return "", err
	}
	var identityFile IdentityTokenFile
	err = json.Unmarshal(file, &identityFile)
	if err != nil {
		return "", err
	}
	return identityFile.IdentityToken, nil
}
