package repositories

type SessionRepository struct {
	Sessions map[string]string
}

func (r *SessionRepository) GetAccount(sessionID string) (string, bool) {
	if r.Sessions == nil {
		return "", false
	}
	account, ok := r.Sessions[sessionID]
	return account, ok
}
