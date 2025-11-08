package dto

import "backend/internal/models"

// MapUserToResponse converts User model to UserResponse DTO
func MapUserToResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Nama:      user.Nama,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// MapUsersToResponse converts User slice to UserResponse slice
func MapUsersToResponse(users []models.User) []UserResponse {
	result := make([]UserResponse, len(users))
	for i, user := range users {
		result[i] = MapUserToResponse(&user)
	}
	return result
}

// MapJabatanToResponse converts Jabatan model to JabatanResponse DTO
func MapJabatanToResponse(jabatan *models.Jabatan) JabatanResponse {
	return JabatanResponse{
		ID:        jabatan.ID,
		Nama:      jabatan.Nama,
		Deskripsi: jabatan.Deskripsi,
		CreatedAt: jabatan.CreatedAt,
		UpdatedAt: jabatan.UpdatedAt,
	}
}

// MapJabatansToResponse converts Jabatan slice to JabatanResponse slice
func MapJabatansToResponse(jabatans []models.Jabatan) []JabatanResponse {
	result := make([]JabatanResponse, len(jabatans))
	for i, jabatan := range jabatans {
		result[i] = MapJabatanToResponse(&jabatan)
	}
	return result
}

// MapAspekToResponse converts Aspek model to AspekResponse DTO
func MapAspekToResponse(aspek *models.Aspek) AspekResponse {
	return AspekResponse{
		ID:         aspek.ID,
		Nama:       aspek.Nama,
		Deskripsi:  aspek.Deskripsi,
		Persentase: aspek.Persentase,
		CreatedAt:  aspek.CreatedAt,
		UpdatedAt:  aspek.UpdatedAt,
	}
}

// MapAspeksToResponse converts Aspek slice to AspekResponse slice
func MapAspeksToResponse(aspeks []models.Aspek) []AspekResponse {
	result := make([]AspekResponse, len(aspeks))
	for i, aspek := range aspeks {
		result[i] = MapAspekToResponse(&aspek)
	}
	return result
}

// MapKriteriaToResponse converts Kriteria model to KriteriaResponse DTO
func MapKriteriaToResponse(kriteria *models.Kriteria) KriteriaResponse {
	response := KriteriaResponse{
		ID:        kriteria.ID,
		AspekID:   kriteria.AspekID,
		Kode:      kriteria.Kode,
		Nama:      kriteria.Nama,
		IsCore:    kriteria.IsCore,
		Bobot:     kriteria.Bobot,
		CreatedAt: kriteria.CreatedAt,
		UpdatedAt: kriteria.UpdatedAt,
	}
	
	if kriteria.Aspek.ID != 0 {
		aspek := MapAspekToResponse(&kriteria.Aspek)
		response.Aspek = &aspek
	}
	
	return response
}

// MapKriteriasToResponse converts Kriteria slice to KriteriaResponse slice
func MapKriteriasToResponse(kriterias []models.Kriteria) []KriteriaResponse {
	result := make([]KriteriaResponse, len(kriterias))
	for i, kriteria := range kriterias {
		result[i] = MapKriteriaToResponse(&kriteria)
	}
	return result
}

// MapTargetProfileToResponse converts TargetProfile model to TargetProfileResponse DTO
func MapTargetProfileToResponse(tp *models.TargetProfile) TargetProfileResponse {
	response := TargetProfileResponse{
		ID:          tp.ID,
		JabatanID:   tp.JabatanID,
		KriteriaID:  tp.KriteriaID,
		TargetNilai: tp.TargetNilai,
		CreatedAt:   tp.CreatedAt,
		UpdatedAt:   tp.UpdatedAt,
	}
	
	if tp.Jabatan.ID != 0 {
		jabatan := MapJabatanToResponse(&tp.Jabatan)
		response.Jabatan = &jabatan
	}
	
	if tp.Kriteria.ID != 0 {
		kriteria := MapKriteriaToResponse(&tp.Kriteria)
		response.Kriteria = &kriteria
	}
	
	return response
}

// MapTargetProfilesToResponse converts TargetProfile slice to TargetProfileResponse slice
func MapTargetProfilesToResponse(tps []models.TargetProfile) []TargetProfileResponse {
	result := make([]TargetProfileResponse, len(tps))
	for i, tp := range tps {
		result[i] = MapTargetProfileToResponse(&tp)
	}
	return result
}

