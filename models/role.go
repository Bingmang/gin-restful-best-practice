package models

import "log"

type Role struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Right string `json:"right"`
}

/*
前4位用于管理，后4位用户用户
目前都是预留功能
*/
func InitTableRole() {
	if DB.HasTable(Role{}) {
		return
	}
	log.Println(`models.role: creating table "role"`)
	if err := DB.CreateTable(Role{}).Error; err != nil {
		log.Panic(err)
	}
	DB.Create(&Role{
		ID:    1,
		Name:  "admin",
		Right: "11111111",
	})
	DB.Create(&Role{
		ID:    2,
		Name:  "manager",
		Right: "01111111",
	})
	DB.Create(&Role{
		ID:    3,
		Name:  "user",
		Right: "00001111",
	})
}
