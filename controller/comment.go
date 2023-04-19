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

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentHandler struct {
	Service *service.Services
}

func NewCommentController(srv *service.Services) CommentController {
	return &commentHandler{
		Service: srv,
	}
}

func (c *commentHandler) CreateComment(ctx *gin.Context) {
	var (
		input dto.CommentCreateUpdateReq
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

	comment := entity.Comment{
		UserID:  uint64(userData["id"].(float64)),
		PhotoID: input.PhotoID,
		Message: input.Message,
	}

	photo, httpStatus, err := c.Service.Photo.GetByID(input.PhotoID)
	if err != nil {
		if httpStatus == http.StatusNotFound {
			ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "photo not found"))
			return
		}
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error check photo"))
		return
	}

	if photo.UserID != comment.UserID {
		ctx.JSON(httpStatus, helpers.APIResponse("Forbidden", http.StatusForbidden, "You are not allowed to access this photo"))
		return
	}

	res, httpStatus, err := c.Service.Comment.Create(comment)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, res)
}

func (c *commentHandler) GetAll(ctx *gin.Context) {
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

	res, httpStatus, err := c.Service.Comment.GetAll(ctx, param)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, res)
}

func (c *commentHandler) GetOne(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	res, httpStatus, err := c.Service.Comment.GetByID(ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}
	ctx.JSON(httpStatus, res)
}

func (c *commentHandler) UpdateComment(ctx *gin.Context) {
	var (
		input dto.CommentCreateUpdateReq
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["id"].(float64))
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

	photo, httpStatus, err := c.Service.Photo.GetByID(input.PhotoID)
	if err != nil {
		if httpStatus == http.StatusNotFound {
			ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "photo not found"))
			return
		}
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error check photo"))
		return
	}

	if photo.UserID != userID {
		ctx.JSON(httpStatus, helpers.APIResponse("Forbidden", http.StatusForbidden, "You are not allowed to access this photo"))
		return
	}

	res, httpStatus, err := c.Service.Comment.UpdateByID(ID, input)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, res)
}

func (c *commentHandler) DeleteComment(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	httpStatus, err := c.Service.Comment.DeleteByID(ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Deleted", httpStatus, "false"))
}
