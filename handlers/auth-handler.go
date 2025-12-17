package handlers

import (
	"github.com/Danuson17-8/corn-backend/config"
	"github.com/Danuson17-8/corn-backend/models"
	"github.com/Danuson17-8/corn-backend/services"
	"github.com/Danuson17-8/corn-backend/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Auth *services.AuthService
	OTP  *services.OTPService
	JWT  *services.JWTService
}

func (h *AuthHandler) SendCode(c *fiber.Ctx) error {
	var body models.SendCodeRequest
	if err := c.BodyParser(&body); err != nil {
		return Error(c, 400, "Invalid JSON")
	}

	if body.Captcha == "" {
		return Error(c, 400, "Captcha required")
	}

	ip := c.Get("CF-Connecting-IP")
	ok, err := utils.VerifyTurnstile(body.Captcha, ip)
	if err != nil || !ok {
		return Error(c, 403, "Captcha verification failed")
	}
	cfg := config.NewEnvConfig()
	otpSession, err := h.OTP.SendOTP(cfg, body.Email)
	if err != nil {
		return Error(c, 500, "Cannot send OTP")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "otp_session",
		Value:    otpSession,
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteStrictMode,
		MaxAge:   300,
		Path:     "/",
	})

	return Success(c, "OTP sent", fiber.Map{
		"email": body.Email,
	})
}

func (h *AuthHandler) VerifyCode(c *fiber.Ctx) error {
	var body models.VerifyCodeRequest
	if err := c.BodyParser(&body); err != nil {
		return Error(c, 400, "Invalid JSON")
	}

	otpSession := c.Cookies("otp_session")
	if otpSession == "" {
		return Error(c, 401, "OTP session missing")
	}

	authSession, err := h.OTP.VerifyOTP(body.Email, body.Code, otpSession)
	if err != nil {
		switch err {
		case services.ErrOTPExpired:
			return Error(c, 400, "Code expired")
		case services.ErrOTPInvalid:
			return Error(c, 400, "Invalid code")
		default:
			return Error(c, 500, "Verification failed")
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth_session",
		Value:    authSession,
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteLaxMode,
		MaxAge:   86400,
		Path:     "/",
	})

	c.ClearCookie("otp_session")

	return Success(c, "Verified!")
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var body models.RegisterRequest
	if err := c.BodyParser(&body); err != nil {
		return Error(c, 400, "Invalid JSON")
	}

	if body.Captcha == "" {
		return Error(c, 400, "Captcha required")
	}

	ok, err := utils.VerifyTurnstile(body.Captcha, c.IP())
	if err != nil || !ok {
		return Error(c, 403, "Captcha verification failed")
	}

	if err := h.Auth.Register(body.Email, body.Password); err != nil {
		if err == services.ErrEmailExists {
			return Error(c, 409, "Email already registered")
		}
		return Error(c, 500, "Register failed")
	}

	return Success(c, "User registered successfully")
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body models.LoginRequest
	if err := c.BodyParser(&body); err != nil {
		return Error(c, 400, "Invalid JSON")
	}

	ok, err := utils.VerifyTurnstile(body.Captcha, c.IP())
	if err != nil || !ok {
		return Error(c, 403, "Captcha verification failed")
	}

	if err := h.Auth.Login(body.Email, body.Password); err != nil {
		return Error(c, 401, "Invalid email or password")
	}

	token, err := h.JWT.Generate(body.Email, "user")
	if err != nil {
		return Error(c, 500, "Token generation failed")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteLaxMode,
		MaxAge:   86400,
		Path:     "/",
	})

	return Success(c, "Login success")
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HTTPOnly: true,
	})

	return Success(c, "Logged out")
}
