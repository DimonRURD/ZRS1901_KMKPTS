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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
