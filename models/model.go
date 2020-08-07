package models

import (
	"log"
	"time"
)

const (
	IMAGE_CLASSIFICATION = iota + 1
	IMAGE_DETECTION
	IMAGE_SEGMENTATION
)

type Model struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index"`
	Name         string     `json:"name" gorm:"not null"`
	Type         uint       `json:"type" gorm:"type:smallint;not null"`
	URL          string     `json:"url"`
	Desc         string     `json:"desc"`
	PriceMonthly uint       `json:"price_monthly"`
	PriceYearly  uint       `json:"price_yearly"`
	PriceTotal   uint       `json:"price_total"`
	User         *User      `json:"user,omitempty" gorm:"foreignkey:UserID"`
	UserID       uint       `json:"user_id" gorm:"not null"`
	Yn           bool       `json:"yn" gorm:"not null;default:true"`
}

func InitTableModel() {
	if !DB.HasTable(Model{}) {
		log.Println(`models.model: creating table "model"`)
		if err := DB.CreateTable(Model{}).Error; err != nil {
			log.Panic(err)
		}
	}
	if err := DB.Model(&Model{}).AddForeignKey(
		"user_id", `"user"(id)`, "RESTRICT", "RESTRICT",
	).Error; err != nil {
		log.Panic(err)
	}
}

func (model *Model) Insert() error {
	return DB.Create(model).Error
}

func (model *Model) Update() error {
	return DB.Model(&model).Updates(model).Error
}

func GetModelCount() (int, error) {
	var count int
	err := DB.Model(Model{}).Count(&count).Error
	return count, err
}

func GetModelList(query map[string]interface{}, offset, limit int) ([]Model, error) {
	var models []Model
	err := DB.Where(query).Offset(offset).Limit(limit).Find(&models).Error
	return models, err
}
