package model

type JsonResult struct {
	ErrorCode int         `json:"retcode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
}

func JsonErrorMsg(code int, msg string) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   msg,
		Data:      nil,
		Success:   false,
	}
}

func JsonSuccessMsg(code int, msg string) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   msg,
		Data:      nil,
		Success:   true,
	}
}

func JsonDataResult(code int, msg string, data interface{}) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   msg,
		Data:      data,
		Success:   true,
	}
}
