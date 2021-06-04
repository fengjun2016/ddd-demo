package persistence

import (
	"ddd-demo/domain/entity"
	"ddd-demo/domain/repository"
	"ddd-demo/infrastructure/auth"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UserRepo struct {
	db *gorm.DB
}

type FoodRepo struct {
	db * gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func NewFoodRepository(db *gorm.DB) *FoodRepo {
	return &FoodRepo{db}
}

//UserRepo implements the repository.UserRepository interface
var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "duplicate") {
			dbErr["email_taken"] = "email is already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}

	return user, nil
}

func (r *UserRepo) GetUser(id int64) (*entity.User, error) {
	var user entity.User
	user.ID = id
	err := r.db.Debug().First(&user).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *UserRepo) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not fount")
	}

	return users, nil
}

func (r *UserRepo) GetUserByEmailAndPassword(u *entity.User) (*entity.User, map[string]string) {
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
	err = auth.Compare(user.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		dbErr["incorrect_password"] = "incorrect_password"
		return nil, dbErr
	}
	return &user, nil
}


func (fr *FoodRepo) SaveFood(food *entity.Food) (*entity.Food, map[string]string) {
	dbErr := map[string]string{}
	err := fr.db.Debug().Create(&food).Error
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "duplicate") {
			dbErr["email_taken"] = "email is already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}

	return food, nil
}

func (fr *FoodRepo) GetFood(id int64) (*entity.Food, error) {
	var food entity.Food
	food.ID = id
	err := fr.db.Debug().First(&food).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}

	return &food, nil
}

func (fr *FoodRepo) GetFoods() ([]entity.Food, error) {
	var foods []entity.Food
	err := fr.db.Debug().Find(&foods).Error
	if err != nil {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not fount")
	}

	return foods, nil
}
