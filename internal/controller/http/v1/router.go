package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/plusik10/anti-bruteforce/internal/usecase"
	"github.com/plusik10/anti-bruteforce/pkg/logger"
)

func NewRouter(handler *gin.Engine, n usecase.NetManager, l logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	h := handler.Group("/v1")
	{
		// add manager
		NewIPManageRoute(h, n, l)
	}
}
