package http

import (
	"gorm.io/gorm"
	"wilder.cn/gogo/sys/db"
)

type Repository interface {
	Save(user db.SUser) (db.SUser, error)
	FindByEmail(email string) (db.SUser, error)
	FindByID(ID string) (db.SUser, error)
	Update(user db.SUser) (db.SUser, error)
	UpdatePwd(email string, pwd string) bool
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository) Save(user db.SUser) (db.SUser, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// FindByEmail 只用于管理员
func (r *repository) FindByEmail(email string) (db.SUser, error) {
	var user db.SUser
	err := r.db.Where("id = ? and state = ?", email, StateAdmin).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID string) (db.SUser, error) {
	var user db.SUser
	err := r.db.Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user db.SUser) (db.SUser, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) UpdatePwd(email string, pwd string) bool {
	tb := r.db.Model(&db.SUser{}).Where("email=?", email).Update("password", pwd)
	return tb.RowsAffected == 1
}
