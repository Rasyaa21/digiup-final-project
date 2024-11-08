package service

import (
	"base-gin/config"
	"base-gin/repository"
)

var (
	accountService   *AccountService
	personService    *PersonService
	publisherService *PublisherService
	bookService      *BookService
	borrowService    *BorrowService
	authorService    *AuthorService
)

func SetupServices(cfg *config.Config) {
	accountService = NewAccountService(cfg, repository.GetAccountRepo())
	personService = NewPersonService(repository.GetPersonRepo())
	publisherService = NewPublisherService(repository.GetPublisherRepo())
	bookService = NewBookService(repository.GetBookRepo())
	borrowService = NewBorrowService(repository.GetBorrowRepo())
	authorService = NewAuthorService(repository.GetAuthorRepo())
}

func GetAccountService() *AccountService {
	if accountService == nil {
		panic("account service is not initialised")
	}

	return accountService
}

func GetPersonService() *PersonService {
	return personService
}

func GetPublisherService() *PublisherService {
	return publisherService
}

func GetBookService() *BookService {
	return bookService
}

func GetBorrowService() *BorrowService {
	return borrowService
}

func GetAuthorService() *AuthorService {
	return authorService
}
