// Package datas
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package datas

import (
	"bytes"
	context "context"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
	gorm "gorm.io/gorm"
	"strings"

	entities "github.com/ikaiguang/go-srv-admin/internal/domain/entity/admin"
	repos "github.com/ikaiguang/go-srv-admin/internal/domain/repo/admin"
	schemas "github.com/ikaiguang/go-srv-admin/internal/infra/mysql/schema/admin"
)

// adminRepo repo
type adminRepo struct {
	dbConn      *gorm.DB      // *gorm.DB
	AdminSchema schemas.Admin // Admin
}

// NewAdminRepo new data repo
func NewAdminRepo(dbConn *gorm.DB) repos.AdminRepo {
	return &adminRepo{
		dbConn: dbConn,
	}
}

// =============== 创建 ===============

// create insert one
func (s *adminRepo) create(ctx context.Context, dbConn *gorm.DB, dataModel *entities.Admin) (err error) {
	err = dbConn.WithContext(ctx).Create(dataModel).Error
	if err != nil {
		return err
	}
	return
}

// Create insert one
func (s *adminRepo) Create(ctx context.Context, dataModel *entities.Admin) error {
	return s.create(ctx, s.dbConn, dataModel)
}

// CreateWithDBConn create
func (s *adminRepo) CreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *entities.Admin) error {
	return s.create(ctx, dbConn, dataModel)
}

// existCreate exist create
func (s *adminRepo) existCreate(ctx context.Context, dbConn *gorm.DB, dataModel *entities.Admin) (anotherModel *entities.Admin, isNotFound bool, err error) {
	anotherModel = new(entities.Admin)
	err = dbConn.WithContext(ctx).
		Table(s.AdminSchema.TableName()).
		Where("id = ?", dataModel.Id).
		First(anotherModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			isNotFound = true
			err = nil
		}
		return
	}
	return
}

// ExistCreate exist create
func (s *adminRepo) ExistCreate(ctx context.Context, dataModel *entities.Admin) (anotherModel *entities.Admin, isNotFound bool, err error) {
	return s.existCreate(ctx, s.dbConn, dataModel)
}

// ExistCreateWithDBConn exist create
func (s *adminRepo) ExistCreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *entities.Admin) (anotherModel *entities.Admin, isNotFound bool, err error) {
	return s.existCreate(ctx, dbConn, dataModel)
}

// createInBatches create many
func (s *adminRepo) createInBatches(ctx context.Context, dbConn *gorm.DB, dataModels []*entities.Admin, batchSize int) (err error) {
	err = dbConn.WithContext(ctx).CreateInBatches(dataModels, batchSize).Error
	if err != nil {
		return err
	}
	return
}

// CreateInBatches create many
func (s *adminRepo) CreateInBatches(ctx context.Context, dataModels []*entities.Admin, batchSize int) error {
	return s.createInBatches(ctx, s.dbConn, dataModels, batchSize)
}

// CreateInBatchesWithDBConn create many
func (s *adminRepo) CreateInBatchesWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*entities.Admin, batchSize int) error {
	return s.createInBatches(ctx, dbConn, dataModels, batchSize)
}

// =============== 更新 ===============

// update update
func (s *adminRepo) update(ctx context.Context, dbConn *gorm.DB, dataModel *entities.Admin) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.AdminSchema.TableName()).
		Where("id = ?", dataModel.Id).
		Save(dataModel).Error
	if err != nil {
		return err
	}
	return
}

// Update update
func (s *adminRepo) Update(ctx context.Context, dataModel *entities.Admin) error {
	return s.update(ctx, s.dbConn, dataModel)
}

// UpdateWithDBConn update
func (s *adminRepo) UpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *entities.Admin) error {
	return s.update(ctx, dbConn, dataModel)
}

// existUpdate exist update
func (s *adminRepo) existUpdate(ctx context.Context, dbConn *gorm.DB, dataModel *entities.Admin) (anotherModel *entities.Admin, isNotFound bool, err error) {
	anotherModel = new(entities.Admin)
	err = dbConn.WithContext(ctx).
		Table(s.AdminSchema.TableName()).
		Where("id = ?", dataModel.Id).
		Where("id != ?", dataModel.Id).
		First(anotherModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			isNotFound = true
			err = nil
		}
		return
	}
	return
}

