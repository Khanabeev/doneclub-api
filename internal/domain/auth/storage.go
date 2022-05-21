package auth

type Storage interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}
