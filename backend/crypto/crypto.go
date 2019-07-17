package crypto

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePwds(pwdHash string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
