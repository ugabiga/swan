package example

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ugabiga/swan/core"
	"github.com/ugabiga/swan/utl"
)

type Handler struct {
	prefix       string
	logger       *slog.Logger
	eventEmitter *core.EventEmitter
}

func NewHandler(
	logger *slog.Logger,
	eventEmitter *core.EventEmitter,
) *Handler {
	return &Handler{
		prefix:       "/example",
		logger:       logger,
		eventEmitter: eventEmitter,
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
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int	false	"Limit"
//	@Param			offset	query		int	false	"Offset"
//	@Success		200		{object}	ListResp
//	@Router			/example [get]
func (h *Handler) List(c echo.Context) error {
	var query ListQuery
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, utl.ValidateRequestAllEmptyResponse())
	}

	validationErrors, ok := utl.ValidateRequestStruct(query)
	if !ok {
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	if err := h.eventEmitter.Emit("example", []byte(`{"message": "List"}`)); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ListResp{
		Tag: "list",
	})
}

type ListQuery struct {
	Limit  int `query:"limit" validate:"required"`
	Offset int `query:"offset"`
}

type ListResp struct {
	Tag string `json:"func"`
}

// One godoc
//
//	@Summary		One
//	@Description	One
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"ID"
//	@Param			limit	query		int		false	"Limit"
//	@Param			offset	query		int		false	"Offset"
//	@Success		200		{object}	ListResp
//	@Router			/example/{id} [get]
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
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type OneResp struct {
	ID string `json:"id"`
}

// Create godoc
//
//	@Summary		Create
//	@Description	Create
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Param			req	body		CreateReq	true	"Request"
//	@Success		200	{object}	CreateResp
//	@Router			/example [post]
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
	Name string `json:"name" validate:"required"`
}

type CreateResp struct {
	Name string `json:"name"`
}

// Edit godoc
//
//	@Summary		Edit
//	@Description	Edit
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID"
//	@Param			req	body		EditReq	true	"Request"
//	@Success		200	{object}	EditResp
//	@Router			/example/{id} [put]
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
	Name string `json:"name"`
}

type EditResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Delete godoc
//
//	@Summary		Delete
//	@Description	Delete
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string		true	"ID"
//	@Param			req	body		DeleteReq	true	"Request"
//	@Success		200	{object}	DeleteResp
//	@Router			/example/{id} [delete]
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
	Name string `json:"name"`
}

type DeleteResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
