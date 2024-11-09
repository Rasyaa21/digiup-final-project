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

type BorrowHandler struct {
	hr      *server.Handler
	service *service.BorrowService
}

func NewBorrowHandler(hr *server.Handler, borrowService *service.BorrowService) *BorrowHandler {
	return &BorrowHandler{hr: hr, service: borrowService}
}

func (h *BorrowHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootBorrow)
	grp.POST("", h.hr.AuthAccess(), h.create)
	grp.GET("", h.getList)
	grp.GET("/:id", h.getByID)
	grp.PUT("/:id", h.hr.AuthAccess(), h.update)
	grp.DELETE("/:id", h.hr.AuthAccess(), h.delete)
}

// create godoc
//
// @Summary Create a borrow record
// @Description Create a new borrow record.
// @Accept json
// @Tags Borrow
// @Produce json
// @Param borrow body dto.BorrowReq true "Borrow data"
// @Success 201 {object} dto.SuccessResponse[string]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrows [post]
func (h *BorrowHandler) create(c *gin.Context) {
	var req dto.BorrowReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}

	newBorrow, err := h.service.BorrowBook(&req)
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse[string]{
		Success: true,
		Message: "Borrow record created successfully",
		Data:    strconv.FormatUint(uint64(newBorrow.ID), 10),
	})
}

// getList godoc
//
// @Summary Get a list of borrow records
// @Description Retrieve a list of borrow records.
// @Produce json
// @Tags Borrow
// @Param s query int false "Data offset"
// @Param l query int false "Data limit"
// @Success 200 {object} dto.SuccessResponse[[]dto.BorrowRes]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrows [get]
func (h *BorrowHandler) getList(c *gin.Context) {
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

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.BorrowRes]{
		Success: true,
		Message: "List of borrows",
		Data:    data,
	})
}

// getByID godoc
//
// @Summary Get borrow record by ID
// @Description Retrieve a borrow record by ID.
// @Produce json
// @Tags Borrow
// @Param id path int true "Borrow ID"
// @Success 200 {object} dto.SuccessResponse[dto.BorrowRes]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrows/{id} [get]
func (h *BorrowHandler) getByID(c *gin.Context) {
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

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.BorrowRes]{
		Success: true,
		Message: "Borrow details",
		Data:    data,
	})
}

// update godoc
//
// @Summary Update a borrow record
// @Description Update a borrow record's information.
// @Accept json
// @Tags Borrow
// @Produce json
// @Param id path int true "Borrow ID"
// @Param borrow body dto.BorrowReq true "Updated borrow data"
// @Success 200 {object} dto.SuccessResponse[any]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrows/{id} [put]
func (h *BorrowHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	var req dto.UpdateBorrowReq
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
		Message: "Borrow details updated successfully",
	})
}

// delete godoc
//
// @Summary Delete a borrow record
// @Description Delete a borrow record by ID.
// @Produce json
// @Tags Borrow
// @Param id path int true "Borrow ID"
// @Success 200 {object} dto.SuccessResponse[string]
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /borrows/{id} [delete]
func (h *BorrowHandler) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("Invalid ID"))
		return
	}

	err = h.service.ReturnBorrow(uint(id))
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
		Message: "Borrow record deleted successfully",
		Data:    "Borrow ID",
	})
}
