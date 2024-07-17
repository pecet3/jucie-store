package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"regexp"
)

// todo: rename to credentials
const saltSize = 32

type password struct {
}
type passwordServices interface {
	hashPassword(salt, password string) (string, error)
	generateSalt() (string, error)
	verifyPassword(salt, password string) (bool, error)
	validatePassword(password string) (bool, error)
}

func (p password) validatePassword(password string) (bool, error) {
	var (
		hasUpperCase   = regexp.MustCompile(`[A-Z]`).MatchString
		hasLowerCase   = regexp.MustCompile(`[a-z]`).MatchString
		hasNumber      = regexp.MustCompile(`[0-9]`).MatchString
		hasSpecialChar = regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString
	)

	if !hasUpperCase(password) {
		return false, errors.New("the password must contain at least one uppercase letter")
	}
	if !hasLowerCase(password) {
		return false, errors.New("the password must contain at least one lowercase letter")
	}
	if !hasNumber(password) {
		return false, errors.New("the password must contain at least one number")
	}
	if !hasSpecialChar(password) {
		return false, errors.New("the password must contain at least one special character")
	}

	return true, nil
}

func (p password) hashPassword(salt, password string) (string, error) {
	hasher := sha256.New()

	_, err := hasher.Write([]byte(salt + password))
	if err != nil {
		log.Println(err)
		return "", err
	}
	hashInBytes := hasher.Sum(nil)

	hashString := hex.EncodeToString(hashInBytes)

	return hashString, nil
}

func (p password) generateSalt() (string, error) {
	bytes := make([]byte, saltSize)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println(err)

		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (p password) verifyPassword(salt, password string) (bool, error) {
	hashedPassword, err := p.hashPassword(salt, password)
	if err != nil {
		return false, err
	}
	if hashedPassword != password {
		return true, nil
	}

	return false, nil
}
