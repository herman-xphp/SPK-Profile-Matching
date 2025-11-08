package dto

import "time"

// CalculationRequest represents profile matching calculation request
type CalculationRequest struct {
	JabatanID      uint   `json:"jabatan_id" binding:"required"`
	TenagaKerjaIDs []uint `json:"tenaga_kerja_ids,omitempty"` // Optional: if empty, calculate for all
}

// ProfileMatchResultResponse represents profile matching result in API response
type ProfileMatchResultResponse struct {
	ID              uint                  `json:"id"`
	TenagaKerjaID   uint                  `json:"tenaga_kerja_id"`
	JabatanID       uint                  `json:"jabatan_id"`
	TotalScore      float64               `json:"total_score"`
	CoreFactor      float64               `json:"core_factor"`
	SecondaryFactor float64               `json:"secondary_factor"`
	TenagaKerja     *TenagaKerjaResponse  `json:"tenaga_kerja,omitempty"`
	Jabatan         *JabatanResponse      `json:"jabatan,omitempty"`
	Rank            int                   `json:"rank,omitempty"`
	ScoreTotal      float64               `json:"score_total,omitempty"` // Alias for TotalScore for frontend compatibility
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
}

// RankingResponse represents ranking response with additional fields for frontend
type RankingResponse struct {
	ID              uint                  `json:"id"`
	Rank            int                   `json:"rank"`
	ScoreTotal      float64               `json:"score_total"`
	CoreFactor      float64               `json:"core_factor"`
	SecondaryFactor float64               `json:"secondary_factor"`
	TenagaKerja     *TenagaKerjaResponse  `json:"tenaga_kerja,omitempty"`
	Jabatan         *JabatanResponse      `json:"jabatan,omitempty"`
	CreatedAt       time.Time             `json:"created_at"`
}

// KriteriaDetail represents detail kriteria in calculation
type KriteriaDetail struct {
	Kode       string  `json:"kode"`
	Nama       string  `json:"nama"`
	Target     float64 `json:"target"`
	Actual     float64 `json:"actual"`
	Gap        float64 `json:"gap"`
	BobotNilai float64 `json:"bobot_nilai"`
	IsCore     bool    `json:"is_core"`
}

// AspekDetail represents detail perhitungan per aspek
type AspekDetail struct {
	Persentase float64         `json:"persentase"`
	Score      float64         `json:"score"`
	CF         float64         `json:"cf"`
	SF         float64         `json:"sf"`
	Kriteria   []KriteriaDetail `json:"kriteria"`
}

// DetailPerhitungan represents detail perhitungan structure
type DetailPerhitungan struct {
	Aspek map[string]AspekDetail `json:"aspek"`
}

// ProfileMatchResultDetailResponse represents detailed profile matching result with calculation details
type ProfileMatchResultDetailResponse struct {
	ID              uint                  `json:"id"`
	TenagaKerjaID   uint                  `json:"tenaga_kerja_id"`
	JabatanID       uint                  `json:"jabatan_id"`
	TotalScore      float64               `json:"total_score"`
	CoreFactor      float64               `json:"core_factor"`
	SecondaryFactor float64               `json:"secondary_factor"`
	TenagaKerja     *TenagaKerjaResponse  `json:"tenaga_kerja,omitempty"`
	Jabatan         *JabatanResponse      `json:"jabatan,omitempty"`
	Rank            int                   `json:"rank,omitempty"`
	ScoreTotal      float64               `json:"score_total,omitempty"`
	Details         DetailPerhitungan     `json:"details"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
}

