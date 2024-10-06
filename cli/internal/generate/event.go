package generate

import (
	"log"
	"os"
)

func CreateEvent(path string) {
	folderPath := "internal/" + path

	//Check if folder exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			panic(err)
		}
	}

	if err := createEvent(folderPath); err != nil {
		panic(err)
	}
}

func createEvent(folderPath string) error {
	fileName := "event"
	filePath := folderPath + "/" + fileName + ".go"
	fullPackageName := folderPath
	packageName := extractPackageName(folderPath)
	funcName := "InvokeSetEventRouter"

	template := `package ` + packageName + `

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ugabiga/swan/core"
	"log/slog"
)

func InvokeSetEventRouter(
	logger *slog.Logger,
	eventEmitter *core.EventEmitter,
) {
	eventEmitter.AddOneWayHandler(
		"eventHandler",
		"event",
		func(msg *message.Message) error {
			logger.Info("Received message",
				slog.Any("uuid", msg.UUID),
				slog.String("payload", string(msg.Payload)),
			)

			return nil
		},
	)
}
`
	if err := os.WriteFile(filePath, []byte(template), 0644); err != nil {
		log.Printf("Error while creating struct: %s", err)
		return err
	}

	if err := registerToInvoker("./internal/config/event.go",
		fullPackageName, packageName, funcName); err != nil {
		log.Printf("Error while register struct %s", err)
		return err
	}

	return nil
}
