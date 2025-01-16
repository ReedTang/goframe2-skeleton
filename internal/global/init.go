package global

import (
	"context"

	"goframe2-skeleton/internal/library/cache"
	"goframe2-skeleton/internal/service"
	"goframe2-skeleton/utility/validate"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmode"
)

func Init(ctx context.Context) {
	// 设置gf运行模式
	SetGFMode(ctx)

	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常 err：%+v", err)
		return
	}

	// fmt.Printf("欢迎使用！\r\n当前运行环境：%v, 运行根路径为：%v \r\nApp版本：v%v, gf版本：%v \n", runtime.GOOS, gfile.Pwd(), consts.VersionApp, gf.VERSION)

	// 设置缓存适配器
	cache.SetAdapter(ctx)

	// 初始化功能库配置
	service.SysConfig().InitConfig(ctx)
}

// SetGFMode 设置gf运行模式
func SetGFMode(ctx context.Context) {
	mode := g.Cfg().MustGet(ctx, "system.mode").String()
	if len(mode) == 0 {
		mode = gmode.NOT_SET
	}

	modes := []string{gmode.DEVELOP, gmode.TESTING, gmode.STAGING, gmode.PRODUCT}

	// 如果是有效的运行模式，就进行设置
	if validate.InSlice(modes, mode) {
		gmode.Set(mode)
	}
}
