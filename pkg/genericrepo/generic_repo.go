package genericrepo

import (
	"context"
	"github.com/sinameshkini/microkit/models"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func New[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{
		db: db,
	}
}

func (r *Repository[T]) Add(entity *T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *Repository[T]) AddAll(entity *[]T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *Repository[T]) GetById(id int, ctx context.Context) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).Model(&entity).Where("id = ? AND is_active = ?", id, true).FirstOrInit(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *Repository[T]) Get(params *T, ctx context.Context) *T {
	var entity T
	r.db.WithContext(ctx).Where(&params).FirstOrInit(&entity)
	return &entity
}

func (r *Repository[T]) GetAll(ctx context.Context, req *models.Request) (*[]T, *models.PaginationResponse, error) {
	var (
		err      error
		entities []T
		entity   T
		total    int64
		meta     *models.PaginationResponse
	)

	query := r.db.WithContext(ctx).Model(entity)

	query = req.AddToQuery(query)

	if req.GetPagination {
		total, err = models.GetCount(query)
		if err != nil {
			return nil, nil, err
		}

		if total == 0 {
			return nil, nil, models.ErrNotfound
		}
	}

	query = req.PaginationRequest.ToQuery(query)

	if err = query.Find(&entities).Error; err != nil {
		return nil, nil, err
	}

	if req.GetPagination {
		meta = models.MakePaginationResponse(total, req.Page, req.PerPage)
	}

	return &entities, meta, nil
}

func (r *Repository[T]) Where(params *T, ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Where(&params).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *Repository[T]) Update(entity *T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *Repository[T]) UpdateAll(entities *[]T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entities).Error
}

func (r *Repository[T]) Delete(id int, ctx context.Context) error {
	var entity T
	return r.db.WithContext(ctx).Delete(&entity, id).Error
}

func (r *Repository[T]) SkipTake(skip int, take int, ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Offset(skip).Limit(take).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *Repository[T]) Count(ctx context.Context) int64 {
	var entity T
	var count int64
	r.db.WithContext(ctx).Model(&entity).Count(&count)
	return count
}

func (r *Repository[T]) CountWhere(params *T, ctx context.Context) int64 {
	var entity T
	var count int64
	r.db.WithContext(ctx).Model(&entity).Where(&params).Count(&count)
	return count
}