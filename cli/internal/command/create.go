package command

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/cli/internal/tpl"
	"github.com/ugabiga/swan/cli/internal/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TplData struct {
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
		tplData := TplData{
			HandlerURL: args[1] + "/" + path,
		}

		initTplData(&tplData, genType, path)

		if err := gen(genType, tplData, path); err != nil {
			log.Printf("Error while creating: %s", err)
			return
		}

		if err := registerRouter(tplData); err != nil {
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
		tplData := TplData{}
		if err := gen(genType, tplData, path); err != nil {
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
		tplData := TplData{}
		if err := gen(genType, tplData, path); err != nil {
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
		tplData := TplData{
			StructName: args[1],
		}
		if err := gen(genType, tplData, path); err != nil {
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

func funcNameByType(genType string, tplData TplData) string {
	switch genType {
	case "command":
		return "SetCommands"
	case "event":
		return "SetEvents"
	case "handler":
		return "NewHandler"
	case "struct":
		return "New" + tplData.StructName
	default:
		return ""
	}
}

func fileNameByType(genType string, tplData TplData) string {
	switch genType {
	case "struct":
		return strings.ToLower(tplData.StructName) + ".go"
	default:
		return strings.ToLower(genType) + ".go"
	}
}

func initTplData(tplData *TplData, genType, path string) {
	filePath := fmt.Sprintf("internal/%s/", path)
	tplData.FuncName = funcNameByType(genType, *tplData)
	tplData.PackageName = extractPackageName(filePath)
}

func gen(genType string, tplData TplData, path string) error {
	filePath := fmt.Sprintf("internal/%s/", path)
	tplData.FuncName = funcNameByType(genType, tplData)
	tplData.PackageName = extractPackageName(filePath)

	if err := createTemplate(filePath, genType, tplData); err != nil {
		return err
	}

	if err := registerInvokerOrProvider(genType, tplData.PackageName, tplData.FuncName); err != nil {
		return err
	}

	log.Printf("Generated %s%s", filePath, fileNameByType(genType, tplData))

	return nil
}

func createTemplate(filePath, genType string, tplData TplData) error {
	f := createFile(filePath, fileNameByType(genType, tplData))
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

	if err = t.Execute(f, tplData); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
		return err
	}

	return nil
}

func registerRouter(tplData TplData) error {
	projectPackage := utils.RetrieveModuleName()
	registerFunc := "SetRouteHTTPServer("
	routerPath := RoutePath
	handlerName := tplData.PackageName + "Handler"
	invokerStruct := tplData.PackageName + "." + "Handler"
	invokerFullPackage := projectPackage + "/internal/" + tplData.PackageName

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
