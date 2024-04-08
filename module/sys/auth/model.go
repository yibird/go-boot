package sys

type AccountLoginQuery struct {
	Account  string `form:"account" json:"account" validate:"required" error:"账号不能为空"`
	Password string `form:"password" json:"password" validate:"required,min=6,max=20" error:"密码不能为空"`
	Captcha  string `form:"captcha" json:"captcha" validate:"required,len=4" error:"验证码不能为空|len=非法长度"`
}
