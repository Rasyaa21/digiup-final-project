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

type BookHandler struct {
	hr      *server.Handler
	service *service.BookService
}

func NewBookHandler(hr *server.Handler, bookService *service.BookService) *BookHandler {
	return &BookHandler{hr: hr, service: bookService}
}

func (h *BookHandler) Route(app *gin.Engine) {
	grp := app.Group("/books")
	grp.POST("", h.hr.AuthAccess(), h.create)
	grp.GET("", h.getList)
	grp.GET("/:id", h.getByID)
	grp.PUT("/:id", h.hr.AuthAccess(), h.update)
	grp.DELETE("/:id", h.hr.AuthAccess(), h.delete)
}

// create godoc
//
// @Summary Create a book
// @Description Create a new book.
// @Accept json
// @Produce json
// @Tags Book
// @Param book body dto.BookReq true "Book data"
// @Success 201 {object} dto.SuccessResponse[string]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books [post]
func (h *BookHandler) create(c *gin.Context) {
	var req dto.BookReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}

	if err := h.service.Create(&req); err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse[string]{
		Success: true,
		Message: "Book created successfully",
	})
}

// getList godoc
//
// @Summary Get a list of books
// @Description Retrieve a list of books.
// @Produce json
// @Tags Book
// @Param s query int false "Data offset"
// @Param l query int false "Data limit"
// @Success 200 {object} dto.SuccessResponse[[]dto.BookRes]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books [get]
func (h *BookHandler) getList(c *gin.Context) {
	var req dto.Filter
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}

	data, err := h.service.GetDataList(&req)
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.BookRes]{
		Success: true,
		Message: "List of books",
		Data:    data,
	})
}

// getByID godoc
//
// @Summary Get book by ID
// @Description Retrieve a book by ID.
// @Produce json
// @Tags Book
// @Param id path int true "Book ID"
// @Success 200 {object} dto.SuccessResponse[dto.BookRes]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books/{id} [get]
func (h *BookHandler) getByID(c *gin.Context) {
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

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.BookRes]{
		Success: true,
		Message: "Book details",
		Data:    data,
	})
}

// update godoc
//
// @Summary Update a book
// @Description Update book information.
// @Accept json
// @Produce json
// @Tags Book
// @Param id path int true "Book ID"
// @Param book body dto.BookUpdateReq true "Updated book data"
// @Success 200 {object} dto.SuccessResponse[any]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books/{id} [put]
func (h *BookHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	var req dto.BookUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}

	err = h.service.Update(uint(id), &req)
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
		Message: "Book updated successfully",
	})
}

// delete godoc
//
// @Summary Delete a book
// @Description Delete a book by ID.
// @Produce json
// @Tags Book
// @Param id path int true "Book ID"
// @Success 200 {object} dto.SuccessResponse[string]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books/{id} [delete]
func (h *BookHandler) delete(c *gin.Context) {
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
		Message: "Book deleted successfully",
	})
}