// ExistUpdate exist update
func (s *adminRepo) ExistUpdate(ctx context.Context, dataModel *entities.Admin) (anotherModel *entities.Admin, isNotFound bool, err error) {
	return s.existUpdate(ctx, s.dbConn, dataModel)
}

// ExistUpdateWithDBConn exist update
func (s *adminRepo) ExistUpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *entities.Admin) (anotherModel *entities.Admin, isNotFound bool, err error) {
	return s.existUpdate(ctx, dbConn, dataModel)
}

// =============== query one : 查一个 ===============

// queryOneById query one by id
func (s *adminRepo) queryOneById(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *entities.Admin, isNotFound bool, err error) {
	dataModel = new(entities.Admin)
	err = dbConn.WithContext(ctx).
		Table(s.AdminSchema.TableName()).
		Where("id = ?", id).
		First(dataModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			isNotFound = true
		}
		return
	}
	return
}

// QueryOneById query one by id
func (s *adminRepo) QueryOneById(ctx context.Context, id interface{}) (dataModel *entities.Admin, isNotFound bool, err error) {
	return s.queryOneById(ctx, s.dbConn, id)
}

// QueryOneByIdWithDBConn query one by id
func (s *adminRepo) QueryOneByIdWithDBConn(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *entities.Admin, isNotFound bool, err error) {
	return s.queryOneById(ctx, dbConn, id)
}

// QueryOneByAdminUuid query one by id
func (s *adminRepo) QueryOneByAdminUuid(ctx context.Context, adminUuid string) (dataModel *entities.Admin, isNotFound bool, err error) {
	dataModel = new(entities.Admin)
	err = s.dbConn.WithContext(ctx).
		Table(s.AdminSchema.TableName()).
		Where("admin_uuid = ?", adminUuid).
		First(dataModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			isNotFound = true
		}
		return
	}
	return
}

// queryOneByConditions query one by conditions
func (s *adminRepo) queryOneByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *entities.Admin, isNotFound bool, err error) {
	dataModel = new(entities.Admin)
	dbConn = dbConn.WithContext(ctx).Table(s.AdminSchema.TableName())
	err = s.WhereConditions(dbConn, conditions).
		First(dataModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			isNotFound = true
		}
		return
	}
	return
}

// QueryOneByConditions query one by conditions
func (s *adminRepo) QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *entities.Admin, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, s.dbConn, conditions)
}

// QueryOneByConditionsWithDBConn query one by conditions
func (s *adminRepo) QueryOneByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *entities.Admin, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, dbConn, conditions)
}

// =============== query all : 查全部 ===============

// queryAllByConditions query all by conditions
func (s *adminRepo) queryAllByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModels []*entities.Admin, err error) {
	dbConn = dbConn.WithContext(ctx).Table(s.AdminSchema.TableName())
	err = s.WhereConditions(dbConn, conditions).
		Find(&dataModels).Error
	if err != nil {
		return dataModels, err
	}
	return
}

// QueryAllByConditions query all by conditions
func (s *adminRepo) QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) ([]*entities.Admin, error) {
	return s.queryAllByConditions(ctx, s.dbConn, conditions)
}

// QueryAllByConditionsWithDBConn query all by conditions
func (s *adminRepo) QueryAllByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) ([]*entities.Admin, error) {
	return s.queryAllByConditions(ctx, dbConn, conditions)
}

// =============== list : 列表 ===============

// list 列表
func (s *adminRepo) list(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormutil.PaginatorArgs) (dataModels []*entities.Admin, recordCount int64, err error) {
	// query where
	dbConn = dbConn.WithContext(ctx).Table(s.AdminSchema.TableName())
	dbConn = s.WhereConditions(dbConn, conditions)
	dbConn = gormutil.AssembleWheres(dbConn, paginatorArgs.PageWheres)

	err = dbConn.Count(&recordCount).Error
	if err != nil {
		return
	} else if recordCount == 0 {
		return // empty
	}

	// pagination
	dbConn = gormutil.AssembleOrders(dbConn, paginatorArgs.PageOrders)
	err = gormutil.Paginator(dbConn, paginatorArgs.PageOption).
		Find(&dataModels).Error
	if err != nil {
		return
	}
	return
}

