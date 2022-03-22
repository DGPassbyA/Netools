package repository

import (
	"main/model"

	"gorm.io/gorm"
)

var BookmarkRepository = newBookmarkRepository()

type bookmarkRepository struct {
}

func newBookmarkRepository() *bookmarkRepository {
	return &bookmarkRepository{}
}

func (r *bookmarkRepository) GetByID(db *gorm.DB, id uint) (bookmark *model.Bookmark) {
	bookmark = new(model.Bookmark)
	db.First(&bookmark, id)
	return
}

func (r *bookmarkRepository) Find(db *gorm.DB) (bookmarks []model.Bookmark) {
	db.Find(&bookmarks)
	return
}

func (r *bookmarkRepository) Create(db *gorm.DB, bookmark *model.Bookmark) (err error) {
	err = db.Create(bookmark).Error
	return
}

func (r *bookmarkRepository) Update(db *gorm.DB, bookmark *model.Bookmark) (err error) {
	err = db.Save(bookmark).Error
	return
}

func (r *bookmarkRepository) Delete(db *gorm.DB, id uint) {
	db.Delete(&model.Bookmark{}, "id = ?", id)
}
