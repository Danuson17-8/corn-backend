package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func VerifyTurnstile(token string, remoteIP string) (bool, error) {

	// Dev mode
	if token == "XXXX.DUMMY.TOKEN.XXXX" {
		return true, nil
	}

	secret := os.Getenv("CF_TURNSTILE_SECRET")
	if secret == "" {
		return false, errors.New("missing turnstile secret")
	}

	data := url.Values{}
	data.Set("secret", secret)
	data.Set("response", token)
	if remoteIP != "" {
		data.Set("remoteip", remoteIP)
	}

	resp, err := http.PostForm(
		"https://challenges.cloudflare.com/turnstile/v0/siteverify",
		data,
	)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("turnstile http status %d", resp.StatusCode)
	}

	var result struct {
		Success     bool     `json:"success"`
		ErrorCodes  []string `json:"error-codes"`
		Hostname    string   `json:"hostname"`
		ChallengeTS string   `json:"challenge_ts"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if !result.Success {
		return false, fmt.Errorf("turnstile failed: %v", result.ErrorCodes)
	}

	return true, nil
}
