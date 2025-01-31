package commands

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/ugabiga/swan/swctl/internal/generate"

	"github.com/spf13/cobra"
)

// validateHandlerName checks if the provided handler name is valid
func validateHandlerName(name string) error {
	// Check if name is empty
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("handler name cannot be empty")
	}

	// Check length
	if len(name) > 50 {
		return fmt.Errorf("handler name is too long (maximum 50 characters)")
	}

	// Check if name follows valid pattern (letters, numbers)
	validName := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*$`)
	if !validName.MatchString(name) {
		return fmt.Errorf("handler name must start with a letter and contain only letters and numbers")
	}

	return nil
}

// validateHandlerPath checks if the provided path is valid
func validateHandlerPath(path string) error {
	if strings.TrimSpace(path) == "" {
		return fmt.Errorf("handler path cannot be empty")
	}

	// Add more path validation if needed
	return nil
}

// validateRoutePrefix checks if the provided route prefix is valid
func validateRoutePrefix(prefix string) error {
	if !strings.HasPrefix(prefix, "/") {
		return fmt.Errorf("route prefix must start with '/'")
	}

	validPrefix := regexp.MustCompile(`^/[a-zA-Z0-9/-]*$`)
	if !validPrefix.MatchString(prefix) {
		return fmt.Errorf("route prefix must contain only letters, numbers, forward slashes and hyphens")
	}

	return nil
}

var MakeHandlerCommand = &cobra.Command{
	Use:   "make:handler [path] [route-prefix] [name]",
	Short: "Create a new handler",
	Args:  cobra.RangeArgs(0, 3),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			handlerPath string
			handlerName string
			routePrefix string
		)

		if len(args) == 3 {
			handlerPath = args[0]
			routePrefix = args[1]
			handlerName = args[2]
		} else {
			if err := huh.NewInput().
				Title("Handler Path (relative to internal/, e.g. 'handlers/user' or 'api/v1/handlers')").
				Value(&handlerPath).
				Run(); err != nil {
				fmt.Println(err)
				return
			}

			if err := huh.NewInput().Title("Route Prefix (eg: /api)").Value(&routePrefix).Run(); err != nil {
				fmt.Println(err)
				return
			}

			if err := huh.NewInput().Title("Handler Name").Value(&handlerName).Run(); err != nil {
				fmt.Println(err)
				return
			}
		}

		// Validate inputs
		if err := validateHandlerPath(handlerPath); err != nil {
			fmt.Printf("Invalid handler path: %s\n", err)
			return
		}

		if err := validateHandlerName(handlerName); err != nil {
			fmt.Printf("Invalid handler name: %s\n", err)
			return
		}

		if err := validateRoutePrefix(routePrefix); err != nil {
			fmt.Printf("Invalid route prefix: %s\n", err)
			return
		}

		if err := generate.CreateHandler(handlerPath, handlerName, routePrefix); err != nil {
			fmt.Printf("Error while creating handler: %s", err)
			return
		}

		fmt.Printf("Handler %s created successfully\n", handlerName)
	},
}
