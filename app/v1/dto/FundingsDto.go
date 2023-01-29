package dto

import (
	"gin01/app/v1/common"
	"gin01/app/v1/model"
	"time"
)




type FundingDto struct {
	Htype		int		`json:"type"`	//类型
	Hdata		string	`json:"data"`	//内容
	User		FArticleGetUserDto	`json:"user"`
	CreatedAt	time.Time	`json:"createdat"`
}
type FArticleGetUserDto struct {
	ID uint `json:"id"`
	Nickname string `json:"nickname"`
	Avatarurl string `json:"avatarurl"`
}
func GetUserByF(user model.User) FArticleGetUserDto{
	return FArticleGetUserDto{
		ID:        user.ID,
		Nickname: common.UnicodeEmojiDecode(user.Nickname) ,
		Avatarurl: user.Avatarurl,
	}
}
func ToFundingDto (funding []model.Funding) []FundingDto{
	if len(funding) == 0 {
		return nil
	}
	data := []FundingDto{}
	for _ , value := range funding{
		data = append(data[0:],FundingDto{
			Htype: value.Htype,
			Hdata: value.Hdata,
			User: GetUserByF(value.User),
			CreatedAt:value.CreatedAt,
		})
	}
	return data
}