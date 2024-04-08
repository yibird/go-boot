package sys

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"go-boot/constant"
	"go-boot/core"
	"go-boot/core/model/res"
	sys "go-boot/module/sys/user"
	"image/color"
	"time"
)

type AuthApi struct {
	Redis     *redis.Client
	Validator *core.Validator

	UserService *sys.UserService
}

func NewAuthApi(Redis *redis.Client, Validator *core.Validator, UserService *sys.UserService) *AuthApi {
	return &AuthApi{Redis, Validator, UserService}
}

func createCaptcha() *base64Captcha.Captcha {
	driver := base64Captcha.NewDriverString(50, 120, 0, 0,
		4, "1234567890qwertyuioplkjhgfdsazxcvbnm",
		&color.RGBA{0, 0, 0, 0},
		&base64Captcha.EmbeddedFontsStorage{},
		//[]string{"3Dumb.ttf", "actionj.ttf", "chromohv.ttf", "wqy-microhei.ttc", "RitaSmith.ttf"},
		[]string{},
	)
	store := base64Captcha.NewMemoryStore(10, time.Duration(1000))
	return base64Captcha.NewCaptcha(driver, store)
}

var captcha = createCaptcha()
var ctx = context.Background()
var expiration = 60 * time.Second

// GetCaptcha 获取验证码
func (s *AuthApi) GetCaptcha(c *gin.Context) {
	_, b64s, answer, err := captcha.Generate()
	if err != nil {
		res.ErrorWithMessage(fmt.Sprintf("generate captcha error,reason: %s", err.Error()), c)
		return
	}
	command := s.Redis.Set(ctx, constant.REDIS_CAPTCHA+answer, answer, expiration)
	if _, err := command.Result(); err != nil {
		res.ErrorWithMessage(fmt.Sprintf("generate captcha error,reason: %s", err.Error()), c)
		return
	}
	res.OkWithData(b64s, c)
}

// AccountLogin 账号登录
func (s *AuthApi) AccountLogin(c *gin.Context) {
	var query AccountLoginQuery
	if err := c.ShouldBind(&query); err != nil {
		return
	}
	if err := s.Validator.Validate(query); err != nil {
		res.ErrorWithMessage(err.Error(), c)
		return
	}
	captcha := s.Redis.Get(ctx, constant.REDIS_CAPTCHA+query.Captcha).Val()
	if captcha != query.Captcha {
		res.ErrorWithMessage("captcha error", c)
		return
	}
	user := s.UserService.GetUserByAccount(query.Account)
	fmt.Println("user:", user)
	if user == nil {
		res.ErrorWithMessage("user does not exist", c)
		return
	}
}

// Register 注册账号
func (s *AuthApi) Register(c *gin.Context) {

}
