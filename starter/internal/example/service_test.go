package example_test

import (
	"testing"

	"github.com/ugabiga/swan/starter/internal/providers"
)

func TestCreate(t *testing.T) {
	service := providers.ProvideTestApp(t).Deps.ExampleService

	t.Run("should return created", func(t *testing.T) {
		result := service.Create()
		if result != "created" {
			t.Error("Expected created, got", result)
		}
	})
}
