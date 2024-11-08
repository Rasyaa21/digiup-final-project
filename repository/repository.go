package repository

import "base-gin/storage"

var (
	accountRepo   *AccountRepository
	personRepo    *PersonRepository
	publisherRepo *PublisherRepository
	bookRepo      *BookRepository
	authorRepo    *AuthorRepository
	borrowRepo    *BorrowRepository
)

func SetupRepositories() {
	db := storage.GetDB()
	accountRepo = NewAccountRepository(db)
	personRepo = NewPersonRepository(db)
	publisherRepo = NewPublisherRepository(db)
	bookRepo = newBookRepository(db)
	authorRepo = NewAuthorRepository(db)
	borrowRepo = NewBorrowRepo(db)
}

func GetAccountRepo() *AccountRepository {
	return accountRepo
}

func GetPersonRepo() *PersonRepository {
	return personRepo
}

func GetPublisherRepo() *PublisherRepository {
	return publisherRepo
}

func GetBookRepo() *BookRepository {
	return bookRepo
}

func GetAuthorRepo() *AuthorRepository {
	return authorRepo
}

func GetBorrowRepo() *BorrowRepository {
	return borrowRepo
}
