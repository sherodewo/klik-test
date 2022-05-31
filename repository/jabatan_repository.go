package repository

import (
	"go-checkin/models"
	"gorm.io/gorm"
)

type JabatanRepository interface {
	FindAll() ([]models.Jabatan, error)
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.Jabatan, error)
	FindById(id string) (models.Jabatan, error)
	FindWhere(email string) (models.Jabatan, error)
	//FindBranch(id string) (models.UserDetail, error)
	Save(role models.Jabatan) (models.Jabatan, error)
	Update(role models.Jabatan) (models.Jabatan, error)
	Delete(role models.Jabatan) error
	Count() (int64, error)
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
	DbInstance() *gorm.DB
}

type jabatanRepository struct {
	*gorm.DB
}

func NewJabatanRepository(db *gorm.DB) JabatanRepository {
	return &jabatanRepository{DB: db}
}
func (r jabatanRepository) FindAll() ([]models.Jabatan, error) {
	var entities []models.Jabatan
	err := r.DB.Find(&entities).Error
	return entities, err
}

func (r jabatanRepository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.Jabatan, error) {
	var entity []models.Jabatan
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

func (r jabatanRepository) FindById(id string) (models.Jabatan, error) {
	var entity models.Jabatan
	err := r.DB.Where("id = ?", id).First(&entity).Error
	return entity, err
}

func (r jabatanRepository) FindWhere(name string) (models.Jabatan, error) {
	var entity models.Jabatan
	err := r.DB.Where("name = ?", name).Find(&entity).Error
	return entity, err
}

//func (r roleRepository) FindBranch(id string) (models.UserDetail, error) {
//	var entity models.UserDetail
//	err := r.DB.Where("user_id = ?", id).Find(&entity).Error
//	return entity, err
//}

func (r jabatanRepository) Save(entity models.Jabatan) (models.Jabatan, error) {
	err := r.DB.Create(&entity).Error
	return entity, err
}

func (r jabatanRepository) Update(entity models.Jabatan) (models.Jabatan, error) {
	err := r.DB.Model(models.Jabatan{ID: entity.ID}).UpdateColumns(&entity).Error
	return entity, err
}

func (r jabatanRepository) Delete(entity models.Jabatan) error {
	return r.DB.Delete(&entity).Error
}

func (r jabatanRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("jabatan").Count(&count).Error
	return count, err
}

func (r jabatanRepository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.Jabatan{})
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
func (r jabatanRepository) DbInstance() *gorm.DB {
	return r.DB
}
