package dto

import "gin01/app/v1/model"

type ImageDto struct {
	Imgpath	string	`json:"imgpath"`	//图片地址
}
func ToImageDto (image []model.Image) []ImageDto{
	if len(image) == 0 {
		return nil
	}
	images := []ImageDto{}

	for _ , value := range image{
		images = append(images[0:],ImageDto{
			Imgpath: value.Imgpath,
		})
	}
	return images
}

func ToAdminImageDto (image []model.Adminimage) []ImageDto{
	if len(image) == 0 {
		return nil
	}
	images := []ImageDto{}

	for _ , value := range image{
		images = append(images[0:],ImageDto{
			Imgpath: value.Imgpath,
		})
	}
	return images
}
