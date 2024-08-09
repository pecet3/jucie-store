package auth

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func changePasswordLoop(as *SessionStore) {
	for {
		err := as.CreateNewPassword()
		if err != nil {
			log.Panicln(err)
			break
		}
		log.Println("Changed a password: ", as.Password)
		time.Sleep(1 * time.Minute)
	}
}

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
