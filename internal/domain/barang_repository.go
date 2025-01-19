package domain

import (
	"base-code-api/internal/db"
	"base-code-api/internal/model"
)

type BarangRepository interface {
	Create(barang model.Barang) model.Barang
	Get(id string) *model.Barang
	Update(id string, barang model.Barang) *model.Barang
	Delete(id string) bool
}

type barangRepository struct{}

func NewBarangRepository() BarangRepository {
	return &barangRepository{}
}

// Create a new Barang in the database
func (r *barangRepository) Create(barang model.Barang) model.Barang {
	db.DB.Create(&barang)
	return barang
}

// Get a Barang by its ID
func (r *barangRepository) Get(id string) *model.Barang {
	var barang model.Barang
	db.DB.First(&barang, "id = ?", id)
	return &barang
}

// Update a Barang by its ID
func (r *barangRepository) Update(id string, barang model.Barang) *model.Barang {
	db.DB.Model(&barang).Where("id = ?", id).Updates(barang)
	return &barang
}

// Delete a Barang by its ID
func (r *barangRepository) Delete(id string) bool {
	result := db.DB.Delete(&model.Barang{}, "id = ?", id)
	return result.RowsAffected > 0
}
