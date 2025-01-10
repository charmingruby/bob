package update

import (
	"os/exec"
)

func Perform() error {
	cmd := exec.Command("bash", "-c", "curl -sSL https://raw.githubusercontent.com/charmingruby/bob/main/install.sh | bash")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
