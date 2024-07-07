package example

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	prefix string
}

func NewHandler() *Handler {
	return &Handler{
		prefix: "/example",
	}
}

func (h *Handler) SetRoutes(e *echo.Group) {
	e.GET(h.prefix, h.List)
}

func (h *Handler) List(c echo.Context) error {
	var query ListQuery
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, ListResp{
		Func: "list",
		Tag:  "example",
	})
}

type ListQuery struct {
	TitleIn string `query:"title_in"`
	Limit   int    `query:"limit"`
	Offset  int    `query:"offset"`
}

type ListResp struct {
	Func string `json:"func"`
	Tag  string `json:"tag"`
}
