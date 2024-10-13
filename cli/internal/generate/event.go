package generate

import (
	"github.com/ugabiga/swan/cli/internal/utils"
	"log"
	"os"
)

func CreateEvent(path string) error {
	folderPath := "internal/" + path

	if err := utils.IfFolderNotExistsCreate(folderPath); err != nil {
		return err
	}

	if err := createEvent(folderPath); err != nil {
		return err
	}

	return nil
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

	if err := registerToInvoker(EventPath,
		fullPackageName, packageName, funcName); err != nil {
		log.Printf("Error while register struct %s", err)
		return err
	}

	return nil
}
