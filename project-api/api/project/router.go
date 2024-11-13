package project

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-api/api/midd"
	"test.com/project-api/router"
)

type RouterProject struct {
}

func init() {
	log.Println("init project router")
	router.Register(&RouterProject{})
}

func (*RouterProject) Route(r *gin.Engine) {
	InitRpcProjectClient()
	h := New()
	group := r.Group("/project/index")
	group.Use(midd.TokenVerify())
	group.POST("", h.index)
	group1 := r.Group("/project/project")
	group1.Use(midd.TokenVerify())
	group1.POST("/selfList", h.myProjectList)
}
