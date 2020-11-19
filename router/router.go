package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
	_ "wh110api/docs"
	"wh110api/middleware/jwt"
	"wh110api/pak/setting"
	"wh110api/router/api"
	v1 "wh110api/router/api/v1"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	gin.SetMode(setting.ServeSetting.RunMode)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.GetAuth)
	r.GET("/text/index.html", func(c *gin.Context) {

		WAIHUI := os.Environ()

		c.JSON(200, gin.H{
			"ms":      "全部环境变量",
			"message": WAIHUI,
		})
	})
	r.GET("/text/index2.html", func(c *gin.Context) {

		WAIHUI := os.Environ()

		c.JSON(200, gin.H{
			"ms":      "全部环境变量",
			"message": WAIHUI,
		})
	})
	r.GET("/text/index3.html", func(c *gin.Context) {

		WAIHUI := os.Environ()

		c.JSON(200, gin.H{
			"ms":      "全部环境变量",
			"message": WAIHUI,
		})
	})
	r.GET("/text/index1.html", func(c *gin.Context) {

		var WAIHUI string
		WAIHUI = os.Getenv("WAIHUI")

		c.JSON(200, gin.H{
			"ms":      "WAIHUI环境变量",
			"message": WAIHUI,
		})
	})
	apiv1 := r.Group("/api/v1")
	apiv1.GET("/wh/getnewsdetailbyid", v1.GetNewsDetailById)
	apiv1.GET("/wh/getnewsdetailbyorder", v1.GetNewsDetailByOrder)
	apiv1.GET("/wh/gettoplist", v1.GetNewsTopList)
	apiv1.GET("/wh/getlist", v1.GetNewsList)
	apiv1.GET("/wh/getfriendlist", v1.GetFriendList)
	apiv1.GET("/wh/getrrnews", v1.GetRecommendRandomNewsById)
	apiv1.GET("/wh/getudnews", v1.GetUpDownNewsById)
	apiv1.GET("/wh/getudnewsbo", v1.GetUpDownNewsByOrder)

	apiv1.GET("/cms/adduser", v1.AddAdminUser)
	apiv1.Use(jwt.JWT())
	{
		//apiv1.GET("/tags", v1.GetTags)
		//apiv1.POST("/tags", v1.AddTag)
		//apiv1.PUT("/tags", v1.EditTag)
		//apiv1.DELETE("/tags", v1.DeleteTag)
		//
		//apiv1.GET("/article", v1.GetArticle)
		//apiv1.POST("/article", v1.AddArticles)
		//apiv1.PUT("/article", v1.EcitArticles)
		//apiv1.DELETE("/article", v1.DeleteArticles)
		//apiv1.GET("/user", v1.GetUser)
		//apiv1.POST("/user", v1.CreateUser)
		//apiv1.POST("/seltrader", v1.SelTrader)
	}

	//该中间件以下需要cookie
	//r.Use(cookie.Refresh())
	//r.GET("/index", view.Index)
	//r.GET("/welcome", view.Welcome)
	//
	//memberGroup := r.Group("/member")
	//memberGroup.Use()
	//{
	//	memberGroup.GET("/list", member.List)
	//}

	return r

}