// List 列表
func (s *adminRepo) List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormutil.PaginatorArgs) ([]*entities.Admin, int64, error) {
	return s.list(ctx, s.dbConn, conditions, paginatorArgs)
}

// ListWithDBConn 列表
func (s *adminRepo) ListWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormutil.PaginatorArgs) ([]*entities.Admin, int64, error) {
	return s.list(ctx, dbConn, conditions, paginatorArgs)
}

// =============== delete : 删除 ===============

// delete delete one
func (s *adminRepo) delete(ctx context.Context, dbConn *gorm.DB, dataModel *entities.Admin) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.AdminSchema.TableName()).
		Where("id = ?", dataModel.Id).
		Delete(dataModel).Error
	if err != nil {
		return err
	}
	return
}

// Delete delete one
func (s *adminRepo) Delete(ctx context.Context, dataModel *entities.Admin) error {
	return s.delete(ctx, s.dbConn, dataModel)
}

// DeleteWithDBConn delete one
func (s *adminRepo) DeleteWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *entities.Admin) error {
	return s.delete(ctx, dbConn, dataModel)
}

// deleteByIds delete by ids
func (s *adminRepo) deleteByIds(ctx context.Context, dbConn *gorm.DB, ids interface{}) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.AdminSchema.TableName()).
		Where("id in (?)", ids).
		Delete(entities.Admin{}).Error
	if err != nil {
		return err
	}
	return
}

// DeleteByIds delete by ids
func (s *adminRepo) DeleteByIds(ctx context.Context, ids interface{}) error {
	return s.deleteByIds(ctx, s.dbConn, ids)
}

// DeleteByIdsWithDBConn delete by ids
func (s *adminRepo) DeleteByIdsWithDBConn(ctx context.Context, dbConn *gorm.DB, ids interface{}) error {
	return s.deleteByIds(ctx, dbConn, ids)
}

// =============== insert : 批量入库 ===============

var _ gormutil.BatchInsertRepo = new(AdminSlice)

// AdminSlice 表切片
type AdminSlice []*entities.Admin

// TableName 表名
func (s *AdminSlice) TableName() string {
	if len(*s) > 0 {
		return (*s)[0].TableName()
	}
	return (&entities.Admin{}).TableName()
}

// Len 长度
func (s *AdminSlice) Len() int {
	return len(*s)
}

// InsertColumns 批量入库的列
func (s *AdminSlice) InsertColumns() (columnList []string, placeholder string) {
	// columns
	columnList = []string{
		"admin_uuid", "created_time", "updated_time", "deleted_time", "is_deleted", "admin_email",
		"admin_nickname", "admin_avatar", "admin_gender", "admin_status", "active_begin_time", "active_end_time",
		"disable_time", "blacklist_time", "password_hash", "register_from",
	}

	// placeholders
	insertPlaceholderBytes := bytes.Repeat([]byte("?, "), len(columnList))
	insertPlaceholderBytes = bytes.TrimSuffix(insertPlaceholderBytes, []byte(", "))

	return columnList, string(insertPlaceholderBytes)
}

