package servers

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/golang-jwt/jwt/v4"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"

	"github.com/ikaiguang/go-srv-admin/internal/setup"
)

// NewWhiteListMatcher 路由白名单
func NewWhiteListMatcher() selector.MatchFunc {

	// 白名单
	whiteList := make(map[string]bool)

	// 测试
	whiteList["/kit.api.ping.pingv1.SrvPing/Ping"] = true

	// 后台
	//whiteList["/admin/v1/admin-auth/login-by-email"] = true
	whiteList["/admin.api.admin.adminservicev1.SrvAdminAuth/LoginByEmail"] = true

	return func(ctx context.Context, operation string) bool {
		//if tr, ok := contextutil.MatchHTTPServerContext(ctx); ok {
		//	if _, ok := whiteList[tr.Request().URL.Path]; ok {
		//		return false
		//	}
		//}

		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewJWTMiddleware jwt中间
func NewJWTMiddleware(engineHandler setup.Engine) (m middleware.Middleware, err error) {
	redisCC, err := engineHandler.GetRedisClient()
	if err != nil {
		return m, errorutil.WithStack(err)
	}
	authTokenRepo := engineHandler.GetAuthTokenRepo(redisCC)

	m = selector.Server(
		authutil.Server(
			authTokenRepo.JWTKeyFunc(),
			authutil.WithSigningMethod(authTokenRepo.JWTSigningMethod()),
			authutil.WithClaims(func() jwt.Claims { return &authutil.Claims{} }),
		),
	).
		Match(NewWhiteListMatcher()).
		Build()

	return m, err
}
