package repository

import "ddd-demo/domain/entity"

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUser(int64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}

type FoodRepository interface {

}