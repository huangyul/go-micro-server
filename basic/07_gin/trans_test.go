package _7_gin

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type SignUpForm struct {
	Age        uint8  `json:"age" binding:"required,gte=1,lte=130"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

var tran ut.Translator

func removeTopStruct(fil map[string]string) map[string]string {
	rsp := map[string]string{}
	for key, v := range fil {
		rsp[key[strings.Index(key, ".")+1:]] = v
	}
	return rsp
}

func InitTrans(locale string) error {
	// 修改gin的validator引擎属性
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册一个获取json的tag的方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			// json约定这种情况会忽略
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译
		enT := en.New() // 英文翻译
		uni := ut.New(enT, zhT, enT)
		tran, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("get translator: %s", locale)
		}

		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, tran)
			break
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, tran)
			break
		default:
			_ = en_translations.RegisterDefaultTranslations(v, tran)
		}

	}
	return nil
}

func TestTrans(t *testing.T) {
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译错误")
	}
	g := gin.Default()
	g.POST("/user", func(ctx *gin.Context) {
		var singReq SignUpForm
		if err := ctx.ShouldBind(&singReq); err != nil {
			var errs validator.ValidationErrors
			ok := errors.As(err, &errs)
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{
					"error": errs.Error(),
				})
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": removeTopStruct(errs.Translate(tran)),
			})
			return
		}
		ctx.String(http.StatusOK, "注册成功")
	})

	_ = g.Run(":8088")
}
