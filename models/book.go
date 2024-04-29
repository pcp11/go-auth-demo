package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Title     string         `json:"title"`
	Authors   []Person       `json:"authors" gorm:"foreignKey:ID"`
	Subjects  pq.StringArray `json:"subjects" gorm:"type:text[]"`
	Formats   datatypes.JSON `json:"formats"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Person struct {
	ID        int            `json:"id" gorm:"primary_key"`
	Name      string         `json:"name"`
	BirthYear int            `json:"birth_year"`
	DeathYear int            `json:"death_year"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
