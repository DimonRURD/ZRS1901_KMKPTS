package route

import (
	"time"

	"github.com/
	"github.com/
	"github.com/
	"github.com/
	"github.com/
	"github.com/
	"github.com/
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/packet", tc.Fetch)
	group.POST("/packet", tc.Create)
}
