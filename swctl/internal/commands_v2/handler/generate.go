package handler

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"text/template"

	"github.com/ugabiga/swan/swctl/internal/generate"
	"github.com/ugabiga/swan/swctl/internal/utils"
)

//go:embed *.tmpl
var templateFS embed.FS

func Generate(routePrefix, handlerPath, handlerName string) error {
	handlerPath = filepath.Join("internal", handlerPath)
	packageName := utils.ExtractPackageName(handlerPath)
	funcName := "NewHandler"

	if err := utils.IfFolderNotExistsCreate(handlerPath); err != nil {
		return err
	}

	if err := generateHandlerFile(handlerPath, handlerName, routePrefix, packageName); err != nil {
		return err
	}

	if err := generate.RegisterStructToContainer(handlerPath, packageName, funcName); err != nil {
		return err
	}

	if err := registerHandlerToRoute(handlerPath, routePrefix, handlerName); err != nil {
		return err
	}

	return nil
}

func generateHandlerFile(
	handlerPath,
	handlerName,
	routePrefix,
	packageName string,
) error {
	filePath := filepath.Join(handlerPath, "handler.go")

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl, err := template.ParseFS(templateFS, "template.tmpl")
	if err != nil {
		return err
	}

	if err := tmpl.Execute(f, map[string]string{
		"PackageName": packageName,
		"HandlerName": handlerName,
		"RoutePrefix": routePrefix,
	}); err != nil {
		return err
	}

	return nil
}

func registerHandlerToRoute(handlerPath, routePrefix, handlerName string) error {
	routerFile := generate.RouterPath
	routerInvokeFunc := "SetRouter"

	bytes, err := os.ReadFile(routerFile)
	if err != nil {
		return err
	}
	appFileContents := string(bytes)

	moduleName := utils.RetrieveModuleName()

	// Format the package path
	packagePath := moduleName + "/" + handlerPath
	log.Printf("packagePath: %s", packagePath)
	appFileContents = strings.ReplaceAll(appFileContents, "import (", "import (\n\t\""+packagePath+"\"")

	if strings.Contains(appFileContents, routerInvokeFunc+"(") {
		appFileContents = strings.Replace(appFileContents,
			routerInvokeFunc+"(",
			routerInvokeFunc+"(\n\t"+handlerName+"Handler *"+handlerName+".Handler,",
			1,
		)

		if err = os.WriteFile(routerFile, []byte(appFileContents), 0644); err != nil {
			return err
		}
	}

	// Extract version from route prefix if it exists
	extractedVersion := ""
	if strings.Contains(routePrefix, "/v") {
		parts := strings.Split(routePrefix, "/v")
		if len(parts) > 1 {
			extractedVersion = parts[1]
		}
	}

	// Set up API versioning
	groupName := "api"
	if extractedVersion != "" {
		if err := setupVersionedGroup(extractedVersion, &appFileContents); err != nil {
			return err
		}
		groupName = "v" + extractedVersion
	}

	// Add handler routes
	if err := addHandlerRoutes(handlerName, groupName, &appFileContents); err != nil {
		return err
	}

	return os.WriteFile(routerFile, []byte(appFileContents), 0644)
}

func setupVersionedGroup(version string, contents *string) error {
	if version == "1" {
		// For v1, replace placeholder with named group
		*contents = strings.Replace(*contents,
			"_ = api.Group(\"/v1\")",
			"v1 := api.Group(\"/v1\")",
			1,
		)
		return nil
	}

	// For other versions, add new version group if not exists
	versionGroup := fmt.Sprintf("v%s := api.Group(\"/v%s\")", version, version)
	if !strings.Contains(*contents, versionGroup) {
		*contents = strings.Replace(*contents,
			"api := e.Group(\"/api\")",
			fmt.Sprintf("api := e.Group(\"/api\")\n\t%s", versionGroup),
			1,
		)
	}
	return nil
}

func addHandlerRoutes(handlerName, groupName string, contents *string) error {
	if strings.Contains(*contents, "}") {
		*contents = strings.Replace(*contents,
			"}",
			fmt.Sprintf("\t%sHandler.SetRoutes(%s)\n}", handlerName, groupName),
			1,
		)
	}
	return nil
}
