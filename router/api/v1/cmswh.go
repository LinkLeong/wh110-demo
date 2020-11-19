package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wh110api/pak/e"
)

//后台接口

// @Summary 添加后台用户
// @Produce  json
// @Tags cms
// @Param name formData string true "用户名"
// @Param pwd formData string true "密码"
// @Param type query int false "项目类型" Enums(1, 2,3 )
// @Security ApiKeyAuth
// @Success 200 {string} string "OK"
// @Router /api/v1/cms/addadminuser [get]
func AddAdminUser(c *gin.Context) {
	username := c.PostForm("name")
	pwd := c.PostForm("order")
	code := e.SUCCESS
	if len(username) == 0 || len(pwd) == 0 {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "",
		})
	}

	//tp := com.StrTo(c.DefaultPostForm("type", "1")).MustInt()

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}
