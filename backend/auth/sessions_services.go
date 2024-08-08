package auth

import (
	"sync"
	"time"

	"github.com/pecet3/my-api/data"
)

const (
	typeAdmin = "admin"
	typeAuth  = "auth"
)

type Session struct {
	UserId       int
	Expiry       time.Time
	Token        string
	ActivateCode string
	UserIp       string
	Type         string
}

type SessionStore struct {
	AuthSessions  AuthSessions
	sMu           sync.RWMutex
	AdminSessions AdminSessions
	eMu           sync.RWMutex
	Password      string
	pMu           sync.RWMutex
	data          data.Data
}

func NewSessionStore(d data.Data) *SessionStore {
	return &SessionStore{
		AuthSessions:  make(AuthSessions),
		AdminSessions: make(AdminSessions),
		data:          d,
		Password:      generatePassword(),
	}
}
