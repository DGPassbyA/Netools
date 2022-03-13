package service

import (
	"context"
	"main/model"
	"main/repository"
)

var FileService = newFileService()

type fileService struct {
	Ctx context.Context
}

func newFileService() *fileService {
	return &fileService{
		Ctx: context.Background(),
	}
}

func (s *fileService) GetAll() []*model.File {
	return repository.FileRepository.GetAll(s.Ctx)
}

func (s *fileService) GetByUUID(uuid string) *model.File {
	return repository.FileRepository.GetByUUID(s.Ctx, uuid)
}

func (s *fileService) Create(filename string) *model.File {
	return repository.FileRepository.Create(s.Ctx, filename)
}
