package service

import (
	"go-checkin/dto"
	"go-checkin/models"
	"go-checkin/repository"
	"gorm.io/gorm"
	"time"
)

type DevisiService struct {
	DevisiRepository repository.DevisiRepository
}

func NewDevisiService(repository repository.DevisiRepository) *DevisiService {
	return &DevisiService{
		DevisiRepository: repository,
	}
}

func (s *DevisiService) QueryDatatable(searchValue string, orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.Devisi, err error) {
	recordTotal, err = s.DevisiRepository.Count()

	if searchValue != "" {
		recordFiltered, err = s.DevisiRepository.CountWhere("or", map[string]interface{}{
			"name LIKE ?": "%" + searchValue + "%",
			"id LIKE ?":   "%" + searchValue + "%",
		})

		data, err = s.DevisiRepository.FindAllWhere("or", orderType, "created_at", limit, offset, map[string]interface{}{
			"name LIKE ?": "%" + searchValue + "%",
			"id LIKE ?":   "%" + searchValue + "%",
		})
		return recordTotal, recordFiltered, data, err
	}

	recordFiltered, err = s.DevisiRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.DevisiRepository.FindAllWhere("or", orderType, "created_at", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}
func (s *DevisiService) SaveDevisi(dto dto.DevisiDto) (*models.Devisi, error) {
	entity := models.Devisi{
		Name:        dto.Name,
		Description: dto.Description,
		CreatedAt:   time.Now(),
	}
	data, err := s.DevisiRepository.Save(entity)
	return &data, err
}
func (s *DevisiService) FindUserById(id string) (*models.Devisi, error) {
	data, err := s.DevisiRepository.FindById(id)

	return &data, err
}
func (s *DevisiService) DeleteDevisi(id string) error {
	entity := models.Devisi{
		ID: id,
	}
	err := s.DevisiRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (s *DevisiService) UpdateDevisi(id string, dto dto.DevisiDto) (*models.Devisi, error) {
	var entity models.Devisi
	entity.ID = id
	entity.Name = dto.Name
	entity.Description = dto.Description

	data, err := s.DevisiRepository.Update(entity)

	return &data, err
}
func (s *DevisiService) GetDbInstance() *gorm.DB {
	return s.DevisiRepository.DbInstance()
}
