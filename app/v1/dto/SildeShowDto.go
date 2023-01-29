package dto

import (
	`gin01/app/v1/model`
)

type SildeShow struct {
	ID uint 		`json:"id"`
	Src string 		`json:"src"`
	Url string 		`json:"url"`
}

//
func ToSildeShowDto(SildeShowD []model.SildeShow ) []SildeShow {
	if len(SildeShowD) == 0 {
		return nil
	}
	//fmt.Println(len(SildeShowD))
	a := []SildeShow{}
	for _ , value := range SildeShowD{
		//fmt.Println(key)
		//fmt.Println(value)

		a = append(a[0:],SildeShow{
			ID: value.ID,
			Src: value.Src,
			Url: value.Url,
		})
	}
	return a
}