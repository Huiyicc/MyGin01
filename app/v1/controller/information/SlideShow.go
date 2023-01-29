package information

import (
	`gin01/app/v1/dto`
	`gin01/app/v1/model`
	`gin01/config`
	"github.com/gin-gonic/gin"
	`net/http`
)

func Init(c *gin.Context)  {
	//
	db := config.InitDB()
	//
	// users := model.SildeShow{Src: "http://cq.people.com.cn/NMediaFile/2021/0428/LOCAL202104281145000271146813654.jpg", Disable: 0, Url: "-"}
	//
	// db.Create(&users) // 通过数据的指针来创建
	// return

	SildeShows := []model.SildeShow{}
	err := db.Where("Disable = ?", 0).Find(&SildeShows)
	if err != nil {
		//fmt.Println("==============")
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"data":dto.ToSildeShowDto(SildeShows)})
	//
	//
	// fmt.Println(SildeShows)
	// fmt.Println(c.Request.RequestURI)
	// return
}