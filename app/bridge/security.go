package bridge

import (
	"github.com/webability-go/xamboo/cms/context"
)

func GetUserKey(ctx *context.Context) int {
	userkey, _ := ctx.Sessionparams.GetInt("userkey")
	return userkey
}
