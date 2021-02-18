package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"

	"github.com/thoohv5/bg/app/api/common"
	"github.com/thoohv5/bg/app/api/markdown"
)

func RegisterRoute(r *gin.Engine, trans ut.Translator) error {
	// 公共模块
	r.LoadHTMLGlob("./html/*")
	r.StaticFS("/static", http.Dir("./static"))
	r.NoRoute(common.New().NoRoute)
	r.NoMethod(common.New().NoRoute)

	// ping
	r.GET("ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 博客
	blog := r.Group("blog")
	{
		// api
		api := blog.Group("api")
		{
			api.GET("list", markdown.New().List)
		}
		// article
		article := blog.Group("article")
		{
			article.GET(":label", markdown.New().Markdown)
		}
	}
	return nil
}
