package persistence

import (
	"ddd-demo/domain/entity"
	"ddd-demo/domain/repository"
	"fmt"
)

type Repositories struct {
	User repository.UserRepository
	Food repository.FoodRepository
	db   *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		User: NewUserRepository(db),
		Food: NewFoodRepository(db),
		db:   db,
	}, nil
}

//closes the database connectio
func (r *Repositories) Close() error {
	return s.db.Close()
}

//This migrate all tables
func (r *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.User{}, &entity.Food{}).Error
}
