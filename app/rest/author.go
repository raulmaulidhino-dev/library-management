package rest

import (
	"base-gin/app/domain/dto"
	"base-gin/app/service"
	"base-gin/exception"
	"base-gin/server"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	hr      *server.Handler
	service *service.AuthorService
}

func newAuthorHandler(
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
}

// create godoc
//
//	@Summary Create new Author
//	@Description Create new Author.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param newItem body dto.AuthorCreateReq true "Author's detail"
//	@Success 201 {object} dto.SuccessResponse[dto.AuthorCreateReq]
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /authors [post]
func (h *AuthorHandler) create(c *gin.Context) {
	var req dto.AuthorCreateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		h.hr.BindingError(err)
		return
	}

	data, err := h.service.Create(&req)
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse[*dto.AuthorCreateResp]{
		Success: true,
		Message: "Data pengarang berhasil disimpan",
		Data:    data,
	})
}

// getList godoc
//
//	@Summary Get a list of author
//	@Description Get a list of author.
//	@Produce json
//	@Param q query string false "Author's name"
//	@Param s query int false "Data offset"
//	@Param l query int false "Data limit"
//	@Success 200 {object} dto.SuccessResponse[[]dto.AuthorCreateResp]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /authors [get]
func (h *AuthorHandler) getList(c *gin.Context) {
	var req dto.Filter
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}

	data, err := h.service.GetList(&req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrUserNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(exception.ErrDataNotFound.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.AuthorCreateResp]{
		Success: true,
		Message: "Daftar pengarang",
		Data:    data,
	})
}

// getByID godoc
//
//	@Summary Get a author's detail
//	@Description Get a author's detail.
//	@Produce json
//	@Param id path int true "Author's ID"
//	@Success 200 {object} dto.SuccessResponse[dto.AuthorCreateResp]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /authors/{id} [get]
func (h *AuthorHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrUserNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(exception.ErrDataNotFound.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.AuthorCreateResp]{
		Success: true,
		Message: "Detail pengarang",
		Data:    data,
	})
}

// update godoc
//
//	@Summary Update a author's detail
//	@Description Update a author's detail.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "Author's ID"
//	@Param detail body dto.AuthorCreateReq true "Author's detail"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /authors/{id} [put]
func (h *AuthorHandler) update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	var req dto.AuthorCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}
	req.ID = uint(id)

	err = h.service.Update(&req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDateParsing):
			c.JSON(http.StatusBadRequest, h.hr.ErrorResponse(err.Error()))
		case errors.Is(err, exception.ErrUserNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Data pengarang berhasil diperbarui",
	})
}
