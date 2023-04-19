package controller

import (
	"errors"
	"net/http"
	"project-mygram/dto"
	"project-mygram/entity"
	"project-mygram/helpers"
	"project-mygram/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type PhotoController interface {
	CreatePhoto(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type photoHandler struct {
	Service *service.Services
}

func NewPhotoController(srv *service.Services) PhotoController {
	return &photoHandler{
		Service: srv,
	}
}

func (c *photoHandler) CreatePhoto(ctx *gin.Context) {
	var (
		input dto.PhotoCreateUpdateReq
	)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
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

	Photo := entity.Photo{
		UserID:   uint64(userData["id"].(float64)),
		Title:    input.Title,
		PhotoUrl: input.PhotoUrl,
		Caption:  input.Caption,
	}

	PhotoRes, httpStatus, err := c.Service.Photo.Create(Photo)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, PhotoRes)
}

func (c *photoHandler) GetAll(ctx *gin.Context) {
	var (
		paramPage  uint64 = 1
		paramLimit uint64 = 10
		err        error
	)

	if ctx.Query("page") == "" {
		paramPage, err = strconv.ParseUint(ctx.Query("page"), 10, 32)
		if err != nil {
			err = errors.New("query param page invalid")
			ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
			return
		}
	}

	if ctx.Query("limit") != "" {
		paramLimit, err = strconv.ParseUint(ctx.Query("limit"), 10, 32)
		if err != nil {
			err = errors.New("query param limit invalid")
			ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
			return
		}
	}

	param := dto.ListParam{
		Page:  paramPage,
		Limit: paramLimit,
	}

	res, httpStatus, err := c.Service.Photo.GetAll(ctx, param)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}
	ctx.JSON(httpStatus, res)
}

func (c *photoHandler) GetOne(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	res, httpStatus, err := c.Service.Photo.GetByID(ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}
	ctx.JSON(httpStatus, res)
}

func (c *photoHandler) UpdatePhoto(ctx *gin.Context) {
	var (
		input dto.PhotoCreateUpdateReq
	)

	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	validate := validator.New()
	err = ctx.ShouldBindJSON(&input)
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

	PhotoRes, httpStatus, err := c.Service.Photo.UpdateByID(ID, input)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, PhotoRes)
}

func (c *photoHandler) DeletePhoto(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	httpStatus, err := c.Service.Photo.DeleteByID(ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Deleted", httpStatus, "false"))
}
