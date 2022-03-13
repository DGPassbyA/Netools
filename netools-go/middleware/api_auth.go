package middleware

import (
	"main/config"
	"main/model"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
)

func CheckAuthToken(ctx iris.Context) {
	token := ctx.GetCookie("token")
	qid_str := ctx.GetCookie("qid")
	if qid_str == "" {
		ctx.JSON(model.JsonErrorMsg(1, "Without qid"))
		ctx.StopExecution()
		return
	}
	qid, err := strconv.ParseInt(qid_str, 10, 64)
	if err != nil {
		ctx.JSON(model.JsonErrorMsg(1, "Without Auth Token"))
		ctx.StopExecution()
		return
	}
	claims := jwt.MapClaims{}
	tokenClaims, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtKey), nil
	})
	if tokenClaims == nil || claims == nil || !tokenClaims.Valid || int64(claims["qid"].(float64)) != qid {
		ctx.JSON(model.JsonErrorMsg(1, "Invalid Token"))
		ctx.StopExecution()
		return
	}
	ctx.Values().Set("qid", qid)
	ctx.Next()
}
