package service

import (
	"errors"
	"log"
	"math"
	"net/http"
	"project-mygram/dto"
	"project-mygram/entity"
	"project-mygram/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentService interface {
	Create(input entity.Comment) (res dto.CommentCreateUpdateResponse, statusCode int, err error)
	GetAll(ctx *gin.Context, queryparam dto.ListParam) (result dto.CommentListRes, statusCode int, err error)
	GetByID(ID uint64) (result entity.Comment, statusCode int, err error)
	UpdateByID(ID uint64, input dto.CommentCreateUpdateReq) (result dto.CommentCreateUpdateResponse, statusCode int, err error)
	DeleteByID(ID uint64) (statusCode int, err error)
}

type commentSrv struct {
	CommentRepository repository.CommentRepository
}

func NewCommentService(CommentRepo repository.CommentRepository) CommentService {
	return &commentSrv{
		CommentRepository: CommentRepo,
	}
}

func (srv *commentSrv) Create(input entity.Comment) (res dto.CommentCreateUpdateResponse, statusCode int, err error) {
	temp, err := srv.CommentRepository.Create(input)
	if err != nil {
		log.Printf("[CommentService-Create] error create data: %+v \n", err)
		return res, http.StatusInternalServerError, err
	}

	res = dto.CommentCreateUpdateResponse{
		ID:        temp.ID,
		Message:   temp.Message,
		PhotoID:   temp.PhotoID,
		UserID:    temp.UserID,
		CreatedAt: temp.CreatedAt,
		UpdatedAt: temp.UpdatedAt,
	}

	return res, http.StatusCreated, nil
}

func (srv *commentSrv) GetAll(ctx *gin.Context, queryparam dto.ListParam) (result dto.CommentListRes, statusCode int, err error) {
	result.Limit = queryparam.Limit
	result.Page = queryparam.Page

	result.Data, result.Total, err = srv.CommentRepository.GetAll(ctx, queryparam)
	if err != nil {
		log.Printf("[CommentService-GetAll] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	result.LastPage = uint64(math.Ceil(float64(result.Total) / float64(queryparam.Limit)))

	return result, http.StatusOK, nil
}

func (srv *commentSrv) GetByID(ID uint64) (result entity.Comment, statusCode int, err error) {
	result, err = srv.CommentRepository.GetByID(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || result.ID == 0 {
		return result, http.StatusNotFound, err
	}

	if err != nil {
		log.Printf("[CommentService-GetByID] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

func (srv *commentSrv) UpdateByID(ID uint64, input dto.CommentCreateUpdateReq) (result dto.CommentCreateUpdateResponse, statusCode int, err error) {
	sm, err := srv.CommentRepository.GetByID(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || sm.ID == 0 {
		return result, http.StatusNotFound, err
	}

	if err != nil {
		log.Printf("[CommentService-UpdateByID] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	temp, err := srv.CommentRepository.UpdateByID(ID, input)
	if err != nil {
		log.Printf("[CommentService-UpdateByID] error update data repo: %+v \n", err)
		return result, http.StatusInternalServerError, err
	}

	result = dto.CommentCreateUpdateResponse{
		ID:        temp.ID,
		Message:   temp.Message,
		PhotoID:   temp.PhotoID,
		UserID:    temp.UserID,
		CreatedAt: temp.CreatedAt,
		UpdatedAt: temp.UpdatedAt,
	}

	return result, http.StatusOK, nil
}

func (srv *commentSrv) DeleteByID(ID uint64) (statusCode int, err error) {
	sm, err := srv.CommentRepository.GetByID(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || sm.ID == 0 {
		return http.StatusNotFound, err
	}

	if err != nil {
		log.Printf("[CommentService-DeleteByID] error get data repo: %+v \n", err)
		return http.StatusInternalServerError, err
	}

	err = srv.CommentRepository.DeleteByID(ID)
	if err != nil {
		log.Printf("[CommentService-DeleteByID] error delete data repo: %+v \n", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
