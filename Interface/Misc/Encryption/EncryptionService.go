package Encryption

type Password interface {
	GetHashedPasswordHex() string
	GetSaltHex() string
}

type EncryptionService interface {
	PasswordMatchHex(guessedPassword, hexSalt, hexHashedPassword string) bool
	CreatePassword(raw_pass string) Password
}
