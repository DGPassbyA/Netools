package api

import (
	"main/config"
	"main/model"
	"main/service"
	"main/util"

	"github.com/kataras/iris/v12"
)

type FileController struct {
	Ctx iris.Context
}

func (c *FileController) GetAll() *model.JsonResult {
	res := service.FileService.GetAll()
	return model.JsonDataResult(0, "success", res)
}

func (c *FileController) PostUpload() *model.JsonResult {
	reader, info, err := c.Ctx.FormFile("file")
	if err != nil {
		return model.JsonErrorMsg(-1, "invalid upload")
	}
	defer reader.Close()
	f := service.FileService.Create(info.Filename)
	if f == nil {
		return model.JsonErrorMsg(-1, "upload failed")
	}
	err = f.Store(reader)
	if err != nil {
		return model.JsonErrorMsg(-1, "save file failed")
	}
	return model.JsonDataResult(0, "success", f)
}

func (c *FileController) GetDownload() *model.JsonResult {
	uuid := c.Ctx.URLParam("uuid")
	if util.IsBlank(uuid) {
		return model.JsonErrorMsg(-1, "empty uuid")
	}
	f := service.FileService.GetByUUID(uuid)
	if f == nil {
		return model.JsonErrorMsg(-1, "invalid uuid")
	}
	c.Ctx.SendFile(config.TempFilePath+f.UUID, f.Filename)
	return model.JsonSuccessMsg(0, "success")
}
