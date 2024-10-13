package command

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ugabiga/swan/cli/internal/generate"
	"github.com/ugabiga/swan/cli/internal/tpl"
	"github.com/ugabiga/swan/cli/internal/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TplData struct {
	PackageName      string
	StructOrFuncName string
	StructType       string
	HandlerURL       string
}

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new handler/command",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var CreateCmdCommand = &cobra.Command{
	Use:   "command",
	Short: "create a new command",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		genType := cmd.Use
		tplData := TplData{}
		if err := gen(tplData, genType, path); err != nil {
			fmt.Printf("Error while creating struct: %s", err)
			return
		}
	},
}

var CreateCmdHandler = &cobra.Command{
	Use:   "handler",
	Short: "create a new handler",
	Long:  "create handler users /api",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("args: %v", args)
		path := args[0]
		genType := cmd.Use
		tplData := TplData{
			HandlerURL: args[1] + "/" + path,
		}
		if err := gen(tplData, genType, path); err != nil {
			fmt.Printf("Error while creating struct: %s", err)
			return
		}
	},
}

func CreateCmds() *cobra.Command {
	CreateCmd.AddCommand(CreateCmdCommand)
	CreateCmd.AddCommand(CreateCmdHandler)
	return CreateCmd
}

func gen(tplData TplData, genType, path string) error {
	filePath := fmt.Sprintf("internal/%s/", path)
	tplData.StructOrFuncName = structOrFuncNameByType(genType)
	tplData.PackageName = extractPackageName(filePath)

	if err := createTemplate(filePath, genType, tplData); err != nil {
		return err
	}

	if err := registerInvokerOrProvider(generate.CommandPath, tplData.PackageName, tplData.StructOrFuncName); err != nil {
		return err
	}

	log.Printf("Generated %s%s", filePath, strings.ToLower(genType)+".go")

	return nil
}

func createTemplate(filePath, generationType string, tplData TplData) error {
	f := createFile(filePath, strings.ToLower(generationType)+".go")
	if f == nil {
		log.Printf("warn: file %s%s %s", filePath, strings.ToLower(generationType)+".go", "already exists.")
		return nil
	}
	defer f.Close()

	t, err := template.ParseFS(tpl.CreateTemplateFS, "create/"+generationType+".tmpl")
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

func readFileContent(filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	fileContent := string(bytes)

	return fileContent, nil
}

func extractPackageName(path string) string {
	packageName := filepath.Base(path)
	packageName = strings.TrimSuffix(packageName, "/")

	return packageName
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

func invokerPathByType(genType string) string {
	switch genType {
	case "command":
		return generate.CommandPath
	default:
		return generate.AppPath
	}
}

func registerFuncByType(genType string) string {
	switch genType {
	case "command":
		return "app.RegisterInvokers("
	default:
		return "app.RegisterProviders("
	}
}

func structOrFuncNameByType(genType string) string {
	switch genType {
	case "command":
		return "SetCommands"
	case "handler":
		return "NewHandler"
	default:
		return ""
	}
}
