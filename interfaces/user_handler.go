package interfaces

import (
	"ddd-demo/application"
	"ddd-demo/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Users struct {
	us application.UserAppInterface
	//rd auth.AuthInterface
	//tk auth.TokenInterface
}

//Users constructor
func NewUsers(us application.UserAppInterface) *Users {
	return &Users{
		us: us,
	}
}

func (u *Users) SaveUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})

		return
	}

	newUser, err := u.us.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newUser.PublicUser())
}

func (u *Users) GetUsers(c *gin.Context) {
	users := entity.Users{} //customize user
	var err error
	//us, err = application.UserApp.GetUsers()
	users, err = u.us.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(
		http.StatusOK,
		users.PublicUsers(),
	)
}

func (u *Users) GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := u.us.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user.PublicUser())
}
