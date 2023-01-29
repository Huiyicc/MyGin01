package users

import (
	"encoding/json"
	"gin01/app/v1/common"
	"gin01/app/v1/model"
	"gin01/app/v1/response"
	"gin01/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)
type ImgId struct {
	Id 		float64 	`JSON:"id"`
	Path 	string 		``
}
type ArticleData struct {
	Title string `JSON:"title"`
	Articletext string `JSON:"articletext"`
	Mobile string `JSON:"mobile"`
	Shippingaddress string `JSON:"shippingaddress"`
	Professional float64 `JSON:"professional"`
	ImgIdList []ImgId `JSON:"imgidlist"`
}

func Publish(c *gin.Context){
	var articleData = &ArticleData{}
	//var codeMap = make(map[string]interface{})
	body , _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &articleData)
	if err != nil {
		response.Response(c,http.StatusUnprocessableEntity,423,nil,"")
		return
	}
	var actstact =model.Article{}
	user, _ := c.Get("user")
	//fmt.Println("body:",string(body))
	if user.(model.User).Ugroup==0{
		actstact = model.Article{
			Upid:             user.(model.User).ID,
			Title :           articleData.Title,                         			//取回标题
			Articletext :     common.UnicodeEmojiCode (articleData.Articletext),	//取回文章并进行实体转义
			Mobile :          articleData.Mobile,                         			//联系电话
			Shippingaddress : articleData.Shippingaddress,                			//收货地址
			Professional : int(articleData.Professional),             	  			//职业
			Heat :            0,                                              		//浏览量
			Like : 			  0,													//点赞量
		}
	}else {
		actstact = model.Article{
			Upid:             user.(model.User).ID,
			Ugroup: 			user.(model.User).Ugroup,
			Title :           articleData.Title,                         			//取回标题
			Articletext :     common.UnicodeEmojiCode (articleData.Articletext),	//取回文章并进行实体转义
			Heat :            0,                                              		//浏览量
			Like : 			  0,													//点赞量
		}


	}


	db := config.GetDB()
	if db.NewRecord (actstact) {
		db.Create(&actstact)
		//是否为封面
		for key ,value := range articleData.ImgIdList {
			var imgData = &model.Image{}
			db.Where("id = ?", value.Id).First(&imgData)
			imgData.Apid = actstact.ID
			db.Save(&imgData)
			if key == 0 {
				actstact.Icover = imgData.Imgpath
				db.Save(&actstact)
			}
		}

		response.Success(c,nil,"ok")
	} else {
		response.Fail(c,401,"Article already exists!")
		return
	}
	return
}





