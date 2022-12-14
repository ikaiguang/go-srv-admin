package dbv1_0_0_admin

import (
	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	"gorm.io/gorm"

	schemas "github.com/ikaiguang/go-srv-admin/internal/infra/mysql/schema/admin"
)

// Upgrade .
func Upgrade(dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) (err error) {
	upgradeHandler := NewMigrateHandler(dbConn, migrateRepo)

	// CreateTableAdmin 创建表 admin
	err = upgradeHandler.CreateTableAdmin()
	if err != nil {
		return err
	}
	// CreateTableAdminRegEmail 创建表 admin
	err = upgradeHandler.CreateTableAdminRegEmail()
	if err != nil {
		return err
	}
	// CreateTableAdminRegMobile 创建表 admin
	err = upgradeHandler.CreateTableAdminRegMobile()
	if err != nil {
		return err
	}
	// CreateTableAdminRegUsername 创建表 admin
	err = upgradeHandler.CreateTableAdminRegUsername()
	if err != nil {
		return err
	}

	// InitializeAdmin 初始化管理员
	err = upgradeHandler.InitializeAdmin()
	if err != nil {
		return err
	}
	return err
}

// migrate 数据库迁移
type migrate struct {
	dbConn      *gorm.DB
	migrateRepo migrationutil.MigrateRepo
}

// NewMigrateHandler 处理手柄
func NewMigrateHandler(dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) *migrate {
	return &migrate{
		dbConn:      dbConn,
		migrateRepo: migrateRepo,
	}
}

// CreateTableAdmin 创建表 admin
func (s *migrate) CreateTableAdmin() (err error) {
	if s.dbConn.Migrator().HasTable(schemas.AdminSchema) {
		return err
	}
	return s.dbConn.Migrator().CreateTable(schemas.AdminSchema)
}

// CreateTableAdminRegEmail 创建表 admin
func (s *migrate) CreateTableAdminRegEmail() (err error) {
	if s.dbConn.Migrator().HasTable(schemas.AdminRegEmailSchema) {
		return err
	}
	return s.dbConn.Migrator().CreateTable(schemas.AdminRegEmailSchema)
}

// CreateTableAdminRegMobile 创建表 admin
func (s *migrate) CreateTableAdminRegMobile() (err error) {
	if s.dbConn.Migrator().HasTable(schemas.AdminRegMobileSchema) {
		return err
	}
	return s.dbConn.Migrator().CreateTable(schemas.AdminRegMobileSchema)
}

// CreateTableAdminRegUsername 创建表 admin
func (s *migrate) CreateTableAdminRegUsername() (err error) {
	if s.dbConn.Migrator().HasTable(schemas.AdminRegUsernameSchema) {
		return err
	}
	return s.dbConn.Migrator().CreateTable(schemas.AdminRegUsernameSchema)
}

// InitializeAdmin 初始化管理员
func (s *migrate) InitializeAdmin() (err error) {
	var (
		serverVersion     = "v1.0.0"
		migrateIdentifier = serverVersion + ":InitializeAdmin"
	)
	// 已进行数据库迁移
	dataModel, _, err := s.migrateRepo.QueryOneByIdentifier(migrateIdentifier)
	if err != nil {
		return
	}
	if dataModel.Id > 0 {
		return
	}
	// 记录数据库迁移
	defer func() {
		if err == nil {
			err = s.migrateRepo.CreateDefaultRecord(serverVersion, migrateIdentifier)
		}
	}()

	err = schemas.AdminSchema.Initialize(s.dbConn)
	if err != nil {
		return err
	}
	return
}
