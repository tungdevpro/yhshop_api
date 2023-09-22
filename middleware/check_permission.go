package middleware

import (
	"coffee_api/commons"

	"github.com/gin-gonic/gin"
)

func CheckPermission(ctx *gin.Context, role *commons.RoleAllowed) bool {
	if role == nil {
		return false
	}

	r := ctx.MustGet(commons.CurrentUser).(commons.Requester)
	return r.GetRole() == string(*role)
}
