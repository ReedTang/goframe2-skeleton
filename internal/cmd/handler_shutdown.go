package cmd

import (
	"context"
	"os"
	"sync"

	"goframe2-skeleton/pkg/simple"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
)

var (
	serverCloseSignal = make(chan struct{}, 1)
	serverWg          = sync.WaitGroup{}
	once              sync.Once
)

// signalHandlerForOverall 关闭信号处理
func signalHandlerForOverall(sig os.Signal) {
	serverCloseSignal <- struct{}{}
	serverCloseEvent(gctx.GetInitCtx())
}

// signalListen 信号监听
func signalListen(ctx context.Context, handler ...gproc.SigHandler) {
	simple.SafeGo(ctx, func(ctx context.Context) {
		gproc.AddSigHandlerShutdown(handler...)
		gproc.Listen()
	})
}

// serverCloseEvent 关闭事件
// 区别于服务收到退出信号后的处理，只会执行一次
func serverCloseEvent(ctx context.Context) {
	// once.Do(func() {
	// 	simple.Event().Call(consts.EventServerClose, ctx)
	// })
}
