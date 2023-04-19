package controller

import (
	"net/http"
	"project-mygram/dto"
	"project-mygram/helpers"
	"project-mygram/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userController struct {
	Service *service.Services
}

func NewUserController(srv *service.Services) UserController {
	return &userController{
		Service: srv,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var (
		input dto.RegisterReq
	)
	validate := validator.New()
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	err = validate.Struct(input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(errors, http.StatusBadRequest, "error"))
		return
	}

	userRes, statusHttp, err := c.Service.User.Register(input)
	if err != nil {
		ctx.JSON(statusHttp, helpers.APIResponse(err.Error(), statusHttp, "error"))
		return
	}

	ctx.JSON(statusHttp, userRes)
}

func (c *userController) Login(ctx *gin.Context) {
	var (
		input dto.LoginReq
	)

	validate := validator.New()
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	err = validate.Struct(input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(errors, http.StatusBadRequest, "error"))
		return
	}

	loginRes, statusHttp, err := c.Service.User.Login(input)
	if err != nil {
		ctx.JSON(statusHttp, helpers.APIResponse(err.Error(), statusHttp, "error"))
		return
	}

	ctx.JSON(statusHttp, loginRes)
}
