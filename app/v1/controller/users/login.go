package users

import (
	"encoding/json"
	"fmt"
	"gin01/app/v1/common"
	"gin01/app/v1/dto"
	"gin01/app/v1/model"
	"gin01/app/v1/response"
	"gin01/config"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type user struct {
	Code string `JSON:"code"`			//验证code1
	Openid string `JSON:"openid"`		//验证code2
	Gender int `JSON:"gender"`			//用户性别
	Nickname string `JSON:"nickname"`	//用户昵称
	AvatarUrl string `JSON:"avatarurl"`	//用户头像链接
	Signature string `JSON:"signature"`	//签名
	Gznum int `JSON:"gznum"`			//关注数
	Numberfans int `JSON:"numberfans"`	//粉丝数
}

func Login (c *gin.Context){
	db := config.GetDB()
	var userMap = &user{}
	//var codeMap = make(map[string]interface{})
	body , _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &userMap)
	if err != nil {
		response.Response(c,http.StatusUnprocessableEntity,423,nil,"")
		return
	}

	openid, loginErr := weapp.Login(viper.GetString("system.wxAppId"),viper.GetString("system.wxSecret"),userMap.Code)
	if loginErr != nil {
		response.Fail(c,40029,"code is invalid")
		//response.Response(ctx,http.StatusUnprocessableEntity,40029,gin.H{"err":res},"code is invalid")
		return
	}

	var ruser model.User
	var reopenid string = ""
	name := common.UnicodeEmojiCode(userMap.Nickname)
	//fmt.Println("code:",userMap.Code,"\nuserMap:",userMap.Openid,"\nOpenID:",openid.OpenID)
	if openid.OpenID != userMap.Openid {
		reopenid = openid.OpenID
	}
	// 登录
	db.Where("openid = ?", openid.OpenID).First(&ruser)
	//fmt.Println("Openid:",ruser.Openid )
	if ruser.ID == 0 {
		ruser = model.User{
			Openid : openid.OpenID,
			Avatarurl : userMap.AvatarUrl,
			Gender : userMap.Gender,
			Nickname : name,
			Signature : "这个人很懒，什么都没留下~",
			Gznum :0,
			Numberfans:0,
		}
		db.Create(&ruser)

	}
	//生成token
	token,errort := common.ReleaseToken (ruser)

	if errort != nil {
		response.Fail(c,491,"no")
		return
	}
	ruser.Nickname = common.UnicodeEmojiDecode(ruser.Nickname)
	response.Success(c,gin.H{"token":token,"openid":reopenid,"userinfo":dto.ToUserDto(ruser) },"OK")
}
func GetUserinfo(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c,gin.H{"user":dto.ToUserDto(user.(model.User))},"ok")
}
func Updateinfo(c *gin.Context) {
	user, _ := c.Get("user")

	avatarurl		,_	:=	c.Get("avatarurl")
	signature		,_	:=	c.Get("signature")
	shippingaddress	,_	:=	c.Get("shippingaddress")
	nickname		,_	:=	c.Get("nickname")
	mobile			,_	:=	c.Get("mobile")
	realname		,_	:=	c.Get("realname")
	idnumber		,_	:=	c.Get("idnumber")

	fmt.Println("ssssssssssssssssssssssssssssssssssss",avatarurl,"   ",
	signature,"   ",
	shippingaddress,"   ",
	nickname,"   ",
	mobile,"   ",
	realname,"   ",
	idnumber,"   ")


	var upda = &model.User{}
	db := config.GetDB()
	db.Where("id = ?",user.(model.User).ID).First(&upda)

	if avatarurl.(string) != "" {
		upda.Avatarurl = avatarurl.(string)
	}
	if  nickname.(string) != ""{
		upda.Nickname = nickname.(string)
	}
	if signature.(string) != "" {
		upda.Signature = signature.(string)
	}
	if shippingaddress.(string) != ""{
		upda.Shippingaddress = shippingaddress.(string)
	}
	if mobile.(string) != ""{
		upda.Mobile = mobile.(string)
	}
	if realname.(string) != ""{
		upda.Realname = realname.(string)
	}
	if idnumber.(string) != ""{
		upda.Idnumber = idnumber.(string)
	}


	db.Save(upda)
	token,errort := common.ReleaseToken (*upda)

	if errort != nil {
		response.Fail(c,491,"no")
		return
	}
	response.Success(c,gin.H{"token":token,"userinfo":dto.ToUserDto(*upda) },"OK")
}


