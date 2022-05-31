package service

import (
	"go-checkin/dto"
	"go-checkin/models"
	"go-checkin/repository"
	"gorm.io/gorm"
	"time"
)

type JabatanService struct {
	JabatanRepository repository.JabatanRepository
}

func NewJabatanService(repository repository.JabatanRepository) *JabatanService {
	return &JabatanService{
		JabatanRepository: repository,
	}
}

func (s *JabatanService) QueryDatatable(searchValue string, orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.Jabatan, err error) {
	recordTotal, err = s.JabatanRepository.Count()

	if searchValue != "" {
		recordFiltered, err = s.JabatanRepository.CountWhere("or", map[string]interface{}{
			"name LIKE ?": "%" + searchValue + "%",
			"id LIKE ?":   "%" + searchValue + "%",
		})

		data, err = s.JabatanRepository.FindAllWhere("or", orderType, "created_at", limit, offset, map[string]interface{}{
			"name LIKE ?": "%" + searchValue + "%",
			"id LIKE ?":   "%" + searchValue + "%",
		})
		return recordTotal, recordFiltered, data, err
	}

	recordFiltered, err = s.JabatanRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.JabatanRepository.FindAllWhere("or", orderType, "created_at", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}
func (s *JabatanService) SaveJabatan(dto dto.JabatanDto) (*models.Jabatan, error) {
	entity := models.Jabatan{
		Name:        dto.Name,
		Description: dto.Description,
		CreatedAt:   time.Now(),
	}
	data, err := s.JabatanRepository.Save(entity)
	return &data, err
}
func (s *JabatanService) FindUserById(id string) (*models.Jabatan, error) {
	data, err := s.JabatanRepository.FindById(id)

	return &data, err
}
func (s *JabatanService) DeleteJabatan(id string) error {
	entity := models.Jabatan{
		ID: id,
	}
	err := s.JabatanRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (s *JabatanService) UpdateJabatan(id string, dto dto.JabatanDto) (*models.Jabatan, error) {
	var entity models.Jabatan
	entity.ID = id
	entity.Name = dto.Name
	entity.Description = dto.Description

	data, err := s.JabatanRepository.Update(entity)

	return &data, err
}
func (s *JabatanService) GetDbInstance() *gorm.DB {
	return s.JabatanRepository.DbInstance()
}