// MapTenagaKerjaToResponse converts TenagaKerja model to TenagaKerjaResponse DTO
func MapTenagaKerjaToResponse(tk *models.TenagaKerja) TenagaKerjaResponse {
	return TenagaKerjaResponse{
		ID:        tk.ID,
		NIK:       tk.NIK,
		Nama:      tk.Nama,
		TglLahir:  tk.TglLahir,
		Alamat:    tk.Alamat,
		Telepon:   tk.Telepon,
		CreatedAt: tk.CreatedAt,
		UpdatedAt: tk.UpdatedAt,
	}
}

// MapTenagaKerjasToResponse converts TenagaKerja slice to TenagaKerjaResponse slice
func MapTenagaKerjasToResponse(tks []models.TenagaKerja) []TenagaKerjaResponse {
	result := make([]TenagaKerjaResponse, len(tks))
	for i, tk := range tks {
		result[i] = MapTenagaKerjaToResponse(&tk)
	}
	return result
}

// MapNilaiTenagaKerjaToResponse converts NilaiTenagaKerja model to NilaiTenagaKerjaResponse DTO
func MapNilaiTenagaKerjaToResponse(ntk *models.NilaiTenagaKerja) NilaiTenagaKerjaResponse {
	response := NilaiTenagaKerjaResponse{
		ID:            ntk.ID,
		TenagaKerjaID: ntk.TenagaKerjaID,
		KriteriaID:    ntk.KriteriaID,
		Nilai:         ntk.Nilai,
		CreatedAt:     ntk.CreatedAt,
		UpdatedAt:     ntk.UpdatedAt,
	}
	
	if ntk.TenagaKerja.ID != 0 {
		tk := MapTenagaKerjaToResponse(&ntk.TenagaKerja)
		response.TenagaKerja = &tk
	}
	
	if ntk.Kriteria.ID != 0 {
		kriteria := MapKriteriaToResponse(&ntk.Kriteria)
		response.Kriteria = &kriteria
	}
	
	return response
}

// MapNilaiTenagaKerjasToResponse converts NilaiTenagaKerja slice to NilaiTenagaKerjaResponse slice
func MapNilaiTenagaKerjasToResponse(ntks []models.NilaiTenagaKerja) []NilaiTenagaKerjaResponse {
	result := make([]NilaiTenagaKerjaResponse, len(ntks))
	for i, ntk := range ntks {
		result[i] = MapNilaiTenagaKerjaToResponse(&ntk)
	}
	return result
}

// MapProfileMatchResultToResponse converts ProfileMatchResult model to ProfileMatchResultResponse DTO
func MapProfileMatchResultToResponse(pmr *models.ProfileMatchResult) ProfileMatchResultResponse {
	response := ProfileMatchResultResponse{
		ID:              pmr.ID,
		TenagaKerjaID:   pmr.TenagaKerjaID,
		JabatanID:       pmr.JabatanID,
		TotalScore:      pmr.TotalScore,
		CoreFactor:      pmr.CoreFactor,
		SecondaryFactor: pmr.SecondaryFactor,
		CreatedAt:       pmr.CreatedAt,
		UpdatedAt:       pmr.UpdatedAt,
	}
	response.ScoreTotal = pmr.TotalScore // Alias for frontend compatibility
	
	if pmr.TenagaKerja.ID != 0 {
		tk := MapTenagaKerjaToResponse(&pmr.TenagaKerja)
		response.TenagaKerja = &tk
	}
	
	if pmr.Jabatan.ID != 0 {
		jabatan := MapJabatanToResponse(&pmr.Jabatan)
		response.Jabatan = &jabatan
	}
	
	return response
}

// MapProfileMatchResultsToResponse converts ProfileMatchResult slice to ProfileMatchResultResponse slice
func MapProfileMatchResultsToResponse(pmrs []models.ProfileMatchResult) []ProfileMatchResultResponse {
	result := make([]ProfileMatchResultResponse, len(pmrs))
	for i, pmr := range pmrs {
		result[i] = MapProfileMatchResultToResponse(&pmr)
	}
	return result
}

