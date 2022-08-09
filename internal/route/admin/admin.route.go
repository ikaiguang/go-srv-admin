package adminroute

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdlog "log"

	adminservicev1 "github.com/ikaiguang/go-srv-admin/api/admin/v1/services"
	adminsrv "github.com/ikaiguang/go-srv-admin/internal/application/service/admin"
	datas "github.com/ikaiguang/go-srv-admin/internal/infra/mysql/data/admin"
	"github.com/ikaiguang/go-srv-admin/internal/setup"
)

// RegisterRoutes 注册路由
func RegisterRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) (err error) {
	// 数据库
	dbConn, err := engineHandler.GetMySQLGormDB()
	if err != nil {
		return err
	}
	redisCC, err := engineHandler.GetRedisClient()
	if err != nil {
		return err
	}
	authTokenRepo := engineHandler.GetAuthTokenRepo(redisCC)

	// admin
	adminRepo := datas.NewAdminRepo(dbConn)
	adminRegEmailRepo := datas.NewAdminRegEmailRepo(dbConn)

	// oauth 授权
	adminAuthSrv := adminsrv.NewAdminAuthService(
		authTokenRepo,
		adminRepo,
		adminRegEmailRepo,
	)
	stdlog.Println("|*** 注册路由：NewAdminAuthService")
	adminservicev1.RegisterSrvAdminAuthHTTPServer(hs, adminAuthSrv)
	adminservicev1.RegisterSrvAdminAuthServer(gs, adminAuthSrv)

	return err
}
