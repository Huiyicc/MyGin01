package users

import (
	`encoding/json`
	`gin01/app/v1/dto`
	`gin01/app/v1/model`
	`gin01/app/v1/response`
	`gin01/config`
	`github.com/gin-gonic/gin`
	`github.com/medivhzhan/weapp`
	`github.com/spf13/viper`
	`io/ioutil`
	`net/http`
)

func GetOpenid(ctx *gin.Context) {
	//var codeMap = map[string]interface{}
	var codeMap = make(map[string]interface{})
	var code string = ""
	body , _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &codeMap)
	if err != nil {
		response.Response(ctx,http.StatusUnprocessableEntity,423,nil,"")
		return
	}
	code = codeMap["code"].(string)

	res, loginErr := weapp.Login(viper.GetString("system.wxAppId"),viper.GetString("system.wxSecret"),code)
	if loginErr != nil {
		response.Fail(ctx,40029,"code is invalid")
		//response.Response(ctx,http.StatusUnprocessableEntity,40029,gin.H{"err":res},"code is invalid")
		return
	}
	response.Success(ctx,gin.H{"openid":res.OpenID},"登录成功")

}

func DoFocus(c *gin.Context)  {
	db := config.GetDB()						//获取数据库连接
	user, _ := c.Get("user")					//当前用户ID
	auid, _ := c.Get("auid")					//文章所属用户id
	uid := dto.ToUserDto(user.(model.User)).ID
	focusStruct := &model.UserAttention{}
	if db.Where("uid = ? and fid = ?",uid,auid).First(&focusStruct); focusStruct.ID == 0 {
		focusStruct = &model.UserAttention{
			Uid: uid,
			Fid: auid.(int),
		}
		db.Create(focusStruct)
		Fuser := model.User{}
		db.Where("id = ?", auid).First(&Fuser)
		Fuser.Numberfans = Fuser.Numberfans + 1
		db.Save(Fuser)
		My := model.User{}
		db.Where("id = ?", uid).First(&My)
		My.Gznum = My.Gznum + 1
		db.Save(My)
		response.Success(c,nil,"ok")
		return
	} else {
		db.Delete(focusStruct)
		Fuser := model.User{}
		db.Where("id = ?", auid).First(&Fuser)
		Fuser.Numberfans = Fuser.Numberfans - 1
		db.Save(Fuser)
		My := model.User{}
		db.Where("id = ?", uid).First(&My)
		My.Gznum = My.Gznum - 1
		db.Save(My)
		response.Success(c,nil,"ok")
		return
	}
}

func GetOtherUserInfo (c *gin.Context)  {
	db := config.GetDB()						//获取数据库连接
	user, _ := c.Get("user")					//当前用户ID
	uid,_ := c.Get("auid")
	otherUser := model.User{}
	if err := db.Where("id = ?", uid).First(&otherUser); err.Error != nil {
		response.Fail(c,403,"no")
	}
	IfFocus := false
	focusStruct := &model.UserAttention{}
	if db.Where("uid = ? and fid = ?",
		dto.ToUserDto(user.(model.User)).ID,
		uid.(int)).
		First(&focusStruct); focusStruct.ID != 0 {
		IfFocus = true
	}
	response.Success(c,gin.H{"user":dto.ToUserDto(otherUser),"IfFocus":IfFocus},"ok")
}

func GetGZList (c *gin.Context)  {
	db := config.GetDB()						//获取数据库连接
	user, _ := c.Get("user")

	uid := dto.ToUserDto(user.(model.User)).ID

	UserAttentionLiset := []model.UserAttention{}
	db.Where("uid = ?",uid).
		Preload("User").
		Find(&UserAttentionLiset)
	response.Success(c,gin.H{"list":dto.ToGZListDto (UserAttentionLiset)},"ok")
}

func GetFocusList (c *gin.Context)  {

	db := config.GetDB()						//获取数据库连接
	user, _ := c.Get("user")				//上下文中传来的page(数据列页码)

	uid := dto.ToUserDto(user.(model.User)).ID

	UserAttentionLiset := []model.UserAttention{}
	db.Where("fid = ?",uid).
		Preload("Fuser").
		Find(&UserAttentionLiset)

	response.Success(c,gin.H{"list":dto.ToFucosListDto (UserAttentionLiset)},"ok")

}



func Getgzarticlist (c *gin.Context) {
	db := config.GetDB()						//获取数据库连接
	page, _ := c.Get("page")					//上下文中传来的page(数据列页码)
	count, _ := c.Get("count")					//上下文中传来的count(数据列条数)
	user,_ := c.Get("user")

	uid := user.(model.User).ID

	articlelist := []model.UserAttention{}
	db.Where("uid = ?",uid).
		Preload("AUser").
		Preload("AUser.Articles").
		Preload("AUser.Articles.ImageList").
		Preload("AUser.Articles.User").
		Order("ID desc").
		Offset((page.(int) -1) * count.(int)).		//算出当前获取数据初始浮标位置
		Limit(count.(int)).
		Find(&articlelist)


	response.Success(c,gin.H{"list":dto.GetUserAttentionGetFUserArticleDto(dto.ToUserAttentionGetFUserArticleDto (articlelist))},"ok")


}

