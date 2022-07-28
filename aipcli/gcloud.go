package aipcli

import (
	"os/exec"
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
