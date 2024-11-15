package service

import (
	"base-gin/app/domain/dto"
	"base-gin/app/repository"
	"base-gin/exception"
)

type PersonService struct {
	repo *repository.PersonRepository
}

func newPersonService(personRepo *repository.PersonRepository) *PersonService {
	return &PersonService{repo: personRepo}
}

func (s *PersonService) GetAccountProfile(accountID uint) (dto.AccountProfileResp, error) {
	var resp dto.AccountProfileResp

	item, err := s.repo.GetByAccountID(accountID)
	if err != nil {
		return resp, err
	}

	resp.FromPerson(&item)

	return resp, nil
}

func (s *PersonService) GetByID(id uint) (dto.PersonDetailResp, error) {
	var resp dto.PersonDetailResp

	item, err := s.repo.GetByID(id)
	if err != nil {
		return resp, err
	}
	if item == nil {
		return resp, exception.ErrUserNotFound
	}

	resp.FromEntity(item)

	return resp, nil
}

func (s *PersonService) GetList(params *dto.Filter) ([]dto.PersonDetailResp, error) {
	var resp []dto.PersonDetailResp

	items, err := s.repo.GetList(params)
	if err != nil {
		return nil, err
	}
	if len(items) < 1 {
		return nil, exception.ErrUserNotFound
	}

	for _, item := range items {
		var t dto.PersonDetailResp
		t.FromEntity(&item)

		resp = append(resp, t)
	}

	return resp, nil
}

func (s *PersonService) Update(params *dto.PersonUpdateReq) error {
	if params.ID <= 0 {
		return exception.ErrUserNotFound
	}

	birthDate, err := params.GetBirthDate()
	if err != nil {
		exception.LogError(err, "PersonService.Update")
		return exception.ErrDateParsing
	}
	params.BirthDate = birthDate

	return s.repo.Update(params)
}
