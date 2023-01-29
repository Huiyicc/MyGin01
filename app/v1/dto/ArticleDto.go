package dto

import (
	`gin01/app/v1/model`
)

type Article struct {
	ID				uint				`json:"id"`					//文章ID
	Ugroup 			int					`json:"ugroup"`				//组
	Heat   			int    				`json:"heat"`				//浏览量
	Like			int					`json:"like"`				//点赞数
	Title   		string 				`json:"title"`				//标题
	User			ArticleGetUserDto 	`json:"user"`
	Mobile 			string 				`json:"mobile"` 			//联系电话
	Icover			string				`json:"icover"`				//封面
	ImageList		[]ImageDto 			`json:"imagelist"`
	Articletext 	string 				`json:"articletext"`		//文章主体
	Professional	int					`json:"professional"` 		//职业
	Shippingaddress string 				`json:"shippingaddress"`	//收货地址
	Loading			int					`json:"loading"`
	Success			int					`json:"success"`			//是否被资助过
}

func ToArticleDto (article []model.Article) []Article{
	if len(article) == 0 {
		return nil
	}
	var data []Article
	for _ , value := range article{
		data = append(data[0:],Article{
			ID:value.ID,
			Title:value.Title,
			Ugroup:value.Ugroup,
			Heat:value.Heat,
			Icover:value.Icover,
			Like: value.Like,
			Loading: 0,
			ImageList:ToImageDto(value.ImageList),
			User:ToArticleGetUserDto(value.User),
			Articletext:value.Articletext,
			Mobile:value.Mobile,
			Shippingaddress:value.Shippingaddress,
			Professional:value.Professional,
			Success:value.Success,
		})
	}
	return data
}

type UserAttentionGeFUserArticleByUserDto struct {
	Article []Article
}

type UserAttentionGetFUserArticleByAttentionDto struct {
	User []UserAttentionGeFUserArticleByUserDto
}

func ToUserAttentionGetFUserArticleByAttentionDto(user []model.User) []UserAttentionGeFUserArticleByUserDto {
	if len(user) == 0 {
		return nil
	}
	var artByUser []UserAttentionGeFUserArticleByUserDto
	for _ , value := range user{
		artByUser = append(artByUser[0:],UserAttentionGeFUserArticleByUserDto{
			Article: ToArticleDto(value.Articles),
		})
	}
	return artByUser
}

func ToUserAttentionGetFUserArticleDto (att []model.UserAttention) []UserAttentionGetFUserArticleByAttentionDto {
	if len(att) == 0 {
		return nil
	}
	var artByAtt []UserAttentionGetFUserArticleByAttentionDto
	for _ , value := range att{
		artByAtt = append(artByAtt[0:],UserAttentionGetFUserArticleByAttentionDto{
			User: ToUserAttentionGetFUserArticleByAttentionDto(value.AUser),
		})
	}
	return artByAtt
}

func GetUserAttentionGetFUserArticleDto (att []UserAttentionGetFUserArticleByAttentionDto) [] Article {
	if len(att) == 0 {
		return nil
	}
	var data []Article
	for _, val1 := range att{

		for _,val2 := range val1.User {

			for _,val3 := range val2.Article {
				data = append(data[0:],Article{
					ID:val3.ID,
					Title:val3.Title,
					Ugroup:val3.Ugroup,
					Heat:val3.Heat,
					Icover:val3.Icover,
					Like: val3.Like,
					Loading: 0,
					ImageList:val3.ImageList,
					User:val3.User,
					Articletext:val3.Articletext,
					Mobile:val3.Mobile,
					Shippingaddress:val3.Shippingaddress,
					Professional:val3.Professional,
					Success:val3.Success,
				})
			}
		}
	}
	return data
}
