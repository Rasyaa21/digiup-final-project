package rest

import (
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/server"
	"base-gin/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	hr      *server.Handler
	service *service.AuthorService
}

func NewAuthorHandler(
	hr *server.Handler,
	authorService *service.AuthorService,
) *AuthorHandler {
	return &AuthorHandler{hr: hr, service: authorService}
}

func (h *AuthorHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootAuthor)
	grp.POST("", h.hr.AuthAccess(), h.create)
	grp.GET("", h.getList)
	grp.GET("/:id", h.getByID)
	grp.PUT("/:id", h.hr.AuthAccess(), h.update)
	grp.DELETE("/:id", h.hr.AuthAccess(), h.delete)
}

// create godoc
//
// @Summary Create an author
// @Description Create a new author in the system.
// @Accept json
// @Produce json
// @Tags Author
// @Param author body dto.AuthorReq true "Author data"
// @Success 201 {object} dto.SuccessResponse[string]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors [post]
func (h *AuthorHandler) create(c *gin.Context) {
	var req dto.AuthorReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}

	err := h.service.Create(&req)
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse[string]{
		Success: true,
		Message: "Author created successfully",
		Data:    "ID of the created author",
	})
}

// getList godoc
//
// @Summary Get a list of authors
// @Description Get a list of authors.
// @Produce json
// @Tags Author
// @Param q query string false "Author's name"
// @Param s query int false "Data offset"
// @Param l query int false "Data limit"
// @Success 200 {object} dto.SuccessResponse[[]dto.AuthorRes]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors [get]
func (h *AuthorHandler) getList(c *gin.Context) {
	var req dto.Filter
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}

	data, err := h.service.GetList(&req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(exception.ErrDataNotFound.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.AuthorRes]{
		Success: true,
		Message: "List of authors",
		Data:    data,
	})
}

// getByID godoc
//
// @Summary Get author by ID
// @Description Get an author's detail by their ID.
// @Produce json
// @Tags Author
// @Param id path int true "Author's ID"
// @Success 200 {object} dto.SuccessResponse[dto.AuthorRes]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors/{id} [get]
func (h *AuthorHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(exception.ErrDataNotFound.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.AuthorRes]{
		Success: true,
		Message: "Author details",
		Data:    data,
	})
}

// update godoc
//
// @Summary Update an author's details
// @Description Update an existing author's information.
// @Accept json
// @Produce json
// @Tags Author
// @Security BearerAuth
// @Param id path int true "Author's ID"
// @Param author body dto.AuthorUpdateReq true "Author's updated data"
// @Success 200 {object} dto.SuccessResponse[any]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors/{id} [put]
func (h *AuthorHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	var req dto.AuthorUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}
	req.ID = uint(id)

	err = h.service.Update(&req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Author details updated successfully",
	})
}

// delete godoc
//
// @Summary Delete an author
// @Description Delete an author by their ID.
// @Produce json
// @Tags Author
// @Param id path int true "Author's ID"
// @Success 200 {object} dto.SuccessResponse[string]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /authors/{id} [delete]
func (h *AuthorHandler) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse[string]{
		Success: true,
		Message: "Author deleted successfully",
		Data:    "Author ID",
	})
}
