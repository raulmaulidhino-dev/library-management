package repository

import "base-gin/storage"

var (
	accountRepo *AccountRepository
	personRepo  *PersonRepository
	publisherRepo *PublisherRepository
	authorRepo *AuthorRepository
)

func SetupRepositories() {
	db := storage.GetDB()
	accountRepo = newAccountRepository(db)
	personRepo = newPersonRepository(db)
	publisherRepo = newPublisherRepository(db)
	authorRepo = newAuthorRepository(db)
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

func GetAuthorRepo() *AuthorRepository {
    return authorRepo
}