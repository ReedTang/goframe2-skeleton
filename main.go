package main

import (
	"goframe2-skeleton/internal/cmd"
	"goframe2-skeleton/internal/global"
	_ "goframe2-skeleton/internal/logic"
	_ "goframe2-skeleton/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func main() {
	_ = gtime.SetTimeZone("Asia/Shanghai")
	global.Init(gctx.GetInitCtx())
	cmd.Main.Run(gctx.GetInitCtx())
}
