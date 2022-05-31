package repository

import (
	"go-checkin/models"
	"gorm.io/gorm"
)

type DevisiRepository interface {
	FindAll() ([]models.Devisi, error)
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.Devisi, error)
	FindById(id string) (models.Devisi, error)
	FindWhere(email string) (models.Devisi, error)
	//FindBranch(id string) (models.UserDetail, error)
	Save(role models.Devisi) (models.Devisi, error)
	Update(role models.Devisi) (models.Devisi, error)
	Delete(role models.Devisi) error
	Count() (int64, error)
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
	DbInstance() *gorm.DB
}

type devisiRepository struct {
	*gorm.DB
}

func NewDevisiRepository(db *gorm.DB) DevisiRepository {
	return &devisiRepository{DB: db}
}
func (r devisiRepository) FindAll() ([]models.Devisi, error) {
	var entities []models.Devisi
	err := r.DB.Find(&entities).Error
	return entities, err
}

func (r devisiRepository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.Devisi, error) {
	var entity []models.Devisi
	q := r.DB.Order(orderBy + " " + orderType).Limit(limit).Offset(offset)

	for k, v := range keyVal {
		switch operation {
		case "and":
			q = q.Where(k, v)
		case "or":
			q = q.Or(k, v)
		}
	}

	err := q.Find(&entity).Error
	return entity, err
}

func (r devisiRepository) FindById(id string) (models.Devisi, error) {
	var entity models.Devisi
	err := r.DB.Where("id = ?", id).First(&entity).Error
	return entity, err
}

func (r devisiRepository) FindWhere(name string) (models.Devisi, error) {
	var entity models.Devisi
	err := r.DB.Where("name = ?", name).Find(&entity).Error
	return entity, err
}

//func (r roleRepository) FindBranch(id string) (models.UserDetail, error) {
//	var entity models.UserDetail
//	err := r.DB.Where("user_id = ?", id).Find(&entity).Error
//	return entity, err
//}

func (r devisiRepository) Save(entity models.Devisi) (models.Devisi, error) {
	err := r.DB.Create(&entity).Error
	return entity, err
}

func (r devisiRepository) Update(entity models.Devisi) (models.Devisi, error) {
	err := r.DB.Model(models.Devisi{ID: entity.ID}).UpdateColumns(&entity).Error
	return entity, err
}

func (r devisiRepository) Delete(entity models.Devisi) error {
	return r.DB.Delete(&entity).Error
}

func (r devisiRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("devisi").Count(&count).Error
	return count, err
}

func (r devisiRepository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.Devisi{})
	for k, v := range keyVal {
		switch operation {
		case "and":
			q = q.Where(k, v)
		case "or":
			q = q.Or(k, v)
		}
	}

	err := q.Count(&count).Error
	return count, err
}
func (r devisiRepository) DbInstance() *gorm.DB {
	return r.DB
}
