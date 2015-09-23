package Authentication

type BaseAuthUserHelperService interface {
	BaseVerifyAndGetUserFromCredentials(email, username, password string) BaseUser
	BaseGetUserWithUUID(uid interface{}) BaseUser
	BaseRegisterUser(email, username, password string) BaseUser
}
