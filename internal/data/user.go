package data

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"master/internal/biz"
	"time"
)

type User struct {
	Uid       int64          `gorm:"column:uid;primary_key"`
	Username  string         `gorm:"column:username"`
	Password  string         `gorm:"column:password"`
	Email     string         `gorm:"column:email"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Active    bool           `gorm:"column:active"`
}

type userRepo struct {
	data  *Data
	idGen UniqueIDGenerator
}

func NewUserRepo(data *Data, id UniqueIDGenerator) biz.UserRepo {
	return &userRepo{
		data:  data,
		idGen: id,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(password)
	err = u.data.db.Create(&User{
		Uid:       u.idGen.GenerateID(),
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Active:    false,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	if user.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(password)
	}
	err := u.data.db.Model(&User{}).Where("uid = ?", user.Id).Updates(User{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		UpdatedAt: time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) GetUserByID(ctx context.Context, id int64) (*biz.User, error) {
	var user = &User{}
	err := u.data.db.Where("uid = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:        user.Uid,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.UnixMilli(),
		UpdatedAt: user.UpdatedAt.UnixMilli(),
		Active:    user.Active,
	}, nil
}

func (u *userRepo) ListUser(ctx context.Context, page int, size int) ([]*biz.User, int64, error) {
	var users []*User
	var count int64
	err := u.data.db.Model(&User{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = u.data.db.Offset((page - 1) * size).Limit(size).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}
	var result []*biz.User
	for _, user := range users {
		result = append(result, &biz.User{
			Id:        user.Uid,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.UnixMilli(),
			UpdatedAt: user.UpdatedAt.UnixMilli(),
			Active:    user.Active,
		})
	}
	return result, count, nil
}

func (u *userRepo) DeleteUser(ctx context.Context, id int64) error {
	err := u.data.db.Where("uid = ?", id).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}
