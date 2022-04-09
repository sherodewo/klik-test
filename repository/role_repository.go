package repository

import (
	"go-checkin/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAll() ([]models.UserRole, error)
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.UserRole, error)
	FindById(id string) (models.UserRole, error)
	FindWhere(email string) (models.UserRole, error)
	//FindBranch(id string) (models.UserDetail, error)
	Save(role models.UserRole) (models.UserRole, error)
	Update(role models.UserRole) (models.UserRole, error)
	Delete(role models.UserRole) error
	Count() (int64, error)
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
	DbInstance() *gorm.DB
}

type roleRepository struct {
	*gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{DB: db}
}
func (r roleRepository) FindAll() ([]models.UserRole, error) {
	var entities []models.UserRole
	err := r.DB.Find(&entities).Error
	return entities, err
}

func (r roleRepository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.UserRole, error) {
	var entity []models.UserRole
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

func (r roleRepository) FindById(id string) (models.UserRole, error) {
	var entity models.UserRole
	err := r.DB.Where("id = ?", id).First(&entity).Error
	return entity, err
}

func (r roleRepository) FindWhere(name string) (models.UserRole, error) {
	var entity models.UserRole
	err := r.DB.Where("name = ?", name).Find(&entity).Error
	return entity, err
}

//func (r roleRepository) FindBranch(id string) (models.UserDetail, error) {
//	var entity models.UserDetail
//	err := r.DB.Where("user_id = ?", id).Find(&entity).Error
//	return entity, err
//}

func (r roleRepository) Save(entity models.UserRole) (models.UserRole, error) {
	err := r.DB.Create(&entity).Error
	return entity, err
}

func (r roleRepository) Update(entity models.UserRole) (models.UserRole, error) {
	err := r.DB.Model(models.UserRole{ID: entity.ID}).UpdateColumns(&entity).Error
	return entity, err
}

func (r roleRepository) Delete(entity models.UserRole) error {
	return r.DB.Delete(&entity).Error
}

func (r roleRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("user_role").Count(&count).Error
	return count, err
}

func (r roleRepository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.UserRole{})
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
func (r roleRepository) DbInstance() *gorm.DB {
	return r.DB
}
