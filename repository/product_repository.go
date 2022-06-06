package repository

import (
	"gorm.io/gorm"
	"klik/models"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.Product, error)
	FindBySku(sku string) (models.Product, error)
	FindWhere(email string) (models.Product, error)
	//FindBranch(id string) (models.UserDetail, error)
	Save(role models.Product) (models.Product, error)
	Update(role models.Product) (models.Product, error)
	Delete(role models.Product) error
	Count() (int64, error)
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
	DbInstance() *gorm.DB
}

type productRepository struct {
	*gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{DB: db}
}
func (r productRepository) FindAll() ([]models.Product, error) {
	var entities []models.Product
	err := r.DB.Find(&entities).Error
	return entities, err
}

func (r productRepository) FindBySku(id string) (models.Product, error) {
	var entity models.Product
	err := r.DB.Where("product.sku = ?", id).First(&entity).Error
	return entity, err
}

func (r productRepository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.Product, error) {
	var entity []models.Product
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

func (r productRepository) FindById(id string) (models.Product, error) {
	var entity models.Product
	err := r.DB.Where("id = ?", id).First(&entity).Error
	return entity, err
}

func (r productRepository) FindWhere(name string) (models.Product, error) {
	var entity models.Product
	err := r.DB.Where("name = ?", name).Find(&entity).Error
	return entity, err
}

//func (r roleRepository) FindBranch(id string) (models.UserDetail, error) {
//	var entity models.UserDetail
//	err := r.DB.Where("user_id = ?", id).Find(&entity).Error
//	return entity, err
//}

func (r productRepository) Save(entity models.Product) (models.Product, error) {
	err := r.DB.Create(&entity).Error
	return entity, err
}

func (r productRepository) Update(entity models.Product) (models.Product, error) {
	err := r.DB.Where("sku = ? ", entity.SKU).UpdateColumns(&entity).Error
	return entity, err
}

func (r productRepository) Delete(entity models.Product) error {
	return r.DB.Delete(&entity).Error
}

func (r productRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("product").Count(&count).Error
	return count, err
}

func (r productRepository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.Product{})
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
func (r productRepository) DbInstance() *gorm.DB {
	return r.DB
}
