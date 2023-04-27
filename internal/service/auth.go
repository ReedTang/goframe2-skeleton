package service

import (
	"context"
	v1 "goframe2-skeleton/api/v1"
	"goframe2-skeleton/internal/consts"
	"goframe2-skeleton/internal/model/entity"
	"goframe2-skeleton/utility/response"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	env, _ := g.Cfg().Get(context.TODO(), "server.env")
	secretKey, _ := g.Cfg().Get(context.TODO(), "jwt.secret_key")
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           secretKey.String(),
		Key:             []byte(secretKey.String() + "_" + env.String()),
		Timeout:         5 * time.Hour * 24,
		MaxRefresh:      15 * time.Hour * 24,
		IdentityKey:     "loginUserId",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	messageMap := g.Map{
		"token is expired":         "登录已过期",
		"auth header is empty":     "登录已过期",
		"cookie token is empty":    "登录已过期",
		"parameter token is empty": "登录已过期",
		"token is invalid":         "登录已失效",
		"signature is invalid":     "登录已失效",
	}
	re, ok := messageMap[message]
	if ok {
		message = gconv.String(re)
	}
	response.JsonExit(r, consts.ErrorNoAuthCode, message, nil)
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (data interface{}, err error) {
	var (
		r   = g.RequestFromCtx(ctx)
		req *v1.LoginDoReq
	)
	if err = r.Parse(&req); err != nil {
		return "", err
	}

	var user *entity.AiUser
	if user, err = User().Login(ctx, req); err == nil {
		if user.Id > 0 {
			if user.Status != 1 {
				return "", gerror.New("账号已被禁用")
			}
			return g.Map{"loginUserId": user.Id, "loginUserName": user.UserName}, nil
		}
	} else {
		return "", err
	}

	return nil, jwt.ErrFailedAuthentication
}
