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

type SocialMediaController interface {
	CreateSocialMedia(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaHandler struct {
	Service *service.Services
}

func NewSocialMediaController(srv *service.Services) SocialMediaController {
	return &socialMediaHandler{
		Service: srv,
	}
}

func (c *socialMediaHandler) CreateSocialMedia(ctx *gin.Context) {
	var (
		input dto.SocialMediaCreateReq
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

	socialMedia := entity.SocialMedia{
		UserID:         uint64(userData["id"].(float64)),
		Name:           input.Name,
		SocialMediaUrl: input.SocialMediaUrl,
	}

	SocialMediaRes, httpStatus, err := c.Service.SocialMedia.Create(socialMedia)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, SocialMediaRes)
}

func (c *socialMediaHandler) GetAll(ctx *gin.Context) {
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

	res, httpStatus, err := c.Service.SocialMedia.GetAll(ctx, param)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, res)
}

func (c *socialMediaHandler) GetOne(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	res, httpStatus, err := c.Service.SocialMedia.GetByID(ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}
	ctx.JSON(httpStatus, res)
}

func (c *socialMediaHandler) UpdateSocialMedia(ctx *gin.Context) {
	var (
		input dto.SocialMediaUpdateReq
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

	res, httpStatus, err := c.Service.SocialMedia.UpdateByID(ID, input)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, res)
}

func (c *socialMediaHandler) DeleteSocialMedia(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	httpStatus, err := c.Service.SocialMedia.DeleteByID(ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Deleted", httpStatus, "false"))
}
