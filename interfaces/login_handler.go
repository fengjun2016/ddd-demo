package interfaces

import (
	"ddd-demo/application"
	"ddd-demo/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Authenticate struct {
	us application.UserAppInterface
	//rd auth.AuthInterface    权限校验
	//tk auth.TokenInterface   token校验
}

func NewAuthenticate(uApp application.UserAppInterface) *Authenticate  {
	return &Authenticate{
		us: uApp,
	}
}

func (au *Authenticate) Login(c *gin.Context)  {
	var user *entity.User
	//var tokenErr = map[string]string{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	//validate request
	validateUser := user.Validate("login")
	if len(validateUser) > 0 {
		c.JSON(http.StatusUnprocessableEntity, validateUser)
		return
	}

	u, userErr := au.us.GetUserByEmailAndPassword(user)
	if userErr != nil {
		c.JSON(http.StatusInternalServerError, userErr)
	}

	//todo token 校验
	//todo auth 校验

	userData := make(map[string]interface{})
	userData["access_token"] = "access token"
	userData["refresh_token"] = "refresh token"
	userData["id"] = u.ID
	userData["first_name"] = u.FirstName
	userData["last_name"] = u.LastName

	c.JSON(http.StatusOK, userData)
}

func (au *Authenticate) Logout(c *gin.Context)  {
	//check is the user is authenticated first
	//即检查header里面是否有存在token
	//检查access token 和 refresh token  是否有效 如果此时全部都有效的情况下 则删除相应token
	c.JSON(http.StatusOK, "Successfully logged out")
}

//Refresh Todo 刷新token
func (au *Authenticate) Refresh(c *gin.Context)  {
	//1.先校验旧的原来的refresh token 是否有效 需要在jwt里面解析refresh token
	//2.refresh token 是否过期 以及有效
	c.JSON(http.StatusOK, "Successfully refresh token")
}