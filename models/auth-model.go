package models

type SendCodeRequest struct {
	Email   string `json:"email"`
	Captcha string `json:"cf-turnstile-captcha"`
}

type VerifyCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Captcha  string `json:"cf-turnstile-captcha"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Captcha  string `json:"cf-turnstile-captcha"`
}
