// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"goframe2-skeleton/api/auth/v1"
)

type IAuthV1 interface {
	LoginDo(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error)
	LoginLogout(ctx context.Context, req *v1.LoginLogoutReq) (res *v1.LoginLogoutRes, err error)
}
