package service

import (
	"klik/dto"
	"klik/models"
	"klik/repository"
	"gorm.io/gorm"
	"time"
)

type RoleService struct {
	RoleRepository repository.RoleRepository
}

func NewRoleService(repository repository.RoleRepository) *RoleService {
	return &RoleService{
		RoleRepository: repository,
	}
}

func (s *RoleService) QueryDatatable(searchValue string, orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.UserRole, err error) {
	recordTotal, err = s.RoleRepository.Count()

	if searchValue != "" {
		recordFiltered, err = s.RoleRepository.CountWhere("or", map[string]interface{}{
			"name LIKE ?": "%" + searchValue + "%",
			"id LIKE ?":   "%" + searchValue + "%",
		})

		data, err = s.RoleRepository.FindAllWhere("or", orderType, "created_at", limit, offset, map[string]interface{}{
			"name LIKE ?": "%" + searchValue + "%",
			"id LIKE ?":   "%" + searchValue + "%",
		})
		return recordTotal, recordFiltered, data, err
	}

	recordFiltered, err = s.RoleRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.RoleRepository.FindAllWhere("or", orderType, "created_at", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}
func (s *RoleService) SaveRole(dto dto.RoleDto) (*models.UserRole, error) {
	entity := models.UserRole{
		Name:        dto.Name,
		Description: dto.Description,
		CreatedAt:   time.Now(),
	}
	data, err := s.RoleRepository.Save(entity)
	return &data, err
}
func (s *RoleService) FindUserById(id string) (*models.UserRole, error) {
	data, err := s.RoleRepository.FindById(id)

	return &data, err
}
func (s *RoleService) DeleteRole(id string) error {
	entity := models.UserRole{
		ID: id,
	}
	err := s.RoleRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (s *RoleService) UpdateRole(id string, dto dto.RoleDto) (*models.UserRole, error) {
	var entity models.UserRole
	entity.ID = id
	entity.Name = dto.Name
	entity.Description = dto.Description

	data, err := s.RoleRepository.Update(entity)

	return &data, err
}
func (s *RoleService) GetDbInstance() *gorm.DB {
	return s.RoleRepository.DbInstance()
}
