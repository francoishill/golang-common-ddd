package DefaultEncryptionService

import (
	"crypto/rand"

	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Encryption"
)

type defaultEncryption struct{}

func (d *defaultEncryption) PasswordMatchHex(guessedPassword, hexSalt, hexHashedPassword string) bool {
	return PasswordMatchHex(guessedPassword, hexSalt, hexHashedPassword)
}

func (d *defaultEncryption) CreatePassword(raw_pass string) Password {
	return CreatePassword_DefaultComplexity(raw_pass)
}

func (d *defaultEncryption) generateRandomString(length int, dictionary string) string {
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func (d *defaultEncryption) GenerateRandomTokenString_AlphaNumeric(length int) string {
	return d.generateRandomString(length, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
}

func (d *defaultEncryption) GenerateRandomTokenString_Alpha(length int) string {
	return d.generateRandomString(length, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
}

func (d *defaultEncryption) GenerateRandomTokenString_Numeric(length int) string {
	return d.generateRandomString(length, "0123456789")
}

func (d *defaultEncryption) GenerateRandomTokenString_SpecifyCharlist(length int, charlist string) string {
	return d.generateRandomString(length, charlist)
}

func New() EncryptionService {
	return &defaultEncryption{}
}
