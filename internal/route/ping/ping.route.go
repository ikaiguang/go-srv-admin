package pingroute

import (
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
	stdlog "log"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	pingsrv "github.com/ikaiguang/go-srv-admin/internal/application/service/ping"
)

// RegisterRoutes 注册路由
func RegisterRoutes(hs *http.Server, gs *grpc.Server, logger log.Logger) {

	ping := pingsrv.NewPingService(logger)
	stdlog.Println("|*** 注册路由：NewPingService")
	pingservicev1.RegisterSrvPingHTTPServer(hs, ping)
	pingservicev1.RegisterSrvPingServer(gs, ping)
}
