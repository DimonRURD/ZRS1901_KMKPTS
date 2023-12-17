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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/packet", tc.Fetch)
	group.POST("/packet", tc.Create)
}
