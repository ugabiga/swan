package main

import "github.com/ugabiga/swan/starter/internal/providers"

func main() {
	if err := providers.ProvideAppAndRun(); err != nil {
		panic(err)
	}
}
