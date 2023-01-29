package dto

import (
	"gin01/app/v1/common"
	"gin01/app/v1/model"
	`time`
)

type UserDto struct {
	ID uint `json:"id"`
	Nickname string `json:"nickname"`
	Avatarurl string `json:"avatarurl"`
	Gender int `json:"gender"`
	Signature string `json:"signature"`
	Gznum int `json:"gznum"`
	Numberfans int `json:"numberfans"`
	Ugroup int `json:"ugroup"`
	Shippingaddress string `json:"shippingaddress"`
	Mobile string `json:"mobile"`
	Realname string `json:"realname"`
	Idnumber string `json:"idnumber"`
}
type ArticleGetUserDto struct {
	ID uint `json:"id"`
	Nickname string `json:"nickname"`
	Avatarurl string `json:"avatarurl"`
}
type GZListDto struct {
	CreatedAt time.Time				`json:"createdat"`
	Eachother 	int					`json:"eachother"`	//是否互关
	User		ListGetUserDto	`json:"user"`
}
type FocusListDto struct {
	CreatedAt time.Time				`json:"createdat"`
	Eachother 	int					`json:"eachother"`	//是否互关
	FUser		ListGetUserDto	`json:"user"`
}
type ListGetUserDto struct {
	ID 			uint 	`json:"id"`
	Nickname 	string 	`json:"nickname"`
	Avatarurl 	string 	`json:"avatarurl"`
	Signature	string 	`json:"signature"`	//签名
}
func ToUserDto(user model.User) UserDto {
	return UserDto{
		ID:user.ID,
		Nickname: common.UnicodeEmojiDecode(user.Nickname),
		Ugroup: user.Ugroup,
		Avatarurl: user.Avatarurl,
		Gender:user.Gender,
		Signature:user.Signature,
		Gznum:user.Gznum,
		Numberfans:user.Numberfans,
		Shippingaddress:user.Shippingaddress,
		Mobile:user.Mobile,
		Realname:user.Realname,
		Idnumber:user.Idnumber,
	}
}
func ToGZListGetUserDto(user model.User) ListGetUserDto {
	return ListGetUserDto{
		ID: user.ID,
		Nickname: common.UnicodeEmojiDecode(user.Nickname),
		Avatarurl: user.Avatarurl,
		Signature: user.Signature,
	}
}
func ToArticleGetUserDto(user model.User) ArticleGetUserDto {
	return ArticleGetUserDto{
		ID:user.ID,
		Nickname: common.UnicodeEmojiDecode(user.Nickname),
		Avatarurl: user.Avatarurl,
	}
}
func ToGZListDto (attention []model.UserAttention) []GZListDto{
	if len(attention) == 0 {
		return nil
	}
	data := []GZListDto{}

	for _ , value := range attention{
		data = append(data[0:],GZListDto{
			Eachother:value.Eachother,
			User: ToGZListGetUserDto(value.User),
			CreatedAt:value.CreatedAt,
		})
	}
	return data
}

func ToFucosListDto (attention []model.UserAttention) []FocusListDto{
	if len(attention) == 0 {
		return nil
	}
	data := []FocusListDto{}

	for _ , value := range attention{
		data = append(data[0:],FocusListDto{
			Eachother:value.Eachother,
			FUser: ToGZListGetUserDto(value.Fuser),
			CreatedAt:value.CreatedAt,
		})

	}

	return data
}