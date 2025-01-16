package auth

import (
	"context"

	v1 "goframe2-skeleton/api/auth/v1"
)

func (c *ControllerV1) LoginDo(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
	res = &v1.LoginDoRes{}
	// loginInfo, err := service.Auth().Login(ctx, req)
	// if err != nil {
	// 	return nil, err
	// }
	// err = gconv.Scan(loginInfo, &res)
	// if err != nil {
	// 	return nil, err
	// }
	return
}
