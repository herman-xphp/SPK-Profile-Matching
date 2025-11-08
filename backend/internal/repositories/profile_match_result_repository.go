package repositories

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type ProfileMatchResultRepository struct {
	db *gorm.DB
}

func NewProfileMatchResultRepository(db *gorm.DB) *ProfileMatchResultRepository {
	return &ProfileMatchResultRepository{db: db}
}

func (r *ProfileMatchResultRepository) Create(pmr *models.ProfileMatchResult) error {
	return r.db.Create(pmr).Error
}

func (r *ProfileMatchResultRepository) GetAll() ([]models.ProfileMatchResult, error) {
	var list []models.ProfileMatchResult
	if err := r.db.Preload("TenagaKerja").Preload("Jabatan").Order("total_score DESC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ProfileMatchResultRepository) GetAllWithRelations() ([]models.ProfileMatchResult, error) {
	var list []models.ProfileMatchResult
	if err := r.db.Preload("TenagaKerja").Preload("Jabatan").Order("total_score DESC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ProfileMatchResultRepository) GetByID(id uint) (*models.ProfileMatchResult, error) {
	var pmr models.ProfileMatchResult
	if err := r.db.Preload("TenagaKerja").Preload("Jabatan").First(&pmr, id).Error; err != nil {
		return nil, err
	}
	return &pmr, nil
}

func (r *ProfileMatchResultRepository) GetByJabatanID(jabatanID uint) ([]models.ProfileMatchResult, error) {
	var list []models.ProfileMatchResult
	if err := r.db.Where("jabatan_id = ?", jabatanID).Preload("TenagaKerja").Preload("Jabatan").Order("total_score DESC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ProfileMatchResultRepository) GetByTenagaKerjaID(tenagaKerjaID uint) ([]models.ProfileMatchResult, error) {
	var list []models.ProfileMatchResult
	if err := r.db.Where("tenaga_kerja_id = ?", tenagaKerjaID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ProfileMatchResultRepository) CreateBatch(results []models.ProfileMatchResult) error {
	return r.db.Create(&results).Error
}

func (r *ProfileMatchResultRepository) DeleteByJabatanID(jabatanID uint) error {
	return r.db.Where("jabatan_id = ?", jabatanID).Delete(&models.ProfileMatchResult{}).Error
}
