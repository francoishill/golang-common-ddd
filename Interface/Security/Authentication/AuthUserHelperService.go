package Authentication

type AuthUserHelperService interface {
	VerifyAndGetUserFromCredentials(email, username, password string) AuthUser
	GetUserWithUUID(uid interface{}) AuthUser
	RegisterUser(email, username, password string) AuthUser
}