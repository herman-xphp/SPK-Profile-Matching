package dto

import "time"

// KriteriaResponse represents kriteria data in API response
type KriteriaResponse struct {
	ID        uint           `json:"id"`
	AspekID   uint           `json:"aspek_id"`
	Kode      string         `json:"kode"`
	Nama      string         `json:"nama"`
	IsCore    bool           `json:"is_core"`
	Bobot     float64        `json:"bobot"`
	Aspek     *AspekResponse `json:"aspek,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// KriteriaCreateRequest represents kriteria creation request
type KriteriaCreateRequest struct {
	AspekID uint    `json:"aspek_id,string" binding:"required"`
	Kode    string  `json:"kode" binding:"required"`
	Nama    string  `json:"nama" binding:"required"`
	IsCore  bool    `json:"is_core"`
	Bobot   float64 `json:"bobot" binding:"required,min=0"`
}

// KriteriaUpdateRequest represents kriteria update request
type KriteriaUpdateRequest struct {
	AspekID uint    `json:"aspek_id,omitempty"`
	Kode    string  `json:"kode,omitempty"`
	Nama    string  `json:"nama,omitempty"`
	IsCore  *bool   `json:"is_core,omitempty"`
	Bobot   float64 `json:"bobot,omitempty" binding:"omitempty,min=0"`
}
