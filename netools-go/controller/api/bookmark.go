package api

import (
	"main/model"
	"main/service"
	"main/util"

	"github.com/kataras/iris/v12"
)

type BookmarkController struct {
	Ctx iris.Context
}

func (c *BookmarkController) GetAll() *model.JsonResult {
	marks := service.BookmarkService.Find()
	return model.JsonDataResult(0, "success", marks)
}

func (c *BookmarkController) GetBy() *model.JsonResult {
	id := uint(c.Ctx.URLParamUint64("id"))
	mark := service.BookmarkService.GetByID(id)
	if mark == nil {
		return model.JsonErrorMsg(1, "no bookmark")
	}
	return model.JsonDataResult(0, "success", mark)
}

func (c *BookmarkController) PostCreate() *model.JsonResult {
	name := c.Ctx.FormValue("name")
	content := c.Ctx.FormValue("content")
	t := c.Ctx.FormValue("type")
	if util.IsBlank(content) || util.IsBlank(name) || util.IsBlank(t) {
		return model.JsonErrorMsg(-1, "empty params")
	}
	err := service.BookmarkService.Create(&model.Bookmark{
		Name:    name,
		Content: content,
		Type:    t,
	})
	if err != nil {
		return model.JsonErrorMsg(1, "create failed")
	}
	return model.JsonSuccessMsg(0, "success")
}

func (c *BookmarkController) GetDelete() *model.JsonResult {
	id := uint(c.Ctx.URLParamUint64("id"))
	service.BookmarkService.Delete(id)
	return model.JsonSuccessMsg(0, "success")
}
