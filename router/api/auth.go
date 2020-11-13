package api

import (
	"github.com/astaxie/beego/validation"
	"net/http"
	"time"
	"wh110api/model"
	"wh110api/pak/e"
	"wh110api/pak/logging"

	"github.com/gin-gonic/gin"
	"wh110api/pak/jwt"
)

type auth struct {
	UserName string `valid:"Required; MaxSize(50)"`
	PassWord string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var code int
	//if username != "lynn" && password != "d58w@#4f" {
	//	code = e.INVALID_PARAMS
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": code,
	//		"msg":  e.GetMsg(code),
	//	})
	//	return
	//}

	valid := validation.Validation{}
	a := auth{UserName: username, PassWord: password}
	ok, er := valid.Valid(&a)
	if er != nil {
		logging.Info("token解析错误", er)
	}
	data := make(map[string]interface{})
	success := false
	code = e.INVALID_PARAMS
	if ok {
		isExist := model.CheckAuth(username, password)
		if isExist {
			token, err := jwt.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				data["expire"] = time.Now().Add(3 * time.Hour).Unix()
				code = e.SUCCESS
				success = true
			}
		} else {

			code = e.INVALID_PARAMS
		}

	} //else {
	//	for _, err := range valid.Errors {
	//		//log.Println(err.Key, err.Message)
	//		logging.Info(err.Key, err.Message)
	//	}
	//}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     e.GetMsg(code),
		"success": success,
		"data":    data,
	})
}
