package generating

import (
	"log"
	"os"
	"strings"

	"github.com/ugabiga/swan/cli/internal/utils"
)

func CreateHandler(
	folderPath string,
	domainName string,
	routePrefix string,
) error {
	template := `package ` + domainName + `

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ugabiga/swan/utl"
)

type Handler struct {
	prefix string
	logger *slog.Logger
}

func NewHandler(
	logger *slog.Logger,
) *Handler {
	return &Handler{
		prefix: "/DOMAIN",
		logger: logger,
	}
}

func (h *Handler) SetRoutes(e *echo.Group) {
	e.GET(h.prefix, h.List)
	e.GET(h.prefix+"/:id", h.One)
	e.POST(h.prefix, h.Create)
	e.PUT(h.prefix+"/:id", h.Edit)
	e.DELETE(h.prefix+"/:id", h.Delete)
}

// List godoc
//
//	@Summary		List
//	@Description	List
//	@Tags			DOMAIN
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		int		false	"Limit"
//	@Param			offset		query		int		false	"Offset"
//	@Success		200			{object}	ListResp
//	@Failure		400			{object}	utl.RequestValidationErrorResponse
//	@Router			ROUTE_PREFIX/DOMAIN [get]
func (h *Handler) List(c echo.Context) error {
	var query ListQuery
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, utl.ValidateRequestAllEmptyResponse())
	}

	validationErrors, ok := utl.ValidateRequestStruct(query)
	if !ok {
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	return c.JSON(http.StatusOK, ListResp{
		Tag: "list",
	})
}

type ListQuery struct {
	Limit  int ` + "`query:\"limit\"`" + `
	Offset int ` + "`query:\"offset\"`" + `
}

type ListResp struct {
	Tag string ` + "`json:\"func\"`" + `
}

// One godoc
//
//	@Summary		One
//	@Description	One
//	@Tags			DOMAIN
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string	true	"ID"
//	@Param			limit		query		int		false	"Limit"
//	@Param			offset		query		int		false	"Offset"
//	@Success		200			{object}	ListResp
//	@Failure		400			{object}	utl.RequestValidationErrorResponse
//	@Router			ROUTE_PREFIX/DOMAIN/{id} [get]
func (h *Handler) One(c echo.Context) error {
	id := c.Param("id")

	var query OneQuery
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, utl.ValidateRequestAllEmptyResponse())
	}

	validationErrors, ok := utl.ValidateRequestStruct(query)
	if !ok {
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	return c.JSON(http.StatusOK, OneResp{
		ID: id,
	})
}

type OneQuery struct {
	Limit  int ` + "`query:\"limit\"`" + `
	Offset int ` + "`query:\"offset\"`" + `
}

type OneResp struct {
	ID string ` + "`json:\"id\"`" + `
}

// Create godoc
//
//	@Summary		Create
//	@Description	Create
//	@Tags			DOMAIN
//	@Accept			json
//	@Produce		json
//	@Param			req	body		CreateReq	true	"Request"
//	@Success		200	{object}	CreateResp
//	@Failure		400	{object}	utl.RequestValidationErrorResponse
//	@Router			ROUTE_PREFIX/DOMAIN [post]
func (h *Handler) Create(c echo.Context) error {
	var req CreateReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utl.ValidateRequestAllEmptyResponse())
	}

	validationErrors, ok := utl.ValidateRequestStruct(req)
	if !ok {
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	return c.JSON(http.StatusOK, CreateResp{
		Name: req.Name,
	})
}

type CreateReq struct {
	Name string ` + "`json:\"name\"`" + `
}

type CreateResp struct {
	Name string ` + "`json:\"name\"`" + `
}

// Edit godoc
//
//	@Summary		Edit
//	@Description	Edit
//	@Tags			DOMAIN
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID"
//	@Param			req	body		EditReq	true	"Request"
//	@Success		200	{object}	EditResp
//	@Failure		400	{object}	utl.RequestValidationErrorResponse
//	@Router			ROUTE_PREFIX/DOMAIN/{id} [put]
func (h *Handler) Edit(c echo.Context) error {
	id := c.Param("id")

	var req EditReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utl.ValidateRequestAllEmptyResponse())
	}

	validationErrors, ok := utl.ValidateRequestStruct(req)
	if !ok {
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	return c.JSON(http.StatusOK, EditResp{
		ID:   id,
		Name: req.Name,
	})
}

type EditReq struct {
	Name string ` + "`json:\"name\"`" + `
}

type EditResp struct {
	ID   string ` + "`json:\"id\"`" + `
	Name string ` + "`json:\"name\"`" + `
}

// Delete godoc
//
//	@Summary		Delete
//	@Description	Delete
//	@Tags			DOMAIN
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string		true	"ID"
//	@Param			req	body		DeleteReq	true	"Request"
//	@Success		200	{object}	DeleteResp
//	@Failure		400	{object}	utl.RequestValidationErrorResponse
//	@Router			ROUTE_PREFIX/DOMAIN/{id} [delete]
func (h *Handler) Delete(c echo.Context) error {
	id := c.Param("id")

	var req DeleteReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utl.ValidateRequestAllEmptyResponse())
	}

	validationErrors, ok := utl.ValidateRequestStruct(req)
	if !ok {
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	return c.JSON(http.StatusOK, DeleteResp{
		ID:   id,
		Name: req.Name,
	})
}

type DeleteReq struct {
	Name string ` + "`json:\"name\"`" + `
}

type DeleteResp struct {
	ID   string ` + "`json:\"id\"`" + `
	Name string ` + "`json:\"name\"`" + `
}
`
	filePath := folderPath + "/handler.go"

	// replace DOMAIN with domainName
	template = strings.ReplaceAll(template, "DOMAIN", domainName)
	template = strings.ReplaceAll(template, "ROUTE_PREFIX", routePrefix)

	// write to file
	if err := os.WriteFile(filePath, []byte(template), 0644); err != nil {
		return err
	}

	if err := registerHandlerToApp(domainName); err != nil {
		return err
	}

	if err := registerHandlerToRoute(domainName); err != nil {
		return err
	}

	return nil
}

