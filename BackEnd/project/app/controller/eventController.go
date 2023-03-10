package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/source/infrastructure/utils"
	"strconv"
)

func (c *Controller) FindUserEvents(ctx *gin.Context) {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	detailedEvents, err := c.Service.SelectEventsByUserId(userIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//ctx.JSON(http.StatusOK, []entity.Event{
	//	detailedEvents,
	//})
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"events": detailedEvents,
	})

}
