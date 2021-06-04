package main

import (
	"ddd-demo/app"
	"ddd-demo/infrastructure/persistence"
	"ddd-demo/interfaces"
	"ddd-demo/interfaces/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	app.InitConfig()
	services, err := persistence.NewRepositories("mysql", app.Config.DB.User, app.Config.DB.Password, app.Config.DB.Port, app.Config.DB.Host, app.Config.DB.Name)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.AutoMigrate()

	//redisService, err := auth.NewRedisDB(redis_host, redis_port, redis_password)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//tk := auth.NewToken()

	users := interfaces.NewUsers(services.User)
	//foods := interfaces.NewFoods(services.Food)
	authenticate := interfaces.NewAuthenticate(services.User)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) //For CORS

	//user routes
	r.POST("/users", users.SaveUser)
	r.GET("/users", users.GetUsers)
	r.GET("/users/:user_id", users.GetUser)

	//authentication routes
	r.POST("/login", authenticate.Login)
	r.POST("/logout", authenticate.Logout)
	r.POST("/refresh", authenticate.Refresh)


	//Starting the application
	app_port := os.Getenv("PORT") //using heroku host
	if app_port == "" {
		app_port = "8888" //localhost
	}
	log.Fatal(r.Run(":"+app_port))
}