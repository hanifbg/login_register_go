package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pass string) (hashedPassword []byte, err error) {
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}