// InsertValues 批量入库的值
func (s *AdminSlice) InsertValues(args *gormutil.BatchInsertValueArgs) (prepareData []interface{}, placeholderSlice []string) {
	dataModels := (*s)[args.StepStart:args.StepEnd]
	for index := range dataModels {
		// placeholder
		placeholderSlice = append(placeholderSlice, "("+args.InsertPlaceholder+")")

		// prepare data
		prepareData = append(prepareData, dataModels[index].AdminUuid)
		prepareData = append(prepareData, dataModels[index].CreatedTime)
		prepareData = append(prepareData, dataModels[index].UpdatedTime)
		prepareData = append(prepareData, dataModels[index].DeletedTime)
		prepareData = append(prepareData, dataModels[index].IsDeleted)
		prepareData = append(prepareData, dataModels[index].AdminEmail)
		prepareData = append(prepareData, dataModels[index].AdminNickname)
		prepareData = append(prepareData, dataModels[index].AdminAvatar)
		prepareData = append(prepareData, dataModels[index].AdminGender)
		prepareData = append(prepareData, dataModels[index].AdminStatus)
		prepareData = append(prepareData, dataModels[index].ActiveBeginTime)
		prepareData = append(prepareData, dataModels[index].ActiveEndTime)
		prepareData = append(prepareData, dataModels[index].DisableTime)
		prepareData = append(prepareData, dataModels[index].BlacklistTime)
		prepareData = append(prepareData, dataModels[index].PasswordHash)
		prepareData = append(prepareData, dataModels[index].RegisterFrom)
	}
	return prepareData, placeholderSlice
}

// UpdateColumns 批量入库的列
func (s *AdminSlice) UpdateColumns() (columnList []string) {
	// columns
	columnList = []string{
		"admin_uuid=excluded.admin_uuid", "created_time=excluded.created_time", "updated_time=excluded.updated_time",
		"deleted_time=excluded.deleted_time", "is_deleted=excluded.is_deleted", "admin_email=excluded.admin_email",
		"admin_nickname=excluded.admin_nickname", "admin_avatar=excluded.admin_avatar", "admin_gender=excluded.admin_gender",
		"admin_status=excluded.admin_status", "active_begin_time=excluded.active_begin_time", "active_end_time=excluded.active_end_time",
		"disable_time=excluded.disable_time", "blacklist_time=excluded.blacklist_time",
		"password_hash=excluded.password_hash", "register_from=excluded.register_from",
	}
	return columnList
}

