package model

import (
	"github.com/jinzhu/gorm"
)
//用户信息表
type User struct {
	gorm.Model
	Openid 			string	`gorm:"type:varchar(255);not null"`	//id
	Gender 			int		`gorm:"type:int(11);not null"`	//性别
	Ugroup			int		`gorm:"type:int"`	//组
	Avatarurl 		string 	`gorm:"type:varchar(255);not null"`	//头像
	Nickname 		string	`gorm:"type:varchar(255);not null"`	//昵称
	Signature 		string 	`gorm:"type:varchar(255)"`	//签名
	Gznum 			int		`gorm:"type:int(5);not null"`	//关注人数
	Numberfans 		int	 	`gorm:"type:int(5);not null"`	//粉丝人数
	Shippingaddress string 	`gorm:"type:varchar(255)"` //收货地址
	Mobile 			string 	`gorm:"type:varchar(11)"` //手机号
	Realname 		string 	`gorm:"type:varchar(255)"` //真实姓名
	Idnumber 		string 	`gorm:"type:varchar(255)"` //身份证号
	Articles		[]Article  `gorm:"FOREIGNKEY:Upid;ASSOCIATION_FOREIGNKEY:id"`
}


//用户关注表
type UserAttention struct {
	gorm.Model
	Uid 		uint		`gorm:"type:int;not null"`		//关注者id
	Fid 		int			`gorm:"type:int;not null"`		//被关注者id
	Eachother 	int			`gorm:"type:int(1);not null"`	//是否互关
	User		User		`gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Fid"`
	Fuser		User		`gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Uid"`
	AUser		[]User		`gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Fid"`
}

