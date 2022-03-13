package service

import (
	"context"
	"main/model"
	"main/repository"
)

var TextService = newTextService()

type textService struct {
	Ctx context.Context
}

func newTextService() *textService {
	return &textService{
		Ctx: context.Background(),
	}
}

func (s *textService) GetAll() []*model.Text {
	return repository.TextRepository.GetAll(s.Ctx)
}

func (s *textService) Create(content string) *model.Text {
	return repository.TextRepository.Create(s.Ctx, content)
}

func (s *textService) GetByUUID(uuid string) *model.Text {
	return repository.TextRepository.GetByUUID(s.Ctx, uuid)
}
