package dir

import (
	"os/exec"
	"strings"
)

func ProjectRoot() string {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	out, err := cmd.Output()
	if err != nil {
		// If we can't find project root via go modules, use current directory
		pwd := exec.Command("pwd")
		out, err = pwd.Output()
		if err != nil {
			return "."
		}
		return strings.TrimSpace(string(out))
	}

	return strings.TrimSpace(string(out))
}
