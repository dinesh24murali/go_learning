package common

import "gorm.io/gorm"

type Repository[T any] interface {
	Create(entity *T) error
	FindAll() ([]T, error)
	FindByID(id string) (*T, error)
	Update(entity *T) error
	Delete(id string) error
}

type GormRepository[T any] struct {
	Db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{Db: db}
}

func (r *GormRepository[T]) Create(entity *T) error {
	return r.Db.Create(entity).Error
}

func (r *GormRepository[T]) FindAll() ([]T, error) {
	var entities []T
	err := r.Db.Find(&entities).Error
	return entities, err
}

func (r *GormRepository[T]) FindByID(id string) (*T, error) {
	var entity T
	err := r.Db.First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GormRepository[T]) Update(entity *T) error {
	return r.Db.Save(entity).Error
}

func (r *GormRepository[T]) Delete(id string) error {
	return r.Db.Delete(new(T), "id = ?", id).Error
}
