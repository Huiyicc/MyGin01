package main

import (
	"fmt"
	"gin01/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"os"
)

func main()  {
	InitConfig()
	db := config.InitDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("error" + err.Error()) // 未做写日志处理 一般不会出现异常
		}
	}(db)
	r := gin.Default()
	r = collectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}