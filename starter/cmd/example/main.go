package main

import (
	"github.com/ugabiga/swan/starter/internal/providers"
)

func main() {
	if err := providers.ProvideApp().RunCommand("example"); err != nil {
		panic(err)
	}
}
