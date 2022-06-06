package service

import (
	"gorm.io/gorm"
	"klik/dto"
	"klik/models"
	"klik/repository"
	"time"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: repository,
	}
}

func (s *ProductService) QueryDatatable(searchValue string, orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.Product, err error) {
	recordTotal, err = s.ProductRepository.Count()

	if searchValue != "" {
		recordFiltered, err = s.ProductRepository.CountWhere("or", map[string]interface{}{
			"sku LIKE ?": "%" + searchValue + "%",
			"product_name LIKE ?":   "%" + searchValue + "%",
		})

		data, err = s.ProductRepository.FindAllWhere("or", orderType, "created_at", limit, offset, map[string]interface{}{
			"sku LIKE ?": "%" + searchValue + "%",
			"product_name LIKE ?":   "%" + searchValue + "%",
		})
		return recordTotal, recordFiltered, data, err
	}

	recordFiltered, err = s.ProductRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.ProductRepository.FindAllWhere("or", orderType, "created_at", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}
func (s *ProductService) SaveProduct(dto dto.ProductDto) (*models.Product, error) {
	entity := models.Product{
		SKU:         dto.SKU,
		ProductName: dto.ProductName,
		CreatedAt:   time.Now(),
	}
	data, err := s.ProductRepository.Save(entity)
	return &data, err
}
func (s *ProductService) FindProductBySku(id string) (*models.Product, error) {
	data, err := s.ProductRepository.FindBySku(id)

	return &data, err
}
func (s *ProductService) DeleteProduct(sku string) error {
	entity := models.Product{
		SKU: sku,
	}
	err := s.ProductRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (s *ProductService) UpdateProduct(sku string, dto dto.ProductDto) (*models.Product, error) {
	var entity models.Product
	entity.SKU = sku
	entity.ProductName = dto.ProductName

	data, err := s.ProductRepository.Update(entity)

	return &data, err
}
func (s *ProductService) GetDbInstance() *gorm.DB {
	return s.ProductRepository.DbInstance()
}
