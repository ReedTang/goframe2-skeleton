package simple

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// Debug debug
func Debug(ctx context.Context) bool {
	return g.Cfg().MustGet(ctx, "debug", true).Bool()
}

// IsCluster 是否为集群部署
func IsCluster(ctx context.Context) bool {
	return g.Cfg().MustGet(ctx, "system.isCluster", true).Bool()
}
