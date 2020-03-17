package routes

import (
	"github.com/alphayan/iris"
)

// GetFollowingHandler handles the GET: /following/{id}
func GetFollowingHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	ctx.Writef("from "+ctx.GetCurrentRoute().Path()+" with ID: %d", id)
}
