package repository

import (
	"context"
	"main/config"
	"main/database"
	"main/model"
	"main/util"

	"github.com/google/uuid"
)

var FileRepository = newFileRepository()

type fileRepository struct {
}

func newFileRepository() *fileRepository {
	return &fileRepository{}
}

func (r *fileRepository) GetAll(ctx context.Context) []*model.File {
	res, err := database.RDB.Keys(ctx, config.FileKeyPrefix+"*").Result()
	if err != nil {
		return nil
	}
	var ls []*model.File
	for _, v := range res {
		r, err := database.RDB.Get(ctx, v).Result()
		if err != nil {
			continue
		}
		f := new(model.File)
		f.Parse(r)
		ls = append(ls, f)
	}
	return ls
}

func (r *fileRepository) GetByUUID(ctx context.Context, uuid string) *model.File {
	res, err := database.RDB.Get(ctx, config.FileKeyPrefix+uuid).Result()
	if err != nil {
		return nil
	}
	f := new(model.File)
	f.Parse(res)
	return f
}

func (r *fileRepository) Create(ctx context.Context, filename string, size int64) *model.File {
	_uuid := uuid.New().String()
	key := config.FileKeyPrefix + _uuid
	file := &model.File{
		Timestamp: util.GetTimestamp(),
		UUID:      _uuid,
		Filename:  filename,
		Size:      size,
	}
	_, err := database.RDB.SetEX(ctx, key, file.Stringify(), config.KeyExpireTime).Result()
	if err != nil {
		return nil
	}
	return file
}
