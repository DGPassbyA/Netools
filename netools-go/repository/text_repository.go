package repository

import (
	"context"
	"main/config"
	"main/database"
	"main/model"
	"main/util"

	"github.com/google/uuid"
)

var TextRepository = newTextRepository()

type textRepository struct {
}

func newTextRepository() *textRepository {
	return &textRepository{}
}

func (r *textRepository) GetAll(ctx context.Context) []*model.Text {
	res, err := database.RDB.Keys(ctx, config.TextKeyPrefix+"*").Result()
	if err != nil {
		return nil
	}
	var ls []*model.Text
	for _, v := range res {
		r, err := database.RDB.Get(ctx, v).Result()
		if err != nil {
			continue
		}
		t := new(model.Text)
		t.Parse(r)
		ls = append(ls, t)
	}
	return ls
}

func (r *textRepository) Create(ctx context.Context, content string) *model.Text {
	_uuid := uuid.New().String()
	key := config.TextKeyPrefix + _uuid
	text := &model.Text{
		Timestamp: util.GetTimestamp(),
		UUID:      _uuid,
		Content:   content,
	}
	_, err := database.RDB.SetEX(ctx, key, text.Stringify(), config.KeyExpireTime).Result()
	if err != nil {
		return nil
	}
	return text
}

func (r *textRepository) GetByUUID(ctx context.Context, uuid string) *model.Text {
	res, err := database.RDB.Get(ctx, config.TextKeyPrefix+uuid).Result()
	if err != nil {
		return nil
	}
	t := new(model.Text)
	t.Parse(res)
	return t
}
