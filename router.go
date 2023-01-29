package main

import (
	`gin01/app/v1/controller/common`
	"gin01/app/v1/controller/information"
	"gin01/app/v1/controller/users"
	"gin01/app/v1/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func collectRoute(r *gin.Engine) *gin.Engine {
	userGroup := r.Group("/api/user")
	{
		userGroup.GET(	"/info",			middleware.AuthMiddleware(), 			users.GetUserinfo)			//获取用户信息
		userGroup.POST(	"/updateinfo",		middleware.AuthMiddleware(),
														middleware.SetUserMiddleware(),			users.Updateinfo)			//修改用户信息
		userGroup.POST(	"/infoother",		middleware.AuthMiddleware(),
														middleware.OtherUserInfoMiddleware(),	users.GetOtherUserInfo)
		userGroup.POST(	"/publish",		middleware.AuthMiddleware(),			users.Publish)				//发布文章
		userGroup.POST(	"/upload",			middleware.AuthMiddleware(),			common.UploadImage)			//上传图片
		userGroup.POST(	"/upavatar",		middleware.AuthMiddleware(),			common.UpAvatar)			//上传头像
		userGroup.POST(	"/articinfo",		middleware.GetArticleInfoMiddleware(),
														middleware.AuthMiddleware(),			users.GetArticlelistInfo)	//获取文章详情
		userGroup.POST("/delartic",		middleware.AuthMiddleware(),			users.DelFundings)		//删除文章
		userGroup.POST(	"/dolike",			middleware.AuthMiddleware(),
														middleware.DoLikeMiddleware(),			users.DoLike)
		userGroup.POST(	"/dofocus",		middleware.AuthMiddleware(),
														middleware.DoFocusMiddleware(),			users.DoFocus)				//关注
		userGroup.GET(	"/gzlist",			middleware.AuthMiddleware(),			users.GetGZList)			//关注列表
		userGroup.GET(	"/focusList",		middleware.AuthMiddleware(),			users.GetFocusList)			//粉丝列表
		userGroup.POST(	"/funding",		middleware.AuthMiddleware(),
														middleware.DofundingsDMiddleware(),		users.Fundings)				//资助功能
		userGroup.GET(	"/getfundinglist",	middleware.AuthMiddleware(),			users.GetFundingsList)		//取资助列表
		userGroup.POST(	"/adminartic",		middleware.AuthMiddleware(),
														middleware.DoFocusMiddleware(),			users.AdminArticle)			//管理员发帖
		userGroup.POST("/gzarticlist",		middleware.AuthMiddleware(),
														middleware.DogzarticlistMiddleware(),	users.Getgzarticlist)		//获取关注用户文章列表

	}

	r.GET("/api/init",		middleware.APIMiddleware(), information.Init)				//初始化
	r.POST("/api/openid",	users.GetOpenid)											//获取openid
	r.POST("/api/login",	users.Login)												//用户登录



	// r.GET("/api/user/likearticle",middleware.LikeMiddleware(), users.LikeArticle)	//喜欢某文章
	//r.POST("/api/user/funding")		//资助详情

	r.POST("/api/articlelist",	middleware.GetArticleMiddleware(),users.GetArticlelist)
	r.POST("/api/successfully",	middleware.GetArticleMiddleware(),users.GetFundingsSuccessList)

	// r.GET("/api/article",middleware.LikeMiddleware(),users.GetArticleInfo)					//获取单篇文章

	r.StaticFS("/image", http.Dir("image"))										//静态路由:访问图片
	r.MaxMultipartMemory = 8 << 20
	//r.POST("/api/auth/register", users.Register)
	//r.POST("/api/auth/users", controller.Login)
	//r.GET("/api/auth/info",middleware.AuthMiddleware(), controller.Info)
	//r.POST("/api/auth/openid",controller.GetOpenid)
	//r.GET("/api/push/init",controller.GetPush)
	return r
}