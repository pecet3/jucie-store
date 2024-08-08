package auth

import (
	"math/rand"
	"strconv"
	"strings"
)

func generatePassword() string {
	randomNumber := rand.Intn(9000) + 1000
	numberString := strconv.Itoa(randomNumber)

	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	firstLetter := string(letters[rand.Intn(len(letters))])

	return firstLetter + strings.ToLower(numberString)
}

func (as *SessionStore) GetCurrentPassword() string {
	as.pMu.RLock()
	defer as.pMu.RUnlock()
	return as.Password
}

func (as *SessionStore) CreateNewPassword() error {
	as.pMu.Lock()
	defer as.pMu.Unlock()

	as.Password = generatePassword()

	return nil
}
