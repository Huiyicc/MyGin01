package model

import (
	"github.com/jinzhu/gorm"
)
//文章表
type Article struct {
	gorm.Model
	Upid 			uint	`gorm:"type:int"`									//发布者ID
	Ugroup			int		`grop:"type:int"`									//发布组
	Icover			string	`gorm:"type:varchar(500)"`							//封面
	Title   		string 	`gorm:"type:varchar(255);not null"`					//标题
	Articletext 	string 	`gorm:"type:varchar(500);not null"`					//文章主体
	Mobile 			string 	`gorm:"type:varchar(11)"` 							//联系电话
	Shippingaddress string 	`gorm:"type:varchar(255)"` 							//收货地址
	Heat   			int    	`gorm:"type:int"`									//浏览量
	Like			int		`grom:"type:int"`									//点赞量
	Professional	int		`gorm:"type:int"` 									//职业
	Success			int		`gorm:"type:int"`									//资助成功
	ImageList 		[]Image `gorm:"FOREIGNKEY:Apid;ASSOCIATION_FOREIGNKEY:ID"`
	User			User	`gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Upid"`
}


// 文章喜欢表
type Likearticle struct {
	gorm.Model
	Uid		uint	`gorm:"typr:int"`		//用户ID
	Aid		int		`gorm:"type:int"`		//文章ID
	Like 	int		`gorm:"type:int"`		//是否关注
}

//文章资助表
type Funding struct {
	gorm.Model
	Upid 		uint	`gorm:"type:int"`			//资助者ID
	Apid		int	`gorm:"type:int"`				//文章ID
	Htype		int		`gorm:"type:int"`			//资助类型
	Hdata		string	`gorm:"type:varchar(255)"`	//资助内容
	User		User	`gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Upid"`
}

//文章图片表
type Image struct {
	ID        uint `gorm:"primary_key"`
	Upid 		uint	`gorm:"type:int"`			//发布者ID
	Apid		uint	`gorm:"type:int"`			//文章ID
	Imgpath		string	`gorm:"type:varchar(255)"`	//图片地址
}

//文章图片表
type Adminimage struct {
	gorm.Model
	Upid 		uint	`gorm:"type:int"`			//发布者ID
	Apid		uint	`gorm:"type:int"`			//文章ID
	Imgpath		string	`gorm:"type:varchar(255)"`	//图片地址
}

//文章图片表
type Adminarticle struct {
	gorm.Model
	Upid 			uint	`gorm:"type:int"`									//发布者ID
	Icover			string	`gorm:"type:varchar(500)"`							//封面
	Title   		string 	`gorm:"type:varchar(255);not null"`					//标题
	Articletext 	string 	`gorm:"type:varchar(500);not null"`					//文章主体
	Heat   			int    	`gorm:"type:int"`									//浏览量
	Like			int		`grom:"type:int"`									//点赞量
	ImageList 		[]Adminimage `gorm:"FOREIGNKEY:Apid;ASSOCIATION_FOREIGNKEY:ID"`
	User			User	`gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Upid"`

}
