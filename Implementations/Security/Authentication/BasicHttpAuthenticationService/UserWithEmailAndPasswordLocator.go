package BasicHttpAuthenticationService

type UserWithEmailAndPasswordLocator interface {
	FindUserWithEmailAndPassword(email, password string) interface{}
}
