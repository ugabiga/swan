package generate

import (
	"os"
	"strings"

	"github.com/ugabiga/swan/swctl/internal/utils"
)

func CreateHandler(path, handlerName string, routePrefix string) error {
	folderPath := "internal/" + path

	if err := utils.IfFolderNotExistsCreate(folderPath); err != nil {
		return err
	}

	if err := createHandler(folderPath, handlerName, routePrefix); err != nil {
		return err
	}

	return nil
}

func createHandler(
	folderPath string,
	handlerName string,
	routePrefix string,
) error {
	packageName := extractPackageName(folderPath)

	template := `package ` + packageName + `

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	prefix string
	logger *slog.Logger
}

func NewHandler(
	logger *slog.Logger,
) *Handler {
	return &Handler{
		prefix: "/HANDLER_NAME",
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
//	@Tags			HANDLER_NAME
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		int		false	"Limit"
//	@Param			offset		query		int		false	"Offset"
//	@Success		200			{object}	ListResp
//	@Router			ROUTE_PREFIX/HANDLER_NAME [get]
func (h *Handler) List(c echo.Context) error {
	var req ListReq 
	if err := c.Bind(&req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ListResp{
	})
}

type ListReq struct {
	Limit  int ` + "`query:\"limit\"`" + `
	Offset int ` + "`query:\"offset\"`" + `
}

type ListResp struct {
}

// One godoc
//
//	@Summary		One
//	@Description	One
//	@Tags			HANDLER_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string	true	"ID"
//	@Success		200			{object}	ListResp
//	@Router			ROUTE_PREFIX/HANDLER_NAME/{id} [get]
func (h *Handler) One(c echo.Context) error {
	var req OneReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, OneResp{
		ID: req.ID,
	})
}

type OneReq struct {
	ID  string ` + "`param:\"id\"`" + `
}

type OneResp struct {
	ID string ` + "`json:\"id\"`" + `
}

// Create godoc
//
//	@Summary		Create
//	@Description	Create
//	@Tags			HANDLER_NAME
//	@Accept			json
//	@Produce		json
//	@Param			req	body		CreateReq	true	"Request"
//	@Success		200	{object}	CreateResp
//	@Router			ROUTE_PREFIX/HANDLER_NAME [post]
func (h *Handler) Create(c echo.Context) error {
	var req CreateReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, CreateResp{
		CreatedName: req.Name,
	})
}

type CreateReq struct {
	Name string ` + "`json:\"name\"`" + `
}

type CreateResp struct {
	CreatedName string ` + "`json:\"created_name\"`" + `
}

// Edit godoc
//
//	@Summary		Edit
//	@Description	Edit
//	@Tags			HANDLER_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID"
//	@Param			req	body		EditReq	true	"Request"
//	@Success		200	{object}	EditResp
//	@Router			ROUTE_PREFIX/HANDLER_NAME/{id} [put]
func (h *Handler) Edit(c echo.Context) error {
	var req EditReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, EditResp{
		ID:   req.ID,
		Name: req.Name,
	})
}

type EditReq struct {
	ID   string ` + "`param:\"id\"`" + `
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
//	@Tags			HANDLER_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string		true	"ID"
//	@Param			req	body		DeleteReq	true	"Request"
//	@Success		200	{object}	DeleteResp
//	@Router			ROUTE_PREFIX/HANDLER_NAME/{id} [delete]
func (h *Handler) Delete(c echo.Context) error {
	var req DeleteReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, DeleteResp{
		ID:   req.ID,
	})
}

type DeleteReq struct {
	ID   string ` + "`param:\"id\"`" + `
	Name string ` + "`json:\"name\"`" + `
}

type DeleteResp struct {
	ID   string ` + "`json:\"id\"`" + `
}
`
	filePath := folderPath + "/handler.go"

	template = strings.ReplaceAll(template, "HANDLER_NAME", handlerName)
	template = strings.ReplaceAll(template, "ROUTE_PREFIX", routePrefix)

	// write to file
	if err := os.WriteFile(filePath, []byte(template), 0644); err != nil {
		return err
	}

	if err := registerHandlerToApp(folderPath, handlerName); err != nil {
		return err
	}

	if err := registerHandlerToRoute(folderPath, handlerName); err != nil {
		return err
	}

	return nil
}
