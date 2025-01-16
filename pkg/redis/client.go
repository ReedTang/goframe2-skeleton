package redis

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func init() {
	ctx := gctx.New()
	Client = redis.NewClient(&redis.Options{
		Addr:     g.Cfg().MustGet(ctx, "redis.default.address").String(),
		Password: g.Cfg().MustGet(ctx, "redis.default.pass").String(),
		DB:       g.Cfg().MustGet(ctx, "redis.default.db").Int(),
	})
}
