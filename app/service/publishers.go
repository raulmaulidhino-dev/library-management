package service

import (
	"base-gin/app/domain/dto"
	"base-gin/app/repository"
)

type PublishersService struct {
	repo *repository.PublishersRepository
}

func newPublishersService(publisherRepo *repository.PublishersRepository) *PublishersService {
	return &PublishersService{repo: publisherRepo}
}

func (s *PublishersService) Create(params *dto.PublishersCreateReq) (*dto.PublishersCreateResp, error) {
	newItem := params.ToEntity()

	err := s.repo.Create(&newItem)
	if err != nil {
		return nil, err
	}

	var resp dto.PublishersCreateResp
	resp.FromEntity(&newItem)

	return &resp, nil
}