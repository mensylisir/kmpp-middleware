package entity

import "github.com/mensylisir/kmpp-middleware/src/model"

type User struct {
	model.User
	Roles []string `json:"roles"`
}

type Profile struct {
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginProfile struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CaptchaId string `json:"captcha_id"`
	Code      string `json:"code"`
}

type RfreshProfile struct {
	RefreshToken string `json:"refresh_token"`
}

type Captcha struct {
	Image     string `json:"image"`
	CaptchaId string `json:"captcha_id"`
}

type OperateUser struct {
	Operation string       `json:"operation"`
	Items     []model.User `json:"items"`
}

type UserChangePassword struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Original string `json:"original"`
}

type UserPasswordForget struct {
	ID string `json:"id"`
}
