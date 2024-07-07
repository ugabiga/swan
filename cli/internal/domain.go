package internal

import (
	"os"
	"strings"
)

func CreateDomain(domainName string, routePrefix string) {
	folderPath := "internal/" + domainName

	if err := os.Mkdir(folderPath, 0755); err != nil {
		panic(err)
	}

	createHandler(folderPath, domainName, routePrefix)
}

func createHandler(
	folderPath string,
	domainName string,
	routePrefix string,
) {
	template := `package ` + domainName + `

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	prefix string
}

func NewHandler() *Handler {
	return &Handler{
		prefix: "/DOMAIN",
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
//	@Router			ROUTE_PREFIX/DOMAIN [get]
func (h *Handler) List(c echo.Context) error {
	var query ListQuery
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
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
//	@Router			ROUTE_PREFIX/DOMAIN/{id} [get]
func (h *Handler) One(c echo.Context) error {
	id := c.Param("id")

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
//	@Router			ROUTE_PREFIX/DOMAIN [post]
func (h *Handler) Create(c echo.Context) error {
	var req CreateReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
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
//	@Router			ROUTE_PREFIX/DOMAIN/{id} [put]
func (h *Handler) Edit(c echo.Context) error {
	id := c.Param("id")

	var req EditReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
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
//	@Failure		400	{object}	nil
//	@Router			ROUTE_PREFIX/DOMAIN/{id} [delete]
func (h *Handler) Delete(c echo.Context) error {
	id := c.Param("id")

	var req DeleteReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
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
		panic(err)
	}

}