func registerStructToApp(fullPackageName, packageName, structName string) error {
	appFilePath := "./internal/config/app.go"
	appRegisterProvidersFunc := "app.RegisterProviders"

	bytes, err := os.ReadFile(appFilePath)
	if err != nil {
		return err
	}
	appFileContents := string(bytes)

	moduleName := utils.RetrieveModuleName()

	// Format the package path
	var packagePath string
	packagePath = moduleName + "/" + fullPackageName
	appFileContents = strings.ReplaceAll(appFileContents, "import (", "import (\n\t\""+packagePath+"\"")

	if strings.Contains(appFileContents, appRegisterProvidersFunc+"(") {
		appFileContents = strings.Replace(appFileContents, appRegisterProvidersFunc+
			"(", appRegisterProvidersFunc+
			"(\n\t\t"+packageName+".New"+structName+",", 1)

		if err = os.WriteFile(appFilePath, []byte(appFileContents), 0644); err != nil {
			return err
		}
	}

	return nil
}

func registerHandlerToApp(domainName string) error {
	appFilePath := "./internal/config/app.go"
	appRegisterProvidersFunc := "app.RegisterProviders"

	bytes, err := os.ReadFile(appFilePath)
	if err != nil {
		return err
	}
	appFileContents := string(bytes)

	moduleName := utils.RetrieveModuleName()

	// Format the package path
	var packagePath string
	packagePath = moduleName + "/internal/" + domainName
	log.Printf("packagePath: %s", packagePath)
	appFileContents = strings.ReplaceAll(appFileContents, "import (", "import (\n\t\""+packagePath+"\"")

	if strings.Contains(appFileContents, appRegisterProvidersFunc+"(") {
		appFileContents = strings.Replace(appFileContents, appRegisterProvidersFunc+
			"(", appRegisterProvidersFunc+
			"(\n\t\t"+domainName+".NewHandler,", 1)

		if err = os.WriteFile(appFilePath, []byte(appFileContents), 0644); err != nil {
			return err
		}
	}

	return nil
}

func registerHandlerToRoute(domainName string) error {
	routerFile := "./internal/config/route.go"
	routerInvokeFunc := "InvokeSetRouteHTTPServer"

	bytes, err := os.ReadFile(routerFile)
	if err != nil {
		return err
	}
	appFileContents := string(bytes)

	moduleName := utils.RetrieveModuleName()

	// Format the package path
	var packagePath string
	packagePath = moduleName + "/internal/" + domainName
	log.Printf("packagePath: %s", packagePath)
	appFileContents = strings.ReplaceAll(appFileContents, "import (", "import (\n\t\""+packagePath+"\"")

	if strings.Contains(appFileContents, routerInvokeFunc+"(") {
		appFileContents = strings.Replace(appFileContents,
			routerInvokeFunc+"(",
			routerInvokeFunc+"(\n\t"+domainName+"Handler *"+domainName+".Handler,",
			1,
		)

		if err = os.WriteFile(routerFile, []byte(appFileContents), 0644); err != nil {
			return err
		}
	}

	if strings.Contains(appFileContents, "}") {
		appFileContents = strings.Replace(appFileContents,
			"}",
			"\t"+domainName+"Handler.SetRoutes(group)\n}",
			1,
		)
		if err = os.WriteFile(routerFile, []byte(appFileContents), 0644); err != nil {
			return err
		}
	}

	return nil
}
