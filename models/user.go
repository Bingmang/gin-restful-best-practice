package models

import (
	"log"
	"time"
)

type User struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index"`
	LastLoginAt  time.Time  `json:"last_login_at"`
	Username     string     `json:"username" gorm:"unique;not null"`
	Password     string     `json:"password,omitempty" gorm:"not null"`
	Phone        string     `json:"phone" gorm:"unique;not null"`
	Avatar       string     `json:"avatar"`
	Mail         string     `json:"mail"`
	Organization string     `json:"organization"`
	Role         *Role      `json:"role,omitempty" gorm:"foreignkey:RoleID"`
	RoleID       uint       `json:"role_id" gorm:"default:3"`
	Yn           bool       `json:"yn" gorm:"not null;default:true"`
}

func InitTableUser() {
	if !DB.HasTable(User{}) {
		log.Println(`models.user: creating table "user"`)
		if err := DB.CreateTable(User{}).Error; err != nil {
			log.Panic(err)
		}
	}
	if err := DB.Model(&User{}).AddForeignKey(
		"role_id", "role(id)", "RESTRICT", "RESTRICT",
	).Error; err != nil {
		log.Panic(err)
	}
}

func (user *User) Insert() error {
	return DB.Create(user).Error
}

func (user *User) Update() error {
	return DB.Model(&user).Updates(user).Error
}

func GetUserCount() (int, error) {
	var count int
	err := DB.Model(User{}).Count(&count).Error
	return count, err
}

func GetUserList(offset, limit int) ([]User, error) {
	var users []User
	err := DB.Where("yn = ?", true).Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func GetUserByID(id int) (User, error) {
	var user User
	DB.Where("yn = ?", true).First(&user, id)
	err := DB.Model(&user).Related(&user.Role).Error
	return user, err
}

func GetUserByUsername(username string) (User, error) {
	var user User
	DB.Where("username = ?", username).First(&user)
	err := DB.Model(&user).Related(&user.Role).Error
	return user, err
}
