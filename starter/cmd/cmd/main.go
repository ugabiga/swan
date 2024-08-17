package main

import (
	"github.com/ugabiga/swan/starter/internal/config"
)

func main() {
	if err := config.ProvideApp().RunCmd(); err != nil {
		panic(err)
	}
}
