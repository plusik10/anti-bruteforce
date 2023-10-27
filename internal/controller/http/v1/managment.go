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
	h := handler.Group("/ip")
	{
		h.GET("/help-method")
		h.GET("/auth-attempt")  // TODO: implement
		h.POST("/bucket-clean") //TODO: implement

		h.POST("/add-to-blacklist", route.addToBlackList)
		h.DELETE("/delete-from-blacklist", route.deleteIPFromList)
		h.POST("/add-to-whitelist", route.addToWhiteList)
		h.DELETE("/delete-from-whitelist", route.deleteIPFromList)
	}
}

func (i *ipManageRoute) help(c *gin.Context) {
	c.JSON(http.StatusOK, "/add-to-blacklist, /delete-from-blacklist, /add-to-whitelist, /delete-from-whitelist")
}

func (i *ipManageRoute) deleteIPFromList(c *gin.Context) {
	var request doRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		i.l.Error(err, "http-v1-deleteIPFromList - doRequest")
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	err := i.netManger.DeleteIpFromStorage(c.Request.Context(), request.IP)
	if err != nil {
		i.l.Error(err, "http - v1 -deleteIPFromList - netManger.DeleteFromBlackList")
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (i *ipManageRoute) addToWhiteList(c *gin.Context) {
	var request doRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		i.l.Error(err, "http-v1-addToWhiteList - doRequest")
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	err := i.netManger.AddIPToWhiteList(c.Request.Context(), request.IP)
	if err != nil {
		i.l.Error(err, "http - v1 -addToWhiteList - netManger.AddIPToWhiteList")
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (i *ipManageRoute) addToBlackList(c *gin.Context) {
	var request doRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		i.l.Error(err, "http-v1-addToBlackList - doRequest")
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	err := i.netManger.AddIPToBlackList(c.Request.Context(), request.IP)
	if err != nil {
		i.l.Error(err, "http - v1 -addToBlackList - netManger.AddIPToBlacklist")
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}
