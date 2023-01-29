package middleware

import (
	"encoding/hex"
	`fmt`
	"gin01/app/v1/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	`strings`
)

type Signst struct {
	Sign  string `form:"sign" binding:"required,Sign"`
}
var Sign validator.Func = func(fl validator.FieldLevel) bool {

	signString := fl.Field().Interface().(string)												//读取Token sign头
	if signString == "" || !strings.HasPrefix(signString, "sBearer") {										// sign头 sBearer
		return false
	}

	return true
}
type ArticleListValidate struct {
	Page 	int `form:"page" json:"page" binding:"required"`
	Count 	int `form:"count" json:"count" binding:"required"`
	Uid		int	`form:"uid" json:"uid"`
	Ugroup 	int	`form:"ugroup" json:"ugroup"`
}
type ArticleInfoValidate struct {
	Aid		int	`json:"aid" binding:"required,min=1"`
	AUid	int	`json:"auid" binding:"required,min=1"`
}
type DoLikeArticleValidate struct {
	Aid		int	`json:"aid" binding:"required,min=1"`
}
type DoFocusUserValidate struct {
	AUid	int	`json:"auid" binding:"required,min=1"`
}
type DofundingValidate struct {
	Aid	int	`json:"aid" binding:"required,min=1"`
	Htype string `json:"htype" binding:"required,max=1"`
	Hdata string `json:"data" binding:"required"`
}

type GzarticlistValidate struct {
	Page 	int `form:"page" json:"page" binding:"required"`
	Count 	int `form:"count" json:"count" binding:"required"`
}

type SetuserValidate struct {
	Avatarurl			string `form:"avatarurl" json:"avatarurl"`
	Nickname			string `form:"nickname" json:"nickname"`
	Signature			string `form:"signature" json:"signature"`
	Numberfans			string `form:"numberfans" json:"numberfans"`
	Shippingaddress		string `form:"shippingaddress" json:"shippingaddress"`
	Mobile				string `form:"mobile" json:"mobile"`
	Realname			string `form:"realname" json:"realname"`
	Idnumber			string `form:"idnumber" json:"idnumber"`
}


func APIMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			err := v.RegisterValidation("Sign", Sign)
			if err != nil {
				return
			}
		}

		var sign Signst
		if err := ctx.ShouldBindWith(&sign, binding.Header); err == nil {
			signString := sign.Sign[8:]
			signByte, _ := hex.DecodeString(signString)
			key := []byte(viper.GetString("system.key")) // 加密的密钥
			decryptCode := common.AesDecryptCBC(signByte, key)
			signMap := common.Parse_str(string(decryptCode))
			ok, did:= VerifyTheSign(signMap); if ok != true {
				ctx.JSON(http.StatusUnauthorized, gin.H{"code":401,"msg":"权限不足"})
				ctx.Abort()
				return
			}
			//fmt.Println(did)
			ctx.Set("did",did)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":400,"msg":"权限不足"})
			ctx.Abort()
			return
		}


	}
}
func VerifyTheSign(sign map[string]interface{}) (bool,string) {

	did, 		didState 	 	:= 	sign["did"];
	version, 	versionState 	:= 	sign["version"];
	timeStamp,  timestampState  :=	sign["timeStamp"];
	// fmt.Println(did)
	// fmt.Println(version)
	// fmt.Println(timeStamp)
	if didState 		== false {
		return false, string("")
	}

	if versionState 	== false || version != viper.GetString("system.version"){
		return false, string("")
	}

	if timestampState 	== false {
		return false, string("")
	}

											timeStampString , _ := strconv.ParseFloat(timeStamp.(string), 64)
											timeStampInt64 := int64(timeStampString)

	if timeStampInt64 	<= 1	 {
		return false, string("")
	}

	// if (time.Now().Unix() - timeStampInt64) > 600 {
	// 	return false
	// }

	// whereDid = did.(string)
	return true, did.(string)
}


// @title    GetArticleMiddleware 获取文章列表路由中间件
// @description   参数验证 uid 可空
// @auth      	lbh             时间（2021/05/08   10:57 ）
// @page     	页码    	    int         "当前文章列的页码"
// @count     	获取条数      int         "获取文章列表条数"
// @uid     	用户id       int         "用作查询我发表的文章，可空（默认查询所有用户发表的文章）"
// @gin.context.next()    [page,count,uid]

func GetArticleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Param ArticleListValidate		//接收的数据模型
		if err := ctx.ShouldBindWith(&Param, binding.JSON); err == nil {
			ctx.Set("page",Param.Page) 		//将page写入上下文
			ctx.Set("count",Param.Count)	//将count写入上下文
			ctx.Set("uid",Param.Uid)		//将uid写入上下文
			ctx.Set("ugroup",Param.Ugroup)		//将uid写入上下文
			ctx.Next()						//进入控制器
		} else {							//数据验证不通过
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":403,"msg":"not supported"})
			ctx.Abort()
			return
		}
	}
}

func SetUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Param SetuserValidate		//接收的数据模型
		if err := ctx.ShouldBindWith(&Param, binding.JSON); err == nil {
			ctx.Set("avatarurl",Param.Avatarurl)
			ctx.Set("signature",Param.Signature)
			ctx.Set("nickname",Param.Nickname)
			ctx.Set("shippingaddress",Param.Shippingaddress)
			ctx.Set("mobile",Param.Mobile)
			ctx.Set("realname",Param.Realname)
			ctx.Set("idnumber",Param.Idnumber)
			ctx.Next()						//进入控制器
		} else {							//数据验证不通过
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":403,"msg":err.Error()})
			ctx.Abort()
			return
		}
	}
}

// @title    GetArticleInfoMiddleware 获取文章详情路由中间件
// @description   参数验证 uid 可空
// @auth      	lbh             时间（2021/05/08   10:57 ）
// @aid     	当前文章id   	int
// @gin.context.next()    	int 		aid
func GetArticleInfoMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Param ArticleInfoValidate		//接收的数据模型
		if err := ctx.ShouldBindWith(&Param, binding.JSON); err == nil {
			fmt.Println(Param)
			ctx.Set("aid",Param.Aid)
			ctx.Set("auid",Param.AUid)
			ctx.Next()						//进入控制器
		} else {							//数据验证不通过
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":403,"msg":err.Error()})
			ctx.Abort()
			return
		}
	}
}

func DoLikeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Param DoLikeArticleValidate		//接收的数据模型
		if err := ctx.ShouldBindWith(&Param, binding.JSON); err == nil {
			fmt.Println(Param)
			ctx.Set("aid",Param.Aid)
			ctx.Next()						//进入控制器
		} else {							//数据验证不通过
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":403,"msg":err.Error()})
			ctx.Abort()
			return
		}
	}
}
func DoFocusMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Param DoFocusUserValidate		//接收的数据模型
		if err := ctx.ShouldBindWith(&Param, binding.JSON); err == nil {
			fmt.Println(Param)
			ctx.Set("auid",Param.AUid)
			ctx.Next()						//进入控制器
		} else {							//数据验证不通过
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":403,"msg":err.Error()})
			ctx.Abort()
			return
		}
	}
}
func DogzarticlistMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Param GzarticlistValidate		//接收的数据模型
		if err := ctx.ShouldBindWith(&Param, binding.JSON); err == nil {
			ctx.Set("page",Param.Page) 		//将page写入上下文
			ctx.Set("count",Param.Count)	//将count写入上下文
			ctx.Next()						//进入控制器
		} else {							//数据验证不通过
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":403,"msg":err.Error()})
			ctx.Abort()
			return
		}
	}
}
func OtherUserInfoMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Param DoFocusUserValidate		//接收的数据模型
		if err := ctx.ShouldBindWith(&Param, binding.JSON); err == nil {
			fmt.Println(Param)
			ctx.Set("auid",Param.AUid)
			ctx.Next()						//进入控制器
		} else {							//数据验证不通过
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":403,"msg":err.Error()})
			ctx.Abort()
			return
		}
	}
}
//资助
func DofundingsDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Param DofundingValidate		//接收的数据模型
		if err := ctx.ShouldBindWith(&Param, binding.JSON); err == nil {
			fmt.Println(Param)
			ctx.Set("aid",Param.Aid)
			ctx.Set("htype",Param.Htype)
			ctx.Set("hdata",Param.Hdata)
			ctx.Next()						//进入控制器
		} else {							//数据验证不通过
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":403,"msg":err.Error()})
			ctx.Abort()
			return
		}
	}
}