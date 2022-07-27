package dbv1_0_0_admin

import (
	stdlog "log"
	"os"
	"testing"

	"github.com/ikaiguang/go-srv-admin/internal/setup"
)

// upHandler handler
var upHandler *migrate

func TestMain(m *testing.M) {
	// 初始化
	configPath := "./../../../../configs"
	if err := setup.Init(setup.WithConfigPath(configPath)); err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	//defer func() { _ = setup.Close() }()

	// 模块
	engineHandler, err := setup.GetEngine()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	// 关闭
	//defer func() { _ = setup.Close() }()

	// 数据库链接
	dbConn, err := engineHandler.GetMySQLGormDB()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	// handler
	upHandler = NewMigrateHandler(dbConn)

	os.Exit(m.Run())
}
