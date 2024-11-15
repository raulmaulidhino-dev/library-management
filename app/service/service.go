package service

import (
	"base-gin/app/repository"
	"base-gin/config"
)

var (
	accountService *AccountService
	personService  *PersonService
	publishersService *PublishersService
)

func SetupServices(cfg *config.Config) {
	accountService = newAccountService(cfg, repository.GetAccountRepo())
	personService = newPersonService(repository.GetPersonRepo())
	publishersService = newPublishersService(repository.GetPublishersRepo())
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

func GetPublishersService() *PublishersService {
	return publishersService
}