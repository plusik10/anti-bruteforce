package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/plusik10/anti-bruteforce/internal/usecase"
	"github.com/plusik10/anti-bruteforce/pkg/logger"
	"net/http"
)

type ipManageRoute struct {
	netManger usecase.NetManager
	l         logger.Interface
}

type doRequest struct {
	IP string `json:"IP" example:"192.0.0.24"`
}

func NewIPManageRoute(handler *gin.RouterGroup, n usecase.NetManager, l logger.Interface) {
	route := &ipManageRoute{n, l}
	_ = route // TODO: REMOVE
	h := handler.Group("/ip")
	{
		h.POST("/add-to-blacklist", route.addToBlackList) // TODO: REMOVE
		h.POST("/tset")                                   // TODO: REMOVE
	}
}

func (i *ipManageRoute) addToBlackList(c *gin.Context) {
	var request doRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		i.l.Error(err, "http-v1-doRequest")
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	err := i.netManger.AddIPToBlackList(c.Request.Context(), request.IP)
	if err != nil {
		i.l.Error(err, "http - v1 - doTranslate")
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}
