package persistence

import (
	"ddd-demo/domain/entity"
	"ddd-demo/domain/repository"
	"errors"
	"strings"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepo {
	return &UserRepo{db}
}

//UserRepo implemenets the repository.UserRepository interface
var _ repository.UserRepository = &UserRepo{}

func (u *UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "duplicate") {
			dbErr["email_taken"] = "email is already taken"
			return nil, dbErr
		}
		//any other db error
		db["db_error"] = "database error"
		return nil, dbErr
	}

	return user, nil
}

func (u *UserRepo) GetUser(id int64) (*entity.User, error) {
	var users []entity.User
	err := r.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (u *UserRepo) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFound(err) {
		return nil, errors.New("user not fount")
	}

	return users, nil
}

func (u *UserRepo) GetUserByEmailAndPassword(u *entity.User) (*entity.User, map[string]string) {
	var user entity.User
	dbErr := map[string]string{}
	err := r.db.Debug().Where("email = ?", u.Email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		dbErr["no_user"] = "user not found"
		return nil, dbErr
	}
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}

	//verify the password
	err = security.VerifyPassword(user.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		dbErr["incorrect_password"] = "incorrect_password"
		return nil, dbErr
	}
	return &user, nil
}