// MapProfileMatchResultToRankingResponse converts ProfileMatchResult to RankingResponse with rank
func MapProfileMatchResultToRankingResponse(pmr *models.ProfileMatchResult, rank int) RankingResponse {
	response := RankingResponse{
		ID:              pmr.ID,
		Rank:            rank,
		ScoreTotal:      pmr.TotalScore,
		CoreFactor:      pmr.CoreFactor,
		SecondaryFactor: pmr.SecondaryFactor,
		CreatedAt:       pmr.CreatedAt,
	}
	
	if pmr.TenagaKerja.ID != 0 {
		tk := MapTenagaKerjaToResponse(&pmr.TenagaKerja)
		response.TenagaKerja = &tk
	}
	
	if pmr.Jabatan.ID != 0 {
		jabatan := MapJabatanToResponse(&pmr.Jabatan)
		response.Jabatan = &jabatan
	}
	
	return response
}

// MapProfileMatchResultsToRankingResponse converts ProfileMatchResult slice to RankingResponse slice with ranks
func MapProfileMatchResultsToRankingResponse(pmrs []models.ProfileMatchResult) []RankingResponse {
	result := make([]RankingResponse, len(pmrs))
	for i, pmr := range pmrs {
		result[i] = MapProfileMatchResultToRankingResponse(&pmr, i+1)
	}
	return result
}

// MapProfileMatchResultToDetailResponse converts ProfileMatchResult with details to ProfileMatchResultDetailResponse
func MapProfileMatchResultToDetailResponse(pmr *models.ProfileMatchResult, details map[string]interface{}, rank int) ProfileMatchResultDetailResponse {
	response := ProfileMatchResultDetailResponse{
		ID:              pmr.ID,
		TenagaKerjaID:   pmr.TenagaKerjaID,
		JabatanID:       pmr.JabatanID,
		TotalScore:      pmr.TotalScore,
		CoreFactor:      pmr.CoreFactor,
		SecondaryFactor: pmr.SecondaryFactor,
		Rank:            rank,
		ScoreTotal:      pmr.TotalScore,
		CreatedAt:       pmr.CreatedAt,
		UpdatedAt:       pmr.UpdatedAt,
	}

	if pmr.TenagaKerja.ID != 0 {
		tk := MapTenagaKerjaToResponse(&pmr.TenagaKerja)
		response.TenagaKerja = &tk
	}

	if pmr.Jabatan.ID != 0 {
		jabatan := MapJabatanToResponse(&pmr.Jabatan)
		response.Jabatan = &jabatan
	}

	// Convert details map to DTO structure
	detailPerhitungan := DetailPerhitungan{
		Aspek: make(map[string]AspekDetail),
	}

	if aspekMap, ok := details["aspek"].(map[string]map[string]interface{}); ok {
		for aspekNama, aspekData := range aspekMap {
			aspekDetail := AspekDetail{}

			if persentase, ok := aspekData["persentase"].(float64); ok {
				aspekDetail.Persentase = persentase
			}

			if score, ok := aspekData["score"].(float64); ok {
				aspekDetail.Score = score
			}

			if cf, ok := aspekData["cf"].(float64); ok {
				aspekDetail.CF = cf
			}

			if sf, ok := aspekData["sf"].(float64); ok {
				aspekDetail.SF = sf
			}

			// Convert kriteria list
			if kriteriaList, ok := aspekData["kriteria"].([]map[string]interface{}); ok {
				kriteriaDetails := make([]KriteriaDetail, 0, len(kriteriaList))
				for _, k := range kriteriaList {
					kriteriaDetail := KriteriaDetail{}
					if kode, ok := k["kode"].(string); ok {
						kriteriaDetail.Kode = kode
					}
					if nama, ok := k["nama"].(string); ok {
						kriteriaDetail.Nama = nama
					}
					if target, ok := k["target"].(float64); ok {
						kriteriaDetail.Target = target
					}
					if actual, ok := k["actual"].(float64); ok {
						kriteriaDetail.Actual = actual
					}
					if gap, ok := k["gap"].(float64); ok {
						kriteriaDetail.Gap = gap
					}
					if bobotNilai, ok := k["bobot_nilai"].(float64); ok {
						kriteriaDetail.BobotNilai = bobotNilai
					}
					if isCore, ok := k["is_core"].(bool); ok {
						kriteriaDetail.IsCore = isCore
					}
					kriteriaDetails = append(kriteriaDetails, kriteriaDetail)
				}
				aspekDetail.Kriteria = kriteriaDetails
			}

			detailPerhitungan.Aspek[aspekNama] = aspekDetail
		}
	}

	response.Details = detailPerhitungan

	return response
}

