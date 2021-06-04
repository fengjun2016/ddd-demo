package persistence

import (
	"ddd-demo/app"
	"ddd-demo/domain/entity"
	"ddd-demo/domain/repository"
	"github.com/jinzhu/gorm"
)

type Repositories struct {
	User repository.UserRepository
	Food repository.FoodRepository
	db   *gorm.DB
}

func NewRepositories(DbDriver, DbUser, DbPassword, DbPort, DbHost,DbName string) (*Repositories, error) {
	//DbUrl := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	mysqlArgs := app.Config.DB.User + ":" + app.Config.DB.Password + "@tcp(" + app.Config.DB.Host + ":" + app.Config.DB.Port + ")/" + app.Config.DB.Name + "?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"
	db, err := gorm.Open(DbDriver, mysqlArgs)
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

//Close the database connection
func (r *Repositories) Close() error {
	return r.db.Close()
}

//AutoMigrate all tables
func (r *Repositories) AutoMigrate() error {
	return r.db.AutoMigrate(&entity.User{}, &entity.Food{}).Error
}
