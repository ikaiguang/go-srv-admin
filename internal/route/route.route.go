package routes

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	adminroute "github.com/ikaiguang/go-srv-admin/internal/route/admin"
	pingroute "github.com/ikaiguang/go-srv-admin/internal/route/ping"
	testdataroute "github.com/ikaiguang/go-srv-admin/internal/route/testdata"
	"github.com/ikaiguang/go-srv-admin/internal/setup"
)

// RegisterRoutes 注册路由
func RegisterRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) (err error) {
	stdlog.Println("|*** 注册路由：...")

	// 日志
	logger, _, err := engineHandler.Logger()
	if err != nil {
		return err
	}

	// testdata
	pingroute.RegisterRoutes(hs, gs, logger)
	testdataroute.RegisterRoutes(hs, gs, logger)

	// admin
	if err = adminroute.RegisterRoutes(engineHandler, hs, gs); err != nil {
		return err
	}

	return err
}
