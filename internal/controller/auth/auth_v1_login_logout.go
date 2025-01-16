package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"goframe2-skeleton/api/auth/v1"
)

func (c *ControllerV1) LoginLogout(ctx context.Context, req *v1.LoginLogoutReq) (res *v1.LoginLogoutRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
