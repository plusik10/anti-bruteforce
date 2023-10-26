package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/plusik10/anti-bruteforce/pkg/logger"
)

// TODO: add manager usecase.NetManager to param NewRouter and to param NewIPManageRoute.
func NewRouter(handler *gin.Engine, l logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	h := handler.Group("/v1")
	{
		// add manager
		NewIPManageRoute(h, l)
	}
}
