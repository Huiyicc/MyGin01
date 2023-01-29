package users

import (
	"gin01/app/v1/dto"
	"gin01/app/v1/model"
	"gin01/app/v1/response"
	"gin01/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Funding struct {
	Apid		int		`json:"aid"`	//文章id
	Htype		int		`json:"type"`	//类型
	Hdata		string	`json:"data"`	//内
}


func Fundings (c *gin.Context) {
	userl, _ := c.Get("user")
	aid , _:= c.Get("aid")
	htype,_:= c.Get("htype")
	hdata,_:= c.Get("hdata")
	if userl.(model.User).ID ==0 {
		response.Fail(c,401,"User Is Null")
		return
	}
	db := config.GetDB()
	fun := model.Funding{}
	fun.Upid = userl.(model.User).ID
	fun.Apid = aid.(int)
	fun.Htype,_ = strconv.Atoi(htype.(string))
	fun.Hdata = hdata.(string)
	db.Create(&fun)
	art := model.Article{}
	db.Where("ID = ?",aid.(int)).First(&art)
	if art.ID != 0 {
		art.Success=1
		db.Save (&art)
	}
	response.Success(c,gin.H{},"ok")
}

func GetFundingsList (c *gin.Context) {
	userl, _ := c.Get("user")
	if userl.(model.User).ID ==0 {
		response.Fail(c,403,"User Is Null")
		return
	}
	lint := c.Query("aid")
	db := config.GetDB()
	FundLisdb := []model.Funding{}

	db.Where("apid = ?" ,lint).Preload("User").Find(&FundLisdb)
	response.Success(c,gin.H{"funlist":dto.ToFundingDto(FundLisdb)},"ok")
}


func GetFundingsSuccessList(c *gin.Context){
	db := config.GetDB()						//获取数据库连接
	page, _ := c.Get("page")					//上下文中传来的page(数据列页码)
	count, _ := c.Get("count")					//上下文中传来的count(数据列条数)
	articlelist := []model.Article{}
	db.Where("success = ?",1).
		Preload("ImageList").
		Preload("User").
		Offset((page.(int) -1) * count.(int)).
		Limit(count.(int)).
		Find(&articlelist)
		response.Success(c,gin.H{"list":dto.ToArticleDto (articlelist)},"ok")

}
