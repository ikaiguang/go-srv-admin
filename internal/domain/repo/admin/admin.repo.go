// Package repos
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package repos

import (
	context "context"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"

	entities "github.com/ikaiguang/go-srv-admin/internal/domain/entity/admin"
)

// AdminRepo repo
type AdminRepo interface {
	Create(ctx context.Context, dataModel *entities.Admin) error
	ExistCreate(ctx context.Context, dataModel *entities.Admin) (anotherModel *entities.Admin, isNotFound bool, err error)
	CreateInBatches(ctx context.Context, dataModels []*entities.Admin, batchSize int) error
	Insert(ctx context.Context, dataModels []*entities.Admin) error
	Update(ctx context.Context, dataModel *entities.Admin) error
	ExistUpdate(ctx context.Context, dataModel *entities.Admin) (anotherModel *entities.Admin, isNotFound bool, err error)
	QueryOneById(ctx context.Context, id interface{}) (dataModel *entities.Admin, isNotFound bool, err error)
	QueryOneByAdminUuid(ctx context.Context, adminUuid string) (dataModel *entities.Admin, isNotFound bool, err error)
	QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *entities.Admin, isNotFound bool, err error)
	QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) (dataModels []*entities.Admin, err error)
	List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormutil.PaginatorArgs) (dataModels []*entities.Admin, totalNumber int64, err error)
	Delete(ctx context.Context, dataModel *entities.Admin) error
	DeleteByIds(ctx context.Context, ids interface{}) error
}
