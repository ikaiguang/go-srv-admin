package servers

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/golang-jwt/jwt/v4"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"

	"github.com/ikaiguang/go-srv-admin/internal/setup"
)

// NewWhiteListMatcher 路由白名单
func NewWhiteListMatcher() selector.MatchFunc {

	// 白名单
	whiteList := make(map[string]bool)

	// 测试
	whiteList["/kit.api.ping.pingv1.SrvPing/Ping"] = true

	return func(ctx context.Context, operation string) bool {
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
		return m, err
	}
	authTokenRepo := engineHandler.GetAuthTokenRepo(redisCC)

	m = selector.Server(
		authutil.Server(
			authTokenRepo.JWTKeyFunc,
			authutil.WithSigningMethod(authTokenRepo.JWTSigningMethod()),
			authutil.WithClaims(func() jwt.Claims { return &authutil.Claims{} }),
		),
	).
		Match(NewWhiteListMatcher()).
		Build()

	return m, err
}
