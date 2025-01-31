package commands

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/swctl/internal/generate"
)

// validateAppName checks if the provided name is valid for a new app
func validateAppName(name string) error {
	// Check if name is empty
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("app name cannot be empty")
	}

	// Check length
	if len(name) > 50 {
		return fmt.Errorf("app name is too long (maximum 50 characters)")
	}

	// Check if name follows valid pattern (lowercase letters, numbers, and hyphens)
	validName := regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)
	if !validName.MatchString(name) {
		return fmt.Errorf("app name must contain only lowercase letters, numbers, and hyphens, and cannot start or end with a hyphen")
	}

	return nil
}

var NewCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new app",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		if err := validateAppName(name); err != nil {
			fmt.Printf("Invalid app name: %s\n", err)
			return
		}

		if err := generate.CreateNew(name); err != nil {
			fmt.Printf("Error while creating app: %s", err)
			return
		}

		fmt.Printf("New App %s created successfully\n", name)
	},
}
