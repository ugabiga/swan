package core

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func ValidateEnvironmentVariables[T interface{}]() *T {
	var env *T = new(T)

	// Fill env with the actual  environment variables
	valueType := reflect.TypeOf(env).Elem()
	value := reflect.ValueOf(env).Elem()

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		envName := field.Tag.Get("env")

		if envName == "" {
			envName = field.Name
		}

		if envValue := os.Getenv(envName); envValue != "" {
			value.Field(i).SetString(envValue)
		}
	}

	// Validate the environment variables
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(env); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)

		for _, validationError := range validationErrors {
			slog.Error(
				"Invalid environment variable",
				"field", validationError.Field(),
				"value", validationError.Value(),
				"tag", validationError.Tag(),
			)
		}

		log.Fatal("Failed to validate environment variables")
	}

	return env
}

func LoadEnv(envFile string) {
	err := godotenv.Load(rootPath(envFile))
	if err != nil {
		panic(fmt.Errorf("error loading .env file: %w", err))
	}
}

func rootPath(file string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			panic(fmt.Errorf("go.mod not found"))
		}
		currentDir = parent
	}

	return filepath.Join(currentDir, file)
}
