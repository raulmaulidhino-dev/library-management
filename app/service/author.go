package service

import (
	"base-gin/app/domain/dto"
	"base-gin/app/repository"
	"base-gin/exception"
)

type AuthorService struct {
	repo *repository.AuthorRepository
}

func newAuthorService(authorRepo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: authorRepo}
}

func (s *AuthorService) Create(params *dto.AuthorCreateReq) (*dto.AuthorCreateResp, error) {
	newItem := params.ToEntity()

	err := s.repo.Create(&newItem)
	if err != nil {
		return nil, err
	}

	var resp dto.AuthorCreateResp
	resp.FromEntity(&newItem)

	return &resp, nil
}

func (s *AuthorService) GetByID(id uint) (dto.AuthorCreateResp, error) {
	var resp dto.AuthorCreateResp

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

func (s *AuthorService) GetList(params *dto.Filter) ([]dto.AuthorCreateResp, error) {
	var resp []dto.AuthorCreateResp

	items, err := s.repo.GetList(params)
	if err != nil {
		return nil, err
	}
	if len(items) < 1 {
		return nil, exception.ErrUserNotFound
	}

	for _, item := range items {
		var t dto.AuthorCreateResp
		t.FromEntity(&item)

		resp = append(resp, t)
	}

	return resp, nil
}

func (s *AuthorService) Update(params *dto.AuthorCreateReq) error {
	if params.ID <= 0 {
		return exception.ErrUserNotFound
	}

	birthDate, err := params.GetBirthDate()
	if err != nil {
		exception.LogError(err, "AuthorService.Update")
		return exception.ErrDateParsing
	}
	params.BirthDate = birthDate.Format("2006-01-02")

	return s.repo.Update(params)
}
