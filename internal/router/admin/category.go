package admin

import "github.com/gin-gonic/gin"

type CategoryRouter struct {
}

func (cr *CategoryRouter) RouterInit(group *gin.RouterGroup) {
	privateRouter := group.Group("category")

}
