package api

import (
	"main/model"

	"github.com/kataras/iris/v12"
)

type TestController struct {
	Ctx iris.Context
}

func (c *TestController) Get() *model.JsonResult {
	qid, err := c.Ctx.Values().GetInt64("qid")
	if err != nil {
		return model.JsonErrorMsg(1, "qid err")
	}
	return model.JsonDataResult(0, "success", map[string]interface{}{
		"qid": qid,
	})
}
