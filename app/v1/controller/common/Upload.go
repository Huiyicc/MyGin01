package common

import (
	"fmt"
	"gin01/app/v1/model"
	"gin01/app/v1/response"
	"gin01/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//type Image struct {
//
//}
func UploadImage(c *gin.Context)  {
	db := config.GetDB()
	fileHeader, err := c.FormFile("image")
	if err != nil {
		response.Fail(c,406,err.Error())
		return
	}
	//fmt.Println(fileHeader.Header)
	fileExt := filepath.Ext(fileHeader.Filename)

	if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".gif" && fileExt != ".jpeg" {
		response.Fail(c,405,"zbc")
		return
	}


	now := time.Now()
	fileDir := fmt.Sprintf("image/%d%d%d/", now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		response.Fail(c,400,err.Error())
		return
	}
	fileName := fmt.Sprintf("%d%s", time.Now().Unix(), fileExt)
	filePathStr := filepath.Join(fileDir,fileName)
	if err = c.SaveUploadedFile(fileHeader, filePathStr); err != nil {
		return
	}
	user, _ := c.Get("user")
	rxstr := strings.Replace(filePathStr, "\\", "/", 2)
	rxstr = "/" +rxstr

	var imgstack = model.Image{
		Upid: user.(model.User).ID,
		Imgpath:rxstr,
	}
	fmt.Println("path:",rxstr)

	db.Create(&imgstack)
	response.Success(c,gin.H{"id":imgstack.ID},"ok")
}

func UpAvatar(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		response.Fail(c,406,err.Error())
		return
	}
	//fmt.Println(fileHeader.Header)
	fileExt := filepath.Ext(fileHeader.Filename)

	if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".gif" && fileExt != ".jpeg" {
		response.Fail(c,405,"zbc")
		return
	}
	now := time.Now()
	fileDir := fmt.Sprintf("image/%d%d%d/", now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		response.Fail(c,400,err.Error())
		return
	}
	fileName := fmt.Sprintf("%d%s", time.Now().Unix(), fileExt)
	println("fileName:",fileName)
	filePathStr := filepath.Join(fileDir,fileName)
	if err = c.SaveUploadedFile(fileHeader, filePathStr); err != nil {
		return
	}
	var rxstr = strings.Replace(filePathStr, "\\", "/", 2)
	rxstr = viper.GetString("server.url") +"/" +rxstr
	response.Success(c,gin.H{"url":rxstr},"ok")
}

