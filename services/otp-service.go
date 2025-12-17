package services

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Danuson17-8/corn-backend/config"
	"github.com/Danuson17-8/corn-backend/repositories"
	"github.com/Danuson17-8/corn-backend/utils"
)

type OTPService struct {
	Repo *repositories.OTPRepository
}

func (s *OTPService) SendOTP(cfg *config.EnvConfig, email string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	expired := time.Now().Add(5 * time.Minute)
	session := utils.RandomToken(32)

	if err := s.Repo.Upsert(email, code, session, expired); err != nil {
		return "", err
	}

	if err := utils.SendEmail(cfg, email, "Corn Corn", fmt.Sprintf("OTP is: %d", code)); err != nil {
		return "", err
	}

	return session, nil
}

var (
	ErrOTPExpired = errors.New("otp expired")
	ErrOTPInvalid = errors.New("invalid otp")
)

func (s *OTPService) VerifyOTP(email, code, session string) (string, error) {
	dbCode, dbSession, expired, err := s.Repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if time.Now().After(expired) {
		return "", ErrOTPExpired
	}

	if dbCode != code || dbSession != session {
		return "", ErrOTPInvalid
	}

	authSession := utils.RandomToken(48)
	return authSession, nil
}
