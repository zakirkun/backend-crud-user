package controller

import (
	"backend-crud-user/helper"
	"backend-crud-user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bindUri struct {
	ID int `uri:"id" binding:"required"`
}

type UserControlelr struct {
	serviceUser service.ServiceUser
}

func NewUserController(serviceUser service.ServiceUser) *UserControlelr {
	return &UserControlelr{serviceUser}
}

func (c *UserControlelr) GetData(ctx *gin.Context) {
	dataUser, err := c.serviceUser.FetchAll()

	if err != nil {
		result := helper.APIResponse("Failed Get Data", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	result := helper.APIResponse("Success Get Data", http.StatusOK, "success", dataUser)
	ctx.JSON(http.StatusOK, result)
}

func (c *UserControlelr) GetUserByID(ctx *gin.Context) {
	var geturi bindUri
	err := ctx.ShouldBindUri(&geturi)

	if err != nil {
		errorMessage := gin.H{"errors": err}

		response := helper.APIResponse("Failed to get", http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var ID = geturi.ID
	dataUser, err := c.serviceUser.GetByID(ID)

	if err != nil {
		response := helper.APIResponse("Failed get user", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	result := helper.APIResponse("Success Get Data", http.StatusOK, "success", dataUser)
	ctx.JSON(http.StatusOK, result)
}
