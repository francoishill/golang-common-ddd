package DefaultEncryptionService

import (
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Encryption"
)

type defaultEncryption struct{}

func (d *defaultEncryption) PasswordMatchHex(guessedPassword, hexSalt, hexHashedPassword string) bool {
	return PasswordMatchHex(guessedPassword, hexSalt, hexHashedPassword)
}

func (d *defaultEncryption) CreatePassword(raw_pass string) Password {
	return CreatePassword_DefaultComplexity(raw_pass)
}

func New() EncryptionService {
	return &defaultEncryption{}
}
