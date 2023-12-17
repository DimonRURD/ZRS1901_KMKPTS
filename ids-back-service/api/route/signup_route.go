package route

import (
	"time"

	"github.com/DimonRURD/ZRS1901_KMKPTS/api/controller"
	"github.com/DimonRURD/ZRS1901_KMKPTS/bootstrap"
	"github.com/DimonRURD/ZRS1901_KMKPTS/domain"
	"github.com/DimonRURD/ZRS1901_KMKPTS/mongo"
	"github.com/DimonRURD/ZRS1901_KMKPTS/repository"
	"github.com/DimonRURD/ZRS1901_KMKPTS/usecase"
	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
