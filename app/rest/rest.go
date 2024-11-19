package rest

import (
	"base-gin/app/service"
	"base-gin/server"

	"github.com/gin-gonic/gin"
)

var (
	accountHandler *AccountHandler
	personHandler  *PersonHandler
	publisherHandler *PublisherHandler
	authorHandler *AuthorHandler
)

func SetupRestHandlers(app *gin.Engine) {
	handler := server.GetHandler()

	accountHandler = newAccountHandler(
		handler, service.GetAccountService(), service.GetPersonService())
	personHandler = newPersonHandler(handler, service.GetPersonService())
	publisherHandler = newPublisherHandler(handler, service.GetPublisherService())
	authorHandler = newAuthorHandler(handler, service.GetAuthorService())

	setupRoutes(app)
}

func setupRoutes(app *gin.Engine) {
	accountHandler.Route(app)
	personHandler.Route(app)
	publisherHandler.Route(app)
	authorHandler.Route(app)
}
