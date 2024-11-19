package service

import (
	"base-gin/app/repository"
	"base-gin/config"
)

var (
	accountService *AccountService
	personService  *PersonService
	publisherService *PublisherService
	authorService *AuthorService
)

func SetupServices(cfg *config.Config) {
	accountService = newAccountService(cfg, repository.GetAccountRepo())
	personService = newPersonService(repository.GetPersonRepo())
	publisherService = newPublisherService(repository.GetPublisherRepo())
	authorService = newAuthorService(repository.GetAuthorRepo())
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

func GetAuthorService() *AuthorService {
    return authorService
}