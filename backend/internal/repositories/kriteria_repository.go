package repositories

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type KriteriaRepository struct {
	db *gorm.DB
}

func NewKriteriaRepository(db *gorm.DB) *KriteriaRepository {
	return &KriteriaRepository{db: db}
}

func (r *KriteriaRepository) Create(k *models.Kriteria) error {
	return r.db.Create(k).Error
}

func (r *KriteriaRepository) GetAll() ([]models.Kriteria, error) {
	var list []models.Kriteria
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *KriteriaRepository) GetAllWithAspek() ([]models.Kriteria, error) {
	var list []models.Kriteria
	if err := r.db.Preload("Aspek").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *KriteriaRepository) GetByID(id uint) (*models.Kriteria, error) {
	var k models.Kriteria
	if err := r.db.First(&k, id).Error; err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *KriteriaRepository) GetByAspekID(aspekID uint) ([]models.Kriteria, error) {
	var list []models.Kriteria
	if err := r.db.Where("aspek_id = ?", aspekID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *KriteriaRepository) Update(id uint, k *models.Kriteria) error {
	return r.db.Model(&models.Kriteria{}).Where("id = ?", id).Updates(k).Error
}

func (r *KriteriaRepository) Delete(id uint) error {
	return r.db.Delete(&models.Kriteria{}, id).Error
}

