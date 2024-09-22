package generating

import (
	"os"
)

func CreateDomain(domainName string, routePrefix string) {
	folderPath := "internal/" + domainName

	if err := os.Mkdir(folderPath, 0755); err != nil {
		panic(err)
	}

	if err := CreateHandler(folderPath, domainName, routePrefix); err != nil {
		panic(err)
	}
}
