package router

import (
	"../api"
	"../db"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Setup ルーターの設定
func Setup() {
	r := gin.Default()

	r.Use(sessions.Sessions("freqmemo", sessions.NewCookieStore([]byte("perfume"))))

	api := api.Handler{
		Db: db.Get(),
	}

	r.POST("/registration", api.Registration)
	r.POST("/login", api.Login)
	r.GET("/logout", api.Logout)

	auth := r.Group("/api")
	auth.Use(api.AuthRequired)
	{
		auth.POST("/add", api.AddFreq)
		auth.GET("/list", api.FreqList)
		auth.GET("/detail/:id", api.FreqDetail)
		auth.DELETE("/delete/:id", api.FreqDelete)
	}

	r.Run(":8080")
}
