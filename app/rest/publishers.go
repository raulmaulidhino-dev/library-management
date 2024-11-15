package rest

import (
	"base-gin/app/domain/dto"
	"base-gin/app/service"
	"base-gin/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublishersHandler struct {
	hr      *server.Handler
	service *service.PublishersService
}

func newPublishersHandler(
    hr *server.Handler,
    publishersService *service.PublishersService,
) *PublishersHandler {
	return &PublishersHandler{hr: hr, service: publishersService}
}

func (h *PublishersHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootPublishers)
	grp.POST("", h.hr.AuthAccess(), h.create)
}

// create godoc
//
//	@Summary Create new Publisher
//	@Description Create new Publisher.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param newItem body dto.PublishersCreateReq true "Publishers's detail"
//	@Success 201 {object} dto.SuccessResponse[dto.PublishersCreateReq]
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 422 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers [post]
func (h *PublishersHandler) create(c *gin.Context) {
	var req dto.PublishersCreateReq
	
	if err := c.ShouldBindJSON(&req); err != nil {
        h.hr.BindingError(err)
        return
	}

	data, err := h.service.Create(&req)
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse[*dto.PublishersCreateResp]{
		Success: true,
		Message: "Data penerbit berhasil disimpam",
		Data: 	data,
	})
}