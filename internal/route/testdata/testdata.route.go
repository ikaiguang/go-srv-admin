package testdataroute

import (
	testdataservicev1 "github.com/ikaiguang/go-srv-kit/api/testdata/v1/services"
	stdlog "log"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	testdatasrv "github.com/ikaiguang/go-srv-admin/internal/application/service/testdata"
)

// RegisterRoutes 注册路由
func RegisterRoutes(hs *http.Server, gs *grpc.Server, logger log.Logger) {

	testdata := testdatasrv.NewTestdataService(logger)
	stdlog.Println("|*** 注册路由：NewTestdataService")
	testdataservicev1.RegisterSrvTestdataHTTPServer(hs, testdata)
	testdataservicev1.RegisterSrvTestdataServer(gs, testdata)
}
