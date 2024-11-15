package repository

import "base-gin/storage"

var (
	accountRepo *AccountRepository
	personRepo  *PersonRepository
	publishersRepo *PublishersRepository
)

func SetupRepositories() {
	db := storage.GetDB()
	accountRepo = newAccountRepository(db)
	personRepo = newPersonRepository(db)
	publishersRepo = newPublishersRepository(db)
}

func GetAccountRepo() *AccountRepository {
	return accountRepo
}

func GetPersonRepo() *PersonRepository {
	return personRepo
}

func GetPublishersRepo() *PublishersRepository {
	return publishersRepo
}