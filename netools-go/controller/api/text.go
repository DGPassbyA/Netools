package api

import (
	"main/model"
	"main/service"
	"main/util"

	"github.com/kataras/iris/v12"
)

type TextController struct {
	Ctx iris.Context
}

func (c *TextController) GetAll() *model.JsonResult {
	res := service.TextService.GetAll()
	return model.JsonDataResult(0, "success", res)
}

func (c *TextController) PostAdd() *model.JsonResult {
	content := c.Ctx.FormValue("content")
	if util.IsBlank(content) {
		return model.JsonErrorMsg(-1, "empty content")
	}
	res := service.TextService.Create(content)
	if res == nil {
		return model.JsonErrorMsg(-1, "failed to add text")
	}
	return model.JsonDataResult(0, "success", res)
}

func (c *TextController) GetGet() *model.JsonResult {
	uuid := c.Ctx.URLParam("uuid")
	if util.IsBlank(uuid) {
		return model.JsonErrorMsg(-1, "empty uuid")
	}
	t := service.TextService.GetByUUID(uuid)
	if t == nil {
		return model.JsonErrorMsg(-1, "invalid uuid")
	}
	return model.JsonDataResult(0, "success", t)
}
