package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"strconv"
	"wh110api/pak/e"
	"wh110api/service"
)

// @Summary 查询新闻详情
// @Produce  json
// @Tags wh
// @Param id query string true "新闻id"
// @Security ApiKeyAuth
// @Success 200 {string} string "OK"
// @Router /api/v1/wh/getnewsdetailbyid [get]
func GetNewsDetailById(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	news := service.GetById(id)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": news,
	})
}

// @Summary 查询新闻详情
// @Produce  json
// @Tags wh
// @Param order query string true "新闻order"
// @Security ApiKeyAuth
// @Success 200 {string} string "OK"
// @Router /api/v1/wh/getnewsdetailbyorder [get]
func GetNewsDetailByOrder(c *gin.Context) {
	id := c.DefaultQuery("order", "0")
	news := service.GetByOrder(id)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": news,
	})
}

// @Summary 查询新闻推荐列表
// @Produce  json
// @Tags wh
// @Param page query int false "页码,默认为1" default(1)
// @Param size query int false "每页数量,默认为10" default(10)
// @Param type query int false "项目类型" Enums(1, 2, 3 )
// @Security ApiKeyAuth
// @Success 200 {string} string "OK"
// @Router /api/v1/wh/gettoplist [get]
func GetNewsTopList(c *gin.Context) {
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt64()
	size := com.StrTo(c.DefaultQuery("size", "10")).MustInt64()
	tp := com.StrTo(c.DefaultQuery("type", "1")).MustInt()
	news := service.GetTopList(page, size, tp)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": news,
	})
}

// @Summary 查询新闻列表
// @Produce  json
// @Tags wh
// @Param page query int false "页码,默认为1" default(1)
// @Param size query int false "每页数量,默认为10" default(10)
// @Param type query int false "项目类型" Enums(1, 2, 3 )
// @Param isasc query bool false "是否为升序"
// @Security ApiKeyAuth
// @Success 200 {string} string "OK"
// @Router /api/v1/wh/getlist [get]
func GetNewsList(c *gin.Context) {
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt64()
	size := com.StrTo(c.DefaultQuery("size", "10")).MustInt64()
	//类型
	tp := com.StrTo(c.DefaultQuery("type", "1")).MustInt()
	fmt.Println(c.Query("isasc"))
	//是否为升序
	isasc, _ := strconv.ParseBool(c.Query("isasc"))
	news, count := service.GetList(page, size, tp, isasc)
	code := e.SUCCESS
	m := make(map[string]interface{})
	m["count"] = count
	m["data"] = news
	//b, err := json.Marshal(m)
	//fmt.Println (err)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": m,
	})
}

// @Summary 查询友链列表
// @Produce  json
// @Tags wh
// @Param type query int false "项目类型" Enums(1, 2, 3 )
// @Security ApiKeyAuth
// @Success 200 {string} string "OK"
// @Router /api/v1/wh/getfriendlist [get]
func GetFriendList(c *gin.Context) {
	//类型
	tp := com.StrTo(c.DefaultQuery("type", "1")).MustInt()
	fr := service.GetFriendList(tp)
	code := e.SUCCESS

	//b, err := json.Marshal(m)
	//fmt.Println (err)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": fr,
	})
}

// @Summary 根据新闻id获取随机推荐
// @Produce  json
// @Tags wh
// @Param id query string true "新闻标识"
// @Param size query int false "推荐的数量" default(10)
// @Param type query int false "项目类型" Enums(1, 2, 3 )
// @Security ApiKeyAuth
// @Success 200 {string} string "OK"
// @Router /api/v1/wh/getrrnews [get]
func GetRecommendRandomNewsById(c *gin.Context) {
	size := com.StrTo(c.DefaultQuery("size", "10")).MustInt64()
	tp := com.StrTo(c.DefaultQuery("type", "1")).MustInt()
	id := c.DefaultQuery("id", "1")
	news := service.GetRecommendRandomNewsById(size, tp, id)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": news,
	})
}

// @Summary 根据新闻id获取上下篇
// @Produce  json
// @Tags wh
// @Param id query string true "新闻标识"
// @Param type query int false "项目类型" Enums(1, 2,3 )
// @Security ApiKeyAuth
// @Success 200 {string} string "OK"
// @Router /api/v1/wh/getudnews [get]
func GetUpDownNewsById(c *gin.Context) {
	tp := com.StrTo(c.DefaultQuery("type", "1")).MustInt()
	id := c.DefaultQuery("id", "1")
	up, down := service.GetUpDownNewsById(tp, id)
	code := e.SUCCESS

	var ud = make(map[string]interface{})
	ud["up"] = up
	ud["down"] = down

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": ud,
	})
}

// @Summary 根据新闻order获取上下篇
// @Produce  json
// @Tags wh
// @Param order query string true "新闻标识"
// @Param type query int false "项目类型" Enums(1, 2,3 )
// @Security ApiKeyAuth
// @Success 200 {string} string "OK"
// @Router /api/v1/wh/getudnewsbo [get]
func GetUpDownNewsByOrder(c *gin.Context) {
	tp := com.StrTo(c.DefaultQuery("type", "1")).MustInt()
	id := c.DefaultQuery("order", "1")
	up, down := service.GetUpDownNewsByOrder(tp, id)
	code := e.SUCCESS

	var ud = make(map[string]interface{})
	ud["up"] = up
	ud["down"] = down

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": ud,
	})
}
