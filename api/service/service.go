package service

import (
	"net/http"

	"github.com/da4nik/swanager/api/common"
	"github.com/da4nik/swanager/core/entities"
	swarm_service "github.com/da4nik/swanager/core/swarm/service"
	"github.com/gin-gonic/gin"
)

// GetRoutesForRouter adds resource routes to api router
func GetRoutesForRouter(router *gin.RouterGroup) {

	services := router.Group("/services", common.Auth(true))
	{
		services.GET("", list)
		services.POST("", create)
	}

	service := services.Group("/:service_id")
	{
		service.GET("", show)
	}

	appServices := router.Group("/apps/:app_id/services", common.Auth(true))
	{
		appServices.GET("", list)
		appServices.POST("", create)
	}

	appService := appServices.Group("/:service_id")
	{
		appService.GET("", show)
	}
}

func list(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func create(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func show(c *gin.Context) {
	currentUser := common.MustGetCurrentUser(c)

	service, err := entities.GetService(gin.H{"user_id": currentUser.ID})
	if err != nil {
		common.RenderError(c, http.StatusNotFound, "Service not found")
		return
	}

	serviceStatus, err := swarm_service.Status(service)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"service": service, "status_error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"service": service, "status": serviceStatus})
}
