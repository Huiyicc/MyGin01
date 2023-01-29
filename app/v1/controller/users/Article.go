package users

import "C"
import (
	"gin01/app/v1/dto"
	"gin01/app/v1/model"
	"gin01/app/v1/response"
	"gin01/config"
	"github.com/gin-gonic/gin"
)

// func LikeArticle (c *gin.Context){
// 	db := config.GetDB()
// 	userl, _ := c.Get("user")
// 	if userl.(model.User).ID ==0 {
// 		response.Fail(c,403,"User Is Null")
// 		return
// 	}
// 	liks := model.Likearticle{}
// 	liks.Uid = userl.(model.User).ID
// 	lint,_ := strconv.Atoi(c.Query("id"))
// 	liks.Aid = lint
// 	liks.Like = 1
//
//
// 	ati := model.Article{}
// 	ati.ID=uint(liks.Aid)
//
// 	db.First(&ati)
// 	fmt.Println(ati)
// 	ati.Like  = ati.Like + 1
// 	db.Save(&ati)
// 	response.Success(c,gin.H{},"ok")
// }


// @title    GetArticlelist 获取文章列
// @auth      	lbh             时间（2021/05/08   10:57 )
// @page     	页码    	    int         "当前文章列的页码"
// @count     	获取条数      int         "获取文章列表条数"
// @uid     	用户id       int         "用作查询我发表的文章，可空（默认查询所有用户发表的文章）"
// @return 		["list":文章列]
func GetArticlelist (c *gin.Context) {
	db := config.GetDB()						//获取数据库连接
	page, _ := c.Get("page")					//上下文中传来的page(数据列页码)
	count, _ := c.Get("count")					//上下文中传来的count(数据列条数)
	uid,_ := c.Get("uid")						//获取某个用户的数据时使用(数据列条数) ps：未传不做处理
	isa,_ := c.Get("ugroup")
	isadmin := 0
	if isa!=nil {
		if isa.(int) != 0 {
			isadmin =1
		}

	}

	articlelist := []model.Article{}
	if uid != 0 {
		db.Where("upid = ? and ugroup = ?",uid.(int),isadmin).
			Preload("ImageList").
			Preload("User").
			Order("ID desc").
			Offset((page.(int) -1) * count.(int)).		//算出当前获取数据初始浮标位置
			Limit(count.(int)).
			Find(&articlelist)
		response.Success(c,gin.H{"list":dto.ToArticleDto (articlelist)},"ok")
	} else {
		db.Where("ugroup = ?",isadmin).
			Preload("ImageList").
			Preload("User").
			Order("ID desc").
			Offset((page.(int) -1) * count.(int)).
			Limit(count.(int)).
			Find(&articlelist)
		response.Success(c,gin.H{"list":dto.ToArticleDto (articlelist)},"ok")
	}
}


// @title    GetArticlelist 获取文章列	该方法只获取是否点赞、是否关注发表文章的用户
// @auth      	lbh             时间（2021/05/08   10:57 )
// @uid     	当前用户ID  	UserDto.id		必填、
// @aid     	文章id      	int         	必填、大于0
// @return 		["like":bool,"gz":bool]
func GetArticlelistInfo(c *gin.Context)  {
	db := config.GetDB()						//获取数据库连接
	user, _ := c.Get("user")					//当前用户ID
	aid, _ := c.Get("aid")						//文章id
	auid,_ := c.Get("auid")
	var like = false 							//是否已经喜欢该文章
	var focus = false							//是否已经关注该文章发布者
	likeStruct := &model.Likearticle{}
	focusStruct := &model.UserAttention{}
	if db.Where("uid = ? and aid = ?",user.(model.User).ID,
		aid.(int)).
		First(&likeStruct); likeStruct.ID != 0 && likeStruct.Like != 0{
		like = true
	}

	if db.Where("uid = ? and fid = ?",
		dto.ToUserDto(user.(model.User)).ID,
		auid.(int)).First(&focusStruct); focusStruct.ID != 0 {
		focus = true
	}


	FundLisdb := []model.Funding{}

	db.Where("apid = ?" ,aid).Preload("User").Find(&FundLisdb)

	response.Success(c,gin.H{"like":like,"focus":focus,"funlist":dto.ToFundingDto(FundLisdb)},"ok")
	Article := &model.Article{}
	Article.ID = uint(aid.(int))
	db.First(&Article)
	Article.Heat = Article.Heat + 1
	db.Save(&Article)
}

func DoLike(c *gin.Context)  {
	userl, _ := c.Get("user")
	if userl.(model.User).ID ==0 {
		response.Fail(c,403,"User Is Null")
		return
	}
	lint,_ := c.Get("aid")
	db := config.GetDB()

	liks := model.Likearticle{}
	ati := model.Article{}
	if db.Where("aid = ? and uid = ?",lint.(int),userl.(model.User).ID).First(&liks);liks.ID == 0 {
		liks.Uid = userl.(model.User).ID
		liks.Aid = lint.(int)
		db.Create(&liks)
	}
	db.Where("id = ?",lint.(int)).First(&ati)
	if liks.Like == 0{
		liks.Like = 1
		ati.Like  = ati.Like + 1
	} else {
		liks.Like = 0
		ati.Like  = ati.Like - 1
	}
	db.Save(&ati)
	db.Save(&liks)
	response.Success(c,gin.H{"like":ati.Like},"ok")

/*	db := config.GetDB()						//获取数据库连接
	user, _ := c.Get("user")					//当前用户ID
	aid, _ := c.Get("aid")						//文章id
	likeStruct := &model.Likearticle { }

	if db.Where("Uid = ? and aid = ?",
		dto.ToUserDto(user.(model.User)).ID,
		aid.(int)).
		First(&likeStruct); likeStruct.ID == 0 {
		db.Create(likeStruct)
		DlikeArticle(c,"inc",aid.(int))
	} else {
		db.Delete(likeStruct)
		DlikeArticle(c,"dec",aid.(int))
	}
	response.Success(c,nil,"ok")*/
}

func DlikeArticle (c *gin.Context, dtype string, aid int) {
	db := config.GetDB()
	Article := &model.Article{}
	Article.ID = uint(aid)
	db.First(&Article)
	if dtype == "dec" {
		Article.Like  = Article.Like - 1
	}
	if dtype == "inc" {
		Article.Like  = Article.Like + 1
	}
	db.Save(&Article)
}

func AdminArticle (c *gin.Context) {

}

func DelFundings (c *gin.Context){
	id := c.Query("id")
	db := config.GetDB()
	art := model.Article{}
	db.Where("ID = ?",id).First(&art)
	if art.ID != 0 {
		db.Delete (art)
		response.Success(c,gin.H{},"ok")
	} else {
		response.Fail(c,403,"文章不存在")
	}

}
