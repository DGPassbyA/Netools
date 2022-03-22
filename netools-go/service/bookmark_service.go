package service

import (
	"main/database"
	"main/model"
	"main/repository"
)

var BookmarkService = newBookmarkService()

type bookmarkService struct {
}

func newBookmarkService() *bookmarkService {
	return &bookmarkService{}
}

func (s *bookmarkService) GetByID(id uint) *model.Bookmark {
	return repository.BookmarkRepository.GetByID(database.SDB, id)
}

func (s *bookmarkService) Find() []model.Bookmark {
	return repository.BookmarkRepository.Find(database.SDB)
}

func (s *bookmarkService) Delete(id uint) {
	repository.BookmarkRepository.Delete(database.SDB, id)
}

func (s *bookmarkService) Update(bookmark *model.Bookmark) error {
	return repository.BookmarkRepository.Update(database.SDB, bookmark)
}

func (s *bookmarkService) Create(bookmark *model.Bookmark) error {
	return repository.BookmarkRepository.Create(database.SDB, bookmark)
}
