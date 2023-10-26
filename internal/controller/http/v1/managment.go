package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/plusik10/anti-bruteforce/pkg/logger"
)

type ipManageRoute struct {
	l logger.Interface
}

// TODO: add param usecase.NetManager, implement all api method.
func NewIPManageRoute(handler *gin.RouterGroup, l logger.Interface) {
	route := &ipManageRoute{l}
	_ = route // TODO: REMOVE
	h := handler.Group("/ip")
	{
		h.GET("/test")  // TODO: REMOVE
		h.POST("/tset") // TODO: REMOVE
	}
}
