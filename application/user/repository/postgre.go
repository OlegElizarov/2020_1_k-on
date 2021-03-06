package repository

import (
	"github.com/go-park-mail-ru/2020_1_k-on/application/models"
	"github.com/go-park-mail-ru/2020_1_k-on/application/user"
	"github.com/go-park-mail-ru/2020_1_k-on/pkg/errors"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type UserDatabase struct {
	conn   *gorm.DB
	logger *zap.Logger
}

func NewUserDatabase(db *gorm.DB, logger *zap.Logger) user.Repository {
	return &UserDatabase{conn: db, logger: logger}
}

func (udb *UserDatabase) Add(user *models.User) (err error) {
	return udb.conn.Table("kinopoisk.users").Create(user).Error
}

func (udb *UserDatabase) Update(id uint, upUser *models.User) error {
	usr := new(models.User)
	err := udb.conn.Table("kinopoisk.users").Where("id = ?", id).First(usr).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.NewNotFoundError(err.Error())
		}
		return err
	}

	upUser.Id = id
	if upUser.Username != "" {
		usr.Username = upUser.Username
	}
	if upUser.Password != "" {
		usr.Password = upUser.Password
	}
	if upUser.Email != "" {
		usr.Email = upUser.Email
	}
	if upUser.Image != "" {
		usr.Image = upUser.Image
	}
	return udb.conn.Table("kinopoisk.users").Save(usr).Error
}

func (udb *UserDatabase) GetById(id uint) (usr *models.User, err error) {
	usr = new(models.User)
	err = udb.conn.Table("kinopoisk.users").Where("id = ?", id).First(usr).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.NewNotFoundError(err.Error())
	}
	return usr, err
}

func (udb *UserDatabase) GetByName(login string) (usr *models.User, err error) {
	usr = new(models.User)
	err = udb.conn.Table("kinopoisk.users").Where("username = ?", login).First(usr).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.NewNotFoundError(err.Error())
	}
	return usr, err
}

func (udb *UserDatabase) Contains(login string) (bool, error) {
	_, err := udb.GetByName(login)

	if err != nil {
		switch err.(type) {
		case *errors.NotFoundError:
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func (udb *UserDatabase) SetImage(id uint, image string) error {
	usr := new(models.User)
	err := udb.conn.Table("kinopoisk.users").Where("id = ?", id).First(usr).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.NewNotFoundError(err.Error())
		}
		return err
	}

	usr.Image = image
	return udb.conn.Table("kinopoisk.users").Save(usr).Error
}

func (udb *UserDatabase) GetOauthConfig() (*models.OauthConfig, error) {
	conf := new(models.OauthConfig)
	err := udb.conn.Table("kinopoisk.oauth").First(conf).Error
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func (udb *UserDatabase) GetUserFromVk(vkUserId int64) (*models.User, error) {
	usr := new(models.User)

	err := udb.conn.Table("kinopoisk.users").
		Select("id, username, password, email, image").
		Joins("inner join kinopoisk.vkusers vk on vk.user_id = id").
		Where("vk.vk_user_id = ?", vkUserId).First(usr).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return usr, nil
}

func (udb *UserDatabase) CreateUserFromVk(vkUserId int64, usr *models.User) error {
	err := udb.Add(usr)
	if err != nil {
		return err
	}

	return udb.conn.Exec("insert into kinopoisk.vkusers(user_id, vk_user_id) values(?, ?)", usr.Id, vkUserId).Error
}
