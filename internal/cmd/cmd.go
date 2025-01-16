package cmd

import (
	"context"

	_ "goframe2-skeleton/pkg/redis"
	"goframe2-skeleton/pkg/simple"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	Main = &gcmd.Command{
		Description: "默认启动全部服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return All.Func(ctx, parser)
		},
	}

	All = &gcmd.Command{
		Name:        "all",
		Brief:       "start all server",
		Description: "this is the command entry for starting all server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, "starting all server")
			// 需要启动的服务
			_ = gtime.SetTimeZone("Asia/Shanghai")
			allServers := []*gcmd.Command{Http}
			for _, server := range allServers {
				cmd := server
				simple.SafeGo(ctx, func(ctx context.Context) {
					if err := cmd.Func(ctx, parser); err != nil {
						g.Log().Fatalf(ctx, "%v start fail:%v", cmd.Name, err)
					}
				})
			}
			// 信号监听
			signalListen(ctx, signalHandlerForOverall)
			<-serverCloseSignal
			serverWg.Wait()
			g.Log().Debug(ctx, "all service successfully closed ..")
			return
		},
	}
)

func init() {
	if err := Main.AddCommand(All, Http); err != nil {
		panic(err)
	}
}
