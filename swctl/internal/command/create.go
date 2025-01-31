package command

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/swctl/internal/tpl"
	"github.com/ugabiga/swan/swctl/internal/utils"
)

type GenData struct {
	FilePath    string
	PackageName string
	FuncName    string
	StructType  string
	HandlerURL  string // for handler
	StructName  string
}

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new handler/command",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var CreateCmdHandler = &cobra.Command{
	Use:   "handler",
	Short: "create a new handler",
	Long:  "create handler users /api",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		genType := cmd.Use
		genData := GenData{
			HandlerURL: args[1] + "/" + path,
		}

		initGenData(&genData, genType, path)

		if err := gen(genType, genData); err != nil {
			log.Printf("Error while creating: %s", err)
			return
		}

		if err := registerRouter(genData); err != nil {
			log.Printf("Error while creating: %s", err)
			return
		}
	},
}

var CreateCmdCommand = &cobra.Command{
	Use:   "command",
	Short: "create a new command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		genType := cmd.Use
		genData := GenData{}
		initGenData(&genData, genType, path)
		if err := gen(genType, genData); err != nil {
			log.Printf("Error while creating: %s", err)
			return
		}
	},
}

var CreateCmdEvent = &cobra.Command{
	Use:   "event",
	Short: "create a new event",
	Long:  "Usage : create event users",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		genType := cmd.Use
		genData := GenData{}
		initGenData(&genData, genType, path)
		if err := gen(genType, genData); err != nil {
			log.Printf("Error while creating: %s", err)
			return
		}
	},
}

var CreateCmdStruct = &cobra.Command{
	Use:   "struct",
	Short: "create a new struct",
	Long:  "Usage : create struct users Service",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		genType := cmd.Use
		genData := GenData{
			StructName: args[1],
		}
		initGenData(&genData, genType, path)
		if err := gen(genType, genData); err != nil {
			log.Printf("Error while creating: %s", err)
			return
		}
	},
}

func CreateCmds() *cobra.Command {
	CreateCmd.AddCommand(CreateCmdHandler)
	CreateCmd.AddCommand(CreateCmdCommand)
	CreateCmd.AddCommand(CreateCmdEvent)
	CreateCmd.AddCommand(CreateCmdStruct)
	return CreateCmd
}

func invokerPathByType(genType string) string {
	switch genType {
	case "command":
		return CommandPath
	case "event":
		return EventPath
	default:
		return AppPath
	}
}

func registerFuncByType(genType string) string {
	switch genType {
	case "command", "event":
		return "app.RegisterInvokers("
	default:
		return "app.RegisterProviders("
	}
}

func funcNameByType(genType string, genData GenData) string {
	switch genType {
	case "command":
		return "SetCommands"
	case "event":
		return "SetEvents"
	case "handler":
		return "NewHandler"
	case "struct":
		return "New" + genData.StructName
	default:
		return ""
	}
}

func fileNameByType(genType string, genData GenData) string {
	switch genType {
	case "struct":
		return strings.ToLower(genData.StructName) + ".go"
	default:
		return strings.ToLower(genType) + ".go"
	}
}

func initGenData(genData *GenData, genType, path string) {
	filePath := fmt.Sprintf("internal/%s/", path)

	genData.FilePath = filePath
	genData.FuncName = funcNameByType(genType, *genData)
	genData.PackageName = extractPackageName(filePath)
}

func gen(genType string, genData GenData) error {
	if err := createTemplate(genData.FilePath, genType, genData); err != nil {
		return err
	}

	if err := registerInvokerOrProvider(genType, genData.PackageName, genData.FuncName); err != nil {
		return err
	}

	log.Printf("Generated %s%s", genData.FilePath, fileNameByType(genType, genData))

	return nil
}

func createTemplate(filePath, genType string, genData GenData) error {
	f := createFile(filePath, fileNameByType(genType, genData))
	if f == nil {
		log.Printf("warn: file %s%s %s", filePath, strings.ToLower(genType)+".go", "already exists.")
		return nil
	}
	defer f.Close()

	t, err := template.ParseFS(tpl.CreateTemplateFS, "create/"+genType+".tmpl")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
		return err
	}

	if err = t.Execute(f, genData); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
		return err
	}

	return nil
}

func registerRouter(genData GenData) error {
	projectPackage := utils.RetrieveModuleName()
	registerFunc := "SetRouteHTTPServer("
	routerPath := RoutePath
	handlerName := genData.PackageName + "Handler"
	invokerStruct := genData.PackageName + "." + "Handler"
	invokerFullPackage := projectPackage + "/internal/" + genData.PackageName

	fileContent, err := readFileContent(routerPath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", routerPath, err)
		return err
	}

	if !strings.Contains(fileContent, registerFunc) {
		return errors.New("app.RegisterInvokers not found in the file")
	}

	if !strings.Contains(fileContent, invokerFullPackage) {
		// Replace package name
		fileContent = strings.ReplaceAll(
			fileContent,
			"import (", "import (\n\t\""+invokerFullPackage+"\"",
		)
	}

	if strings.Contains(fileContent, "*"+invokerStruct) {
		log.Printf("Already registered %s in the file %s", invokerStruct, routerPath)
		return nil
	}

	fileContent = strings.Replace(
		fileContent,
		registerFunc,
		registerFunc+"\n\t"+handlerName+" *"+invokerStruct+",",
		1,
	)

	if strings.Contains(fileContent, "}") {
		fileContent = strings.Replace(fileContent,
			"}",
			"\t"+handlerName+".SetRoutes(g)\n}",
			1,
		)
	}

	if err = os.WriteFile(routerPath, []byte(fileContent), 0644); err != nil {
		log.Fatalf("Failed to write file %s: %v", routerPath, err)
		return err
	}

	return nil
}

func registerInvokerOrProvider(genType, packageName, structName string) error {
	projectPackage := utils.RetrieveModuleName()
	registerFunc := registerFuncByType(genType)
	invokerPath := invokerPathByType(genType)
	invokerFunc := packageName + "." + structName
	invokerFullPackage := projectPackage + "/internal/" + packageName

	invokerContent, err := readFileContent(invokerPath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", invokerPath, err)
		return err
	}

	if !strings.Contains(invokerContent, registerFunc) {
		return errors.New("app.RegisterInvokers not found in the file")
	}

	if strings.Contains(invokerContent, invokerFunc) {
		log.Printf("Already registered %s in the file %s", invokerFunc, invokerPath)
		return nil
	}

	if !strings.Contains(invokerContent, invokerFullPackage) {
		// Replace package name
		invokerContent = strings.ReplaceAll(
			invokerContent,
			"import (", "import (\n\t\""+invokerFullPackage+"\"",
		)
	}

	invokerContent = strings.Replace(
		invokerContent,
		registerFunc,
		registerFunc+"\n\t\t"+invokerFunc+",",
		1,
	)

	if err = os.WriteFile(invokerPath, []byte(invokerContent), 0644); err != nil {
		log.Fatalf("Failed to write file %s: %v", invokerPath, err)
		return err
	}

	return nil
}

func extractPackageName(path string) string {
	packageName := filepath.Base(path)
	packageName = strings.TrimSuffix(packageName, "/")

	return packageName
}

func readFileContent(filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	fileContent := string(bytes)

	return fileContent, nil
}

func createFile(dirPath string, filename string) *os.File {
	filePath := filepath.Join(dirPath, filename)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create dir %s: %v", dirPath, err)
	}
	stat, _ := os.Stat(filePath)
	if stat != nil {
		return nil
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filePath, err)
	}

	return file
}
