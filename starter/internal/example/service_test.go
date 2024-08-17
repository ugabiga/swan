package example_test

import (
	"testing"

	"github.com/ugabiga/swan/starter/internal/config"
)

func TestCreate(t *testing.T) {
	service := config.ProvideTestApp(t).Deps.ExampleService

	t.Run("should return created", func(t *testing.T) {
		result := service.Create()
		if result != "created" {
			t.Error("Expected created, got", result)
		}
	})
}
