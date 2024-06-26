package aipcli

import (
	"encoding/json"
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
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	configFile := path.Join(configDir, tokenFile)
	file, err := os.ReadFile(configFile)
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
