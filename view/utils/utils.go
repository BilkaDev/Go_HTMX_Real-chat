package utils

import (
	"context"

	"github.com/bilkadev/Go_HTMX_Real-chat/config"
)

func GetUserName(c context.Context) string {
	user, ok := c.Value(config.CurrentUserName).(string)
	if !ok {

		return ""
	}
	return user
}
