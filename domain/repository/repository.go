package repository

import "ddd-demo/domain/entity"

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUser(int64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}

type FoodRepository interface {
	SaveFood(*entity.Food) (*entity.Food, map[string]string)
	GetFood(int64) (*entity.Food, error)
	GetFoods() ([]entity.Food, error)
}