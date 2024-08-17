package main

import (
	"github.com/ugabiga/swan/starter/internal/config"
)

func main() {
	if err := config.ProvideApp().RunCommand("example"); err != nil {
		panic(err)
	}
}
