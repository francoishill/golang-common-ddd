package Encryption

type Password interface {
	GetHashedPasswordHex() string
	GetSaltHex() string
}

type EncryptionService interface {
	PasswordMatchHex(guessedPassword, hexSalt, hexHashedPassword string) bool
	CreatePassword(raw_pass string) Password

	GenerateRandomTokenString_AlphaNumeric(length int) string
	GenerateRandomTokenString_Alpha(length int) string
	GenerateRandomTokenString_Numeric(length int) string
	GenerateRandomTokenString_SpecifyCharlist(length int, charlist string) string
}