// ConflictActionForMySQL 更新冲突时的操作
func (s *AdminSlice) ConflictActionForMySQL() (req *gormutil.BatchInsertConflictActionReq) {
	req = &gormutil.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "AS excluded",
		OnConflictTarget:      "ON DUPLICATE KEY",
		OnConflictAction:      "UPDATE " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// ConflictActionForPostgres 更新冲突时的操作
func (s *AdminSlice) ConflictActionForPostgres() (req *gormutil.BatchInsertConflictActionReq) {
	req = &gormutil.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "",
		OnConflictTarget:      "ON CONFLICT(id)",
		OnConflictAction:      "DO UPDATE SET " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// insert 批量插入
func (s *adminRepo) insert(ctx context.Context, dbConn *gorm.DB, dataModels AdminSlice) error {
	return gormutil.BatchInsertWithContext(ctx, dbConn, &dataModels)
}

// Insert 批量插入
func (s *adminRepo) Insert(ctx context.Context, dataModels []*entities.Admin) error {
	return s.insert(ctx, s.dbConn, dataModels)
}

// InsertWithDBConn 批量插入
func (s *adminRepo) InsertWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*entities.Admin) error {
	return s.insert(ctx, dbConn, dataModels)
}

// =============== conditions : 条件 ===============

// WhereConditions orm where
func (s *adminRepo) WhereConditions(dbConn *gorm.DB, conditions map[string]interface{}) *gorm.DB {

	// table name
	//tableName := s.AdminSchema.TableName()

	// On-demand loading

	// id
	//if data, ok := conditions["id"]; ok {
	//	dbConn = dbConn.Where(tableName+".id = ?", data)
	//}

	// admin_uuid
	//if data, ok := conditions["admin_uuid"]; ok {
	//	dbConn = dbConn.Where(tableName+".admin_uuid = ?", data)
	//}

	// created_time
	//if data, ok := conditions["created_time"]; ok {
	//	dbConn = dbConn.Where(tableName+".created_time = ?", data)
	//}

	// updated_time
	//if data, ok := conditions["updated_time"]; ok {
	//	dbConn = dbConn.Where(tableName+".updated_time = ?", data)
	//}

	// deleted_time
	//if data, ok := conditions["deleted_time"]; ok {
	//	dbConn = dbConn.Where(tableName+".deleted_time = ?", data)
	//}

	// is_deleted
	//if data, ok := conditions["is_deleted"]; ok {
	//	dbConn = dbConn.Where(tableName+".is_deleted = ?", data)
	//}

	// admin_email
	//if data, ok := conditions["admin_email"]; ok {
	//	dbConn = dbConn.Where(tableName+".admin_email = ?", data)
	//}

	// admin_nickname
	//if data, ok := conditions["admin_nickname"]; ok {
	//	dbConn = dbConn.Where(tableName+".admin_nickname = ?", data)
	//}

	// admin_avatar
	//if data, ok := conditions["admin_avatar"]; ok {
	//	dbConn = dbConn.Where(tableName+".admin_avatar = ?", data)
	//}

	// admin_gender
	//if data, ok := conditions["admin_gender"]; ok {
	//	dbConn = dbConn.Where(tableName+".admin_gender = ?", data)
	//}

	// admin_status
	//if data, ok := conditions["admin_status"]; ok {
	//	dbConn = dbConn.Where(tableName+".admin_status = ?", data)
	//}

	// active_begin_time
	//if data, ok := conditions["active_begin_time"]; ok {
	//	dbConn = dbConn.Where(tableName+".active_begin_time = ?", data)
	//}

	// active_end_time
	//if data, ok := conditions["active_end_time"]; ok {
	//	dbConn = dbConn.Where(tableName+".active_end_time = ?", data)
	//}

	// disable_time
	//if data, ok := conditions["disable_time"]; ok {
	//	dbConn = dbConn.Where(tableName+".disable_time = ?", data)
	//}

	// blacklist_time
	//if data, ok := conditions["blacklist_time"]; ok {
	//	dbConn = dbConn.Where(tableName+".blacklist_time = ?", data)
	//}

	return dbConn
}

// UpdateColumns update columns
func (s *adminRepo) UpdateColumns(conditions map[string]interface{}) map[string]interface{} {

	// update columns
	updateColumns := make(map[string]interface{})

	// On-demand loading

	// id
	//if data, ok := conditions["id"]; ok {
	//	updateColumns["id"] = data
	//}

	// admin_uuid
	//if data, ok := conditions["admin_uuid"]; ok {
	//	updateColumns["admin_uuid"] = data
	//}

	// created_time
	//if data, ok := conditions["created_time"]; ok {
	//	updateColumns["created_time"] = data
	//}

	// updated_time
	//if data, ok := conditions["updated_time"]; ok {
	//	updateColumns["updated_time"] = data
	//}

	// deleted_time
	//if data, ok := conditions["deleted_time"]; ok {
	//	updateColumns["deleted_time"] = data
	//}

	// is_deleted
	//if data, ok := conditions["is_deleted"]; ok {
	//	updateColumns["is_deleted"] = data
	//}

	// admin_email
	//if data, ok := conditions["admin_email"]; ok {
	//	updateColumns["admin_email"] = data
	//}

	// admin_nickname
	//if data, ok := conditions["admin_nickname"]; ok {
	//	updateColumns["admin_nickname"] = data
	//}

	// admin_avatar
	//if data, ok := conditions["admin_avatar"]; ok {
	//	updateColumns["admin_avatar"] = data
	//}

	// admin_gender
	//if data, ok := conditions["admin_gender"]; ok {
	//	updateColumns["admin_gender"] = data
	//}

	// admin_status
	//if data, ok := conditions["admin_status"]; ok {
	//	updateColumns["admin_status"] = data
	//}

	// active_begin_time
	//if data, ok := conditions["active_begin_time"]; ok {
	//	updateColumns["active_begin_time"] = data
	//}

	// active_end_time
	//if data, ok := conditions["active_end_time"]; ok {
	//	updateColumns["active_end_time"] = data
	//}

	// disable_time
	//if data, ok := conditions["disable_time"]; ok {
	//	updateColumns["disable_time"] = data
	//}

	// blacklist_time
	//if data, ok := conditions["blacklist_time"]; ok {
	//	updateColumns["blacklist_time"] = data
	//}

	return updateColumns
}
