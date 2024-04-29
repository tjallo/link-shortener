package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"io"
	"log"

	"github.com/tjalle/link_shortener/initializers"
	"github.com/tjalle/link_shortener/models"
	"golang.org/x/crypto/argon2"
)

const (
	SALT_BYTES      = 32
	HASH_BYTES      = 64
	HASHING_TIME    = 10
	HASHING_MEMORY  = HASH_BYTES * 1024
	HASHING_THREADS = 1
)

func hashPassword(password []byte, salt []byte) []byte {
	return argon2.Key(password, salt, HASHING_TIME, HASHING_MEMORY, HASHING_THREADS, HASH_BYTES)
}

func generateRandomSalt() []byte {
	salt := make([]byte, SALT_BYTES)

	_, err := io.ReadFull(rand.Reader, salt)

	if err != nil {
		log.Fatal(err)
	}

	return salt
}

func CreateUser(username string, password string) error {
	salt := generateRandomSalt()
	pwdHash := hashPassword([]byte(password), salt)

	user := models.User{
		Username:     username,
		PasswordHash: pwdHash,
		Salt:         salt,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return fmt.Errorf("error creating new user")
	}

	return nil
}

func VerifyUser(username string, password string) error {
	var user = models.User{Username: username}

	result := initializers.DB.First(&user)

	if result.Error != nil {
		return fmt.Errorf("user not found")
	}

	hashedPwd := hashPassword([]byte(password), user.Salt)

	verified := subtle.ConstantTimeCompare(hashedPwd, user.PasswordHash)

	if verified != 1 {
		return fmt.Errorf("invalid username or password")
	}

	return nil
}
