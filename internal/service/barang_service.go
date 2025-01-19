package service

import (
	"base-code-api/internal/domain"
	"base-code-api/internal/model"
)

type BarangService struct {
	repo domain.BarangRepository
}

func NewBarangService(repo domain.BarangRepository) *BarangService {
	return &BarangService{repo: repo}
}

func (s *BarangService) CreateBarang(barang model.Barang) model.Barang {
	return s.repo.Create(barang)
}

func (s *BarangService) GetBarang(id string) *model.Barang {
	return s.repo.Get(id)
}

func (s *BarangService) UpdateBarang(id string, barang model.Barang) *model.Barang {
	return s.repo.Update(id, barang)
}

func (s *BarangService) DeleteBarang(id string) bool {
	return s.repo.Delete(id)
}
