package rest

import (
	"base-gin/server"
	"base-gin/service"

	"github.com/gin-gonic/gin"
)

var (
	accountHandler   *AccountHandler
	personHandler    *PersonHandler
	publisherHandler *PublisherHandler
	authorHandler    *AuthorHandler
	borrowHandler    *BorrowHandler
	bookHandler      *BookHandler
)

func SetupRestHandlers(app *gin.Engine) {
	handler := server.GetHandler()

	accountHandler = NewAccountHandler(
		handler, service.GetAccountService(), service.GetPersonService())
	personHandler = NewPersonHandler(handler, service.GetPersonService())
	publisherHandler = NewPublisherHandler(handler, service.GetPublisherService())
	authorHandler = NewAuthorHandler(handler, service.GetAuthorService())
	borrowHandler = NewBorrowHandler(handler, service.GetBorrowService())
	bookHandler = NewBookHandler(handler, service.GetBookService())

	setupRoutes(app)
}

func setupRoutes(app *gin.Engine) {
	accountHandler.Route(app)
	personHandler.Route(app)
	publisherHandler.Route(app)
	authorHandler.Route(app)
	borrowHandler.Route(app)
	bookHandler.Route(app)
}
