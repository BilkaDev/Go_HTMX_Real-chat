package config

type ContextKey string

const (
	CurrentUserName ContextKey = "currentUserName"
	CurrentUserId   ContextKey = "currentUserId"
	EnvKey          ContextKey = "env"
)

func (c ContextKey) String() string {
	return string(c)
}
